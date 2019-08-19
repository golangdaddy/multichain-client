package multichain

func (client *Client) CreateStreamFilter(name, restrictions, code string) (Response, error) {

	msg := client.Command(
		"create",
		[]interface{}{
			"streamfilter",
			name,
			restrictions,
			code,
		},
	)

	return client.Post(msg)
}
