package multichain

func (client *Client) GetNewAddress() (Response, error) {

	msg := client.Command(
		"getnewaddress",
		[]interface{}{},
	)

	return client.Post(msg)
}
