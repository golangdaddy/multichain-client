package multichain

func (client *Client) GetAddresses(verbose bool) (Response, error) {

	msg := client.NodeMsg(
		"getaddresses",
		[]interface{}{
			verbose,
		},
	)

	return client.post(msg)
}
