package multichain

func (client *Client) ListAddressTransactions(address string) (Response, error) {

	msg := client.NodeMsg(
		"listaddresstransactions",
		[]interface{}{
			address,
			0,
			0,
			true,
		},
	)

	return client.post(msg)
}
