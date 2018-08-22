package multichain

//Adds a privkey private key (or an array thereof) to the wallet, together with its associated public address. If rescan is true, the entire blockchain is checked for transactions relating to all addresses in the wallet, including the added ones. Returns null if successful.
func (client *Client) ImportPrivKey(privKey, label string, rescan bool) (Response, error) {

	msg := map[string]interface{}{
		"jsonrpc": "1.0",
		"id":      CONST_ID,
		"method":  "importprivkey",
		"params": []interface{}{
			[]string{privKey},
			label,
			rescan,
		},
	}

	return client.Post(msg)
}
