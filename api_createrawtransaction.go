package multichain

func (client *Client) CreateRawTransaction(destinationAddress string, assets map[string]float64, unspentOutputs ...*Unspent) (Response, error) {

	msg := client.Msg(
		"createrawtransaction",
		[]interface{}{
			unspentOutputs,
			map[string]interface{}{
				destinationAddress: assets,
			},
		},
	)

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
