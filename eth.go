package bitwallet

import "github.com/eavesmy/bitwallet/lib"
import "github.com/eavesmy/bitwallet/model"
import "strings"

type EthWallet struct {
	address    string
	apikey     string
	privatekey string
}

func (w *EthWallet) BalanceOf(addr string, contracts ...interface{}) (balance string, err error) {

	addr = strings.ToLower(addr)

	contract := ""

	if len(contracts) != 0 {
		contract = contracts[0].(string)
	}

	if contract == "" {
		return lib.QueryEthBalance(addr, w.apikey)
	}
	return lib.QueryEthTokenBalance(addr, contract, w.apikey)
}

func (w *EthWallet) QueryTransactionByContract(addr, contract string) ([]model.Transaction, error) {

	addr = strings.ToLower(addr)

	return lib.QueryEthTransactionByToken(addr, contract, w.apikey)
}
