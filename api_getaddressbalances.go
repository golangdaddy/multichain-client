package multichain

func (client *Client) GetAddressBalances(address string) (Response, error) {

	msg := client.NodeMsg(
		"getaddressbalances",
		[]interface{}{
			address,
		},
	)

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
