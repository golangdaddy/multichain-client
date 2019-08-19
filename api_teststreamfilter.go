package multichain

func (client *Client) TestStreamFilter(restrictions, code string) (Response, error) {

	msg := client.Command(
		"teststreamfilter",
		[]interface{}{
			restrictions,
			code,
		},
	)

	return client.Post(msg)
}
