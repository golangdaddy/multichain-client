package multichain

func (client *Client) GetBestBlockHash(heightOrHash string) (Response, error) {

	msg := client.NodeMsg(
		"getbestblockhash",
		[]interface{}{},
	)

	return client.post(msg)
}
