package multichain

func (client *Client) ListAddresses(address string) (Response, error) {

	msg := client.NodeMsg(
		"listaddresses",
		[]interface{}{},
	)

	return client.post(msg)
}
