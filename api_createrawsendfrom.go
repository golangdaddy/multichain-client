package multichain

// This works like createrawtransaction, except it automatically selects the transaction inputs from those belonging to from-address, to cover the appropriate amounts. One or more change outputs going back to from-address will also be added to the end of the transaction.
func (client *Client) CreateRawSendFrom(watchAddress, destinationAddress string, assets map[string]float64) (Response, error) {

	msg := client.Command(
		"createrawsendfrom",
		[]interface{}{
			watchAddress,
			map[string]interface{}{
				destinationAddress: assets,
			},
		},
	)

	return client.Post(msg)
}
