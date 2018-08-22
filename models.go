package multichain

type AddressKeyPair struct {
	Address string `json:"address"`
	PubKey  string `json:"pubkey"`
	PrivKey string `json:"privkey"`
}

type Unspent struct {
	Txid string `json:"txid"`
	Vout int    `json:"vout"`
}

type TxData struct {
	Txid         string `json:"txid"`
	Vout         int    `json:"vout"`
	ScriptPubKey string `json:"scriptPubKey"`
}
