package multichain

func (client *Client) GetInfo() (Response, error) {

	msg := client.Command(
		"getinfo",
		[]interface{}{},
	)

	return client.Post(msg)
}
