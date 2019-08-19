package multichain

func (client *Client) GetStreamFilters(filters string, verbose bool) (Response, error) {

	msg := client.Command(
		"liststeamfilters",
		[]interface{}{
			filters,
			verbose,
		},
	)

	return client.Post(msg)
}
