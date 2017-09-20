package multichain

func (client *Client) Issue(accountAddress, assetName string, value float64, units float64) (Response, error) {

	msg := map[string]interface{}{
		"jsonrpc": "1.0",
		"id": CONST_ID,
		"method": "issue",
		"params": []interface{}{
			accountAddress,
			assetName,
			value,
            units,
		},
	}

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
