package vo

import "time"

type ChainLinkRequest struct {
	Name   string `json:"name"`
	Script string `json:"script"`
	UserId uint64 `json:"userId"`
}
type ChainLinkRequestParam struct {
	Name   string `json:"name"`
	Script string `json:"script"`
}

type ChainLinkConsumers struct {
	Address    string    `json:"address"`
	Network    string    `json:"network"`
	DeployTime time.Time `json:"deploy_time"`
}

type ChainLinkRequestExecParam struct {
	SubscriptionId  int64   `json:"subscriptionId"`
	ConsumerAddress string  `json:"consumerAddress"`
	Secretsloction  int8    `json:"secretsloction"`
	SecretUrl       string  `json:"secretUrl"`
	Args            string  `json:"args"`
	TransactionTx   string  `json:"transactionTx"`
	RequestName     string  `json:"requestName"`
	RequestId       string  `json:"requestId"`
	Amount          float64 `json:"amount"`
}

type ChainLinkSubscriptionCreateParam struct {
	Chain          string `json:"chain"`
	Network        string `json:"network"`
	Name           string `json:"name"`
	SubscriptionId int64  `json:"subscriptionId"`
	Admin          string `json:"admin"`
	TransactionTx  string `json:"transactionTx"`
}

type ChainLinkFoundParam struct {
	ConsumerAddress string `json:"consumer_address"`
	Incr            string `json:"incr"`
	TransactionTx   string `json:"transaction_tx"`
}

type ChainLinkConsumerCreateParam struct {
	SubscriptionId  int64  `json:"subscriptionId"`
	ConsumerAddress string `json:"consumer_address"`
	TransactionTx   string `json:"transactionTx"`
}

type ChainLinkFoundUpdateParam struct {
	Id              int64  `json:"id"`
	SubscriptionId  int64  `json:"subscriptionId"`
	ConsumerAddress string `json:"consumerAddress"`
	TransactionTx   string `json:"transactionTx"`
	NewStatus       string `json:"newStatus"`
}

type ChainLinkConsumerUpdateParam struct {
	Id              int64  `json:"id"`
	SubscriptionId  int64  `json:"subscriptionId"`
	ConsumerAddress string `json:"consumerAddress"`
	TransactionTx   string `json:"transactionTx"`
	NewStatus       string `json:"newStatus"`
	UserId          uint64 `json:"userId"`
}

type ChainLinkSubscriptionUpdateParam struct {
	Id                  uint   `json:"id"`
	ChainSubscriptionId uint   `gorm:"column:chain_subscription_id" json:"subscriptionId"`
	Chain               string `json:"chain"`
	Network             string `json:"network"`
	UserId              uint64 `json:"userId"`
	TransactionTx       string `json:"transactionTx"`
	NewStatus           string `json:"newStatus"`
}
