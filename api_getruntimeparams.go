package multichain

func (client *Client) GetRuntimeParams() (Response, error) {

	msg := client.Command(
		"getruntimeparams",
		[]interface{}{},
	)

	return client.Post(msg)
}
