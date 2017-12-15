package multichain

func (client *Client) GetAddresses(verbose bool) (Response, error) {

	msg := client.NodeMsg(
		"getaddresses",
		[]interface{}{
			verbose,
		},
	)

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
