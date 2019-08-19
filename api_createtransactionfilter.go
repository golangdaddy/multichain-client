package multichain

func (client *Client) CreateTransactionFilter(name, restrictions, code string) (Response, error) {

	msg := client.Command(
		"create",
		[]interface{}{
			"txfilter",
			name,
			restrictions,
			code,
		},
	)

	return client.Post(msg)
}
