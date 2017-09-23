package multichain

func (client *Client) Issue(isOpen bool, accountAddress, assetName string, quantity float64, units float64) (Response, error) {

	msg := map[string]interface{}{
		"jsonrpc": "1.0",
		"id": CONST_ID,
		"method": "issue",
		"params": []interface{}{
			accountAddress,
			map[string]interface{}{
				"name": assetName,
				"open": isOpen,
			},
			quantity,
			units,
		},
	}

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
