package multichain

func (client *Client) SignRawTransaction(rawTransaction map[string]interface{}) (Response, error) {

	msg := map[string]interface{}{
		"jsonrpc": "1.0",
		"id": CONST_ID,
		"method": "signrawtransaction",
		"params": []interface{}{
			rawTransaction,
		},
	}

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
