package multichain

func (client *Client) ListAddressTransactions(address string) (Response, error) {

	msg := map[string]interface{}{
		"jsonrpc": "1.0",
		"id": CONST_ID,
		"method": "listaddresstransactions",
		"params": []interface{}{
			address,
			"verbose=true",
		},
	}

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
