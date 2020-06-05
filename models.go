package multichain

type Unspent struct {
	Txid string `json:"txid"`
	Vout int `json:"vout"`
	ScriptPubKey string `json:"scriptPubKey"`
}
