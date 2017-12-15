package multichain

func (client *Client) GetInfo() (Response, error) {

	msg := client.NodeMsg(
		"getinfo",
		[]interface{}{},
	)

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
