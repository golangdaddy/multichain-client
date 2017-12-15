package multichain

func (client *Client) DumpPrivKey(address string) (Response, error) {

	msg := client.NodeMsg(
		"dumpprivkey",
		[]interface{}{
			address,
		},
	)

	return client.post(msg)
}
