package multichain

// This works like createrawtransaction, except it automatically selects the transaction inputs from those belonging to from-address, to cover the appropriate amounts. One or more change outputs going back to from-address will also be added to the end of the transaction.
func (client *Client) CreateRawSendFrom(watchAddress, destinationAddress string, assets map[string]float64) (Response, error) {

	msg := client.NodeMsg(
		"createrawsendfrom",
		[]interface{}{
			watchAddress,
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
