package model

type Ret struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type RetTokenBalance struct {
	Network     string    `json:"network"`
	Hash        string    `json:"hash"`
	TokenInfo   TokenInfo `json:"tokenInfo"`
	TransferCnt int       `json:"transferCnt"`
	Balance     string    `json:"balance"`
}

type TokenInfo struct {
	H string `json:"h"`
	P string `json:"p"`
	F string `json:"f"`
	S string `json:"s"`
	D string `json:"d"`
	T string `json:"t"`
}

type Transaction struct {
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Valid             bool   `json:"valid"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	From              string `json:"from"`
	ContractAddress   string `json:"contractAddress"`
	To                string `json:"to"`
	Value             string `json:"value"`
	TokenName         string `json:"tokenName"`
	TokenSymbol       string `json:"tokenSymbol"`
	TokenDecimal      string `json:"tokenDecimal"`
	TransactionIndex  string `json:"transactionIndex"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	GasUsed           string `json:"gasUsed"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	Input             string `json:"input"`
	Confimations      string `json:"confirmations"`
}
