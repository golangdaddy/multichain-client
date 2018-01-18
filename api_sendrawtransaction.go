package multichain

func (client *Client) SendRawTransaction(rawTransaction string) (Response, error) {

	msg := client.Command(
		"sendrawtransaction",
		[]interface{}{
			rawTransaction,
		},
	)

	return client.Post(msg)
}
