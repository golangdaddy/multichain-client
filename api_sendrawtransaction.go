package multichain

func (client *Client) SendRawTransaction(rawTransaction string) (Response, error) {

	msg := map[string]interface{}{
		"jsonrpc": "1.0",
		"id": CONST_ID,
		"method": "sendrawtransaction",
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
