package multichain

func (client *Client) GetBlock(heightOrHash interface{}, verbosity bool) (Response, error) {

	msg := client.Command(
		"getblock",
		[]interface{}{
			heightOrHash,
			verbosity,
		},
	)

	return client.Post(msg)
}
