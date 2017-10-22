package multichain

func (client *Client) SignRawTransaction(rawTransaction string, privateKey string) (Response, error) {

	msg := map[string]interface{}{
		"jsonrpc": "1.0",
		"id": CONST_ID,
		"method": "signrawtransaction",
		"params": []interface{}{
			rawTransaction,
			[]string{privateKey},
		},
	}

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
