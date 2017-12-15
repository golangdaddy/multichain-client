package multichain

func (client *Client) CreateRawTransaction(destinationAddress string, assets map[string]float64, unspentOutputs ...*Unspent) (Response, error) {

	msg := client.NodeMsg(
		"createrawtransaction",
		[]interface{}{
			unspentOutputs,
			map[string]interface{}{
				destinationAddress: assets,
			},
		},
	)

	return client.post(msg)
}
