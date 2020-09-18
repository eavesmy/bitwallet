package bitwallet

import "github.com/eavesmy/bitwallet/model"

type Chain string

const (
	CHAINBTC = "btc"
	CHAINETH = "eth"
)

type Wallet interface {
	BalanceOf(string, ...interface{}) (string, error)
	QueryTransactionByContract(string, string) ([]model.Transaction, error)
}

func New(apikey string, chain Chain) (wallet Wallet) {

	if apikey == "" {
		panic("Required apikey.")
	}

	if chain == CHAINBTC {
		return nil
	}

	return &EthWallet{apikey: apikey}

}

func NewFromPrivateKey(key string, chain Chain) {

}
