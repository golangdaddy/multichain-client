package multichain

type Unspent struct {
	Txid string `json:"txid"`
	Vout int `json:"vout"`
}
