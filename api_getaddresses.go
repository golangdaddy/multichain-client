package multichain

func (client *Client) GetAddresses(verbose bool) (Response, error) {

	msg := client.Command(
		"getaddresses",
		[]interface{}{
			verbose,
		},
	)

	return client.Post(msg)
}
