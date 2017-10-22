package multichain

func (client *Client) DumpPrivKey(address string) (Response, error) {

	msg := map[string]interface{}{
		"jsonrpc": "1.0",
		"id": CONST_ID,
		"method": "dumpprivkey",
		"params": []interface{}{
			address,
		},
	}

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
