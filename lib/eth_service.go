package lib

import (
	"encoding/json"
	"errors"
	// "github.com/eavesmy/glog"
	"github.com/eavesmy/bitwallet/model"
	gtype "github.com/eavesmy/golang-lib/type"
	"io/ioutil"
)

// var log = glog.New("gogo")

func QueryEthBalance(addr, apikey string) (balance string, err error) {
	res, err := Get("/addr/b/eth/" + addr + "?apikey=" + apikey)
	if err != nil || res == nil {
		return
	}
	defer res.Close()

	ret := &model.Ret{}
	b, _ := ioutil.ReadAll(res)

	json.Unmarshal(b, &ret)

	if ret.Code != 1 {
		err = errors.New(ret.Msg)
		return
	}

	return ret.Data.(string), nil
}

func QueryEthTokenBalance(addr, contract, apikey string) (balance string, err error) {
	res, err := Get("/eth/address/tokenbalance/" + addr + "?apikey=" + apikey)
	if err != nil || res == nil {
		return
	}
	defer res.Close()

	ret := &model.Ret{}

	b, _ := ioutil.ReadAll(res)
	json.Unmarshal(b, &ret)

	if ret.Code != 1 {
		err = errors.New(ret.Msg)
		return
	}

	for _, item := range ret.Data.([]interface{}) {

		_item := item.(map[string]interface{})

		if _item["hash"].(string) == contract {
			tokenInfo := _item["tokenInfo"].(map[string]interface{})
			keep := gtype.String2Int(tokenInfo["d"].(string))

			i := "1"

			for t := 0; t < keep; t++ {
				i += "0"
			}

			f_balance := gtype.String2Float64(_item["balance"].(string)) / gtype.String2Float64(i)
			balance = gtype.Float642String(f_balance, gtype.String2Int(tokenInfo["d"].(string)))
			break
		}
	}

	return
}

func QueryEthTransactionByToken(addr, contract, apikey string) (ts []model.Transaction, err error) {
	res, err := Get("/eth/address/tokentrans/" + addr + "/" + contract + "/1/50?apikey=" + apikey)

	if err != nil || res == nil {
		return
	}
	defer res.Close()

	ret := &model.Ret{}

	b, _ := ioutil.ReadAll(res)
	json.Unmarshal(b, &ret)

	if ret.Code != 1 {
		err = errors.New(ret.Msg)
		return
	}

	ts = []model.Transaction{}
	for _, item := range ret.Data.([]interface{}) {
		_item := item.(map[string]interface{})
		t := model.Transaction{
			TransactionIndex: gtype.Float642String(_item["index"].(float64), 0),
			BlockNumber:      gtype.Float642String(_item["block_no"].(float64), 0),
			TokenDecimal:     _item["tokenDecimals"].(string),
			ContractAddress:  _item["tokenAddr"].(string),
			TimeStamp:        gtype.Float642String(_item["time"].(float64), 0),
			Hash:             _item["txid"].(string),
			From:             _item["from"].(string),
			Confimations:     gtype.Float642String(_item["conformations"].(float64), 0),
			To:               _item["to"].(string),
			Value:            _item["value"].(string),
		}

		ts = append(ts, t)
	}

	return
}
