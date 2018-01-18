package multichain

func (client *Client) GetBestBlockHash(heightOrHash string) (Response, error) {

	msg := client.Command(
		"getbestblockhash",
		[]interface{}{},
	)

	return client.Post(msg)
}
