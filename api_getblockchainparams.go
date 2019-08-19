package multichain

func (client *Client) GetBlockchainParams() (Response, error) {

	msg := client.Command(
		"getblockchainparams",
		[]interface{}{},
	)

	return client.Post(msg)
}
