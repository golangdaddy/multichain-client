package multichain

func (client *Client) AppendRawChange(tx, address string) (Response, error) {

	return client.Post(
		client.Command(
			"appendrawchange",
			[]interface{}{
				tx,
				address,
			},
		),
	)
}
