package multichain

func (client *Client) GetAddressBalances(address string) (Response, error) {

	msg := client.NodeMsg(
		"getaddressbalances",
		[]interface{}{
			address,
		},
	)

	return client.post(msg)
}
