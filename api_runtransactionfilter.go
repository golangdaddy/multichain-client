package multichain

func (client *Client) RunTransactionFilter(filter string) (Response, error) {

	msg := client.Command(
		"runtxfilter",
		[]interface{}{
			filter,
		},
	)

	return client.Post(msg)
}
