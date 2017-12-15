package multichain

func (client *Client) SendRawTransaction(rawTransaction string) (Response, error) {

	msg := client.NodeMsg(
		"sendrawtransaction",
		[]interface{}{
			rawTransaction,
		},
	)

	return client.post(msg)
}
