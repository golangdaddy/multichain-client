package multichain

func (client *Client) GetInfo() (Response, error) {

	msg := client.NodeMsg(
		"getinfo",
		[]interface{}{},
	)

	return client.post(msg)
}
