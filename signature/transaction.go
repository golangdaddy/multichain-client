package signature

type Transaction struct {
	Txid     string `json:"txid"`
	Version  int    `json:"version"`
	Locktime int    `json:"locktime"`
	Vin      []struct {
		Txid      string `json:"txid"`
		Vout      int    `json:"vout"`
		ScriptSig struct {
			Asm string `json:"asm"`
			Hex string `json:"hex"`
		} `json:"scriptSig"`
		Sequence int64 `json:"sequence"`
	} `json:"vin"`
	Vout []struct {
		Value        int `json:"value"`
		N            int `json:"n"`
		ScriptPubKey struct {
			Asm       string   `json:"asm"`
			Hex       string   `json:"hex"`
			ReqSigs   int      `json:"reqSigs"`
			Type      string   `json:"type"`
			Addresses []string `json:"addresses"`
		} `json:"scriptPubKey"`
		Assets []struct {
			Name      string      `json:"name"`
			Issuetxid string      `json:"issuetxid"`
			Assetref  interface{} `json:"assetref"`
			Qty       int         `json:"qty"`
			Raw       int         `json:"raw"`
			Type      string      `json:"type"`
		} `json:"assets"`
	} `json:"vout"`
}
