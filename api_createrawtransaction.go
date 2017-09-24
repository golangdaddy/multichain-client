package multichain

func (client *Client) CreateRawTransaction(destinationAddress string, assets map[string]float64, unspentOutputs ...*Unspent) (Response, error) {

	msg := map[string]interface{}{
		"jsonrpc": "1.0",
		"id": CONST_ID,
		"method": "createrawtransaction",
		"params": []interface{}{
			unspentOutputs,
			map[string]interface{}{
				destinationAddress: assets,
			},
		},
	}

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
