package eth

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"hamster-paas/pkg/utils/logger"
)

type EthNetwork string

const MAINNET EthNetwork = "mainnet"
const GOERLI EthNetwork = "goerli"
const HAMSTER EthNetwork = "hamster"
const BSC_MAINNET EthNetwork = "bsc_mainnet"
const BSC_TESTNET EthNetwork = "bsc_testnet"

var netMap map[EthNetwork]string = make(map[EthNetwork]string)

func setup() {
	netMap[GOERLI] = "https://goerli.infura.io/v3/ce58d7af0a4a47ec9f3d18a3545f6d18"
	netMap[MAINNET] = "https://mainnet.infura.io/v3/ce58d7af0a4a47ec9f3d18a3545f6d18"
	netMap[HAMSTER] = "https://rpc-moonbeam.hamster.newtouch.com"
	netMap[BSC_MAINNET] = "https://bsc-dataseed1.defibit.io/"
	netMap[BSC_TESTNET] = "https://data-seed-prebsc-2-s1.binance.org:8545/"
}

func init() {
	setup()
}

type EthereumProxy interface {
	TransactionByHash(hash string) (tx *types.Transaction, isPending bool, err error)
}

type EthereumProxyFactory struct {
	clients map[EthNetwork]EthereumProxy
}

func NewEthereumProxyFactory() *EthereumProxyFactory {

	return &EthereumProxyFactory{
		clients: make(map[EthNetwork]EthereumProxy),
	}
}

// GetClient  get ethereum client
func (e *EthereumProxyFactory) GetClient(network EthNetwork) EthereumProxy {
	if val, ok := e.clients[network]; ok {
		//do
		return val
	} else {
		client, err := NewRPCEthereumProxy(netMap[network])
		if err != nil {
			return nil
		}
		e.clients[network] = client
		return client
	}
}

type RPCEthereumProxy struct {
	client *ethclient.Client
}

func NewRPCEthereumProxy(endpoints string) (*RPCEthereumProxy, error) {
	client, err := ethclient.Dial(endpoints)

	if err != nil {
		logger.Error("Oops! There was a problem", err)
		return nil, err
	} else {
		fmt.Println("Success! you are connected to the ", endpoints)
		return &RPCEthereumProxy{
			client,
		}, nil
	}
}

func (rpc *RPCEthereumProxy) TransactionByHash(hash string) (tx *types.Transaction, isPending bool, err error) {

	ctx := context.Background()
	hashTx := common.Hash(common.FromHex(hash))
	return rpc.client.TransactionByHash(ctx, hashTx)
}
