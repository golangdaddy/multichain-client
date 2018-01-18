package multichain

func (client *Client) GetAddressBalances(address string) (Response, error) {

	msg := client.Command(
		"getaddressbalances",
		[]interface{}{
			address,
		},
	)

	return client.Post(msg)
}
