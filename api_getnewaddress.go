package multichain

func (client *Client) GetNewAddress() (Response, error) {

	msg := client.ChainMsg(
		"getnewaddress",
		[]interface{}{},
	)

	return client.post(msg)
}
