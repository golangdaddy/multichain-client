package multichain

func (client *Client) AppendRawData(tx, stream, key, data string) (Response, error) {

	return client.Post(
		client.Command(
			"appendrawdata",
			[]interface{}{
				tx,
				map[string]string{
					"for": stream,
					"key": key,
					"data": data,
				},
			},
		),
	)
}
