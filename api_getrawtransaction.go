package multichain

func (client *Client) GetRawTransaction(id string) (Response, error) {

	msg := client.Command(
		"getrawtransaction",
		[]interface{}{
			id,
		},
	)

	return client.Post(msg)
}
