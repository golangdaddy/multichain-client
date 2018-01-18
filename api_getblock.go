package multichain

func (client *Client) GetBlock(heightOrHash string) (Response, error) {

	msg := client.Command(
		"getblock",
		[]interface{}{
			heightOrHash,
		},
	)

	return client.Post(msg)
}
