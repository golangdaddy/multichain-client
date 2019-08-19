package multichain

func (client *Client) TestTransactionFilter(restrictions, code string) (Response, error) {

	msg := client.Command(
		"testtxfilter",
		[]interface{}{
			restrictions,
			code,
		},
	)

	return client.Post(msg)
}
