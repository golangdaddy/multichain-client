package multichain

func (client *Client) ApproveTransactionFilter(address, filter string) (Response, error) {

	msg := client.Command(
		"approvefrom",
		[]interface{}{
			address,
			filter,
		},
	)

	return client.Post(msg)
}
