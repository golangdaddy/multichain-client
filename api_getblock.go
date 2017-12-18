package multichain

func (client *Client) GetBlock(heightOrHash string) (Response, error) {

	msg := client.NodeMsg(
		"getblock",
		[]interface{}{
			heightOrHash,
		},
	)

	return client.post(msg)
}
