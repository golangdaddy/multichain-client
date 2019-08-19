package multichain

func (client *Client) GetTransactionFilters(filters string, verbose bool) (Response, error) {

	msg := client.Command(
		"listtxfilters",
		[]interface{}{
			filters,
			verbose,
		},
	)

	return client.Post(msg)
}
