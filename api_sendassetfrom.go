package multichain

func (client *Client) SendAssetFrom(fromAddress, toAddress, assetType string, value float64) (Response, error) {

	msg := map[string]interface{}{
		"jsonrpc": "1.0",
		"id": CONST_ID,
		"method": "getaddressbalances",
		"params": []interface{}{
			fromAddress,
			toAddress,
			assetType,
			value,
		},
	}

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
