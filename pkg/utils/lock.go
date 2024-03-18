package utils

import (
	"hamster-paas/pkg/utils/logger"
	"os"
	"time"

	"github.com/werf/kubedog/pkg/kube"
	"github.com/werf/lockgate"
	"github.com/werf/lockgate/pkg/distributed_locker"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	ICP_LOCK_KEY = "icp-lock"
	NAMESPACE    = "hamster"
)

var locker *distributed_locker.DistributedLocker

func init() {
	//locker, err := file_locker.NewFileLocker("/tmp/mylock")

	if err := kube.Init(kube.InitOptions{}); err != nil {
		logger.Errorf("cannot initialize kube: %s", err)
		os.Exit(1)
	}

	locker = distributed_locker.NewKubernetesLocker(
		kube.DynamicClient, schema.GroupVersionResource{
			Group:    "",
			Version:  "v1",
			Resource: "configmaps",
		}, ICP_LOCK_KEY, NAMESPACE,
	)
}

func Lock() (*lockgate.LockHandle, error) {
	// Case 1: simple blocking lock
	logger.Info("try to lock:", ICP_LOCK_KEY)

	_, lock, err := locker.Acquire(ICP_LOCK_KEY, lockgate.AcquireOptions{Shared: false, Timeout: 30 * time.Second})
	if err != nil {
		logger.Error(os.Stderr, "ERROR: failed to lock %s: %s\n", ICP_LOCK_KEY, err)
		return nil, err
	}
	return &lock, err
}

func Unlock(lock *lockgate.LockHandle) error {
	logger.Info("try to release lock:", ICP_LOCK_KEY)
	return locker.Release(*lock)
}
