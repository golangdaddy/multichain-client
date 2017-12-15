package multichain

func (client *Client) DumpPrivKey(address string) (Response, error) {

	msg := client.NodeMsg(
		"dumpprivkey",
		[]interface{}{
			address,
		},
	)

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
