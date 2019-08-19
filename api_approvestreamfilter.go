package multichain

func (client *Client) ApproveStreamFilter(address, stream, filter string) (Response, error) {

	in := "{\"for\":" + stream + ", \"approve\":true}"
	msg := client.Command(
		"approvefrom",
		[]interface{}{
			address,
			filter,
			in,
		},
	)

	return client.Post(msg)
}
