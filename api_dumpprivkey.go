package multichain

func (client *Client) DumpPrivKey(address string) (Response, error) {

	msg := client.Command(
		"dumpprivkey",
		[]interface{}{
			address,
		},
	)

	return client.Post(msg)
}
