package multichain

func (client *Client) GetTxOut(txid string, vout int) (Response, error) {

	msg := map[string]interface{}{
		"jsonrpc": "1.0",
		"id": CONST_ID,
		"method": "gettxout",
		"params": []interface{}{
            txid,
            vout,
        },
	}

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
