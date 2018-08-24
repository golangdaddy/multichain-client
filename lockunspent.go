package multichain

func (client *Client) LockUnspent(unlock bool, unspentOutputs []*Unspent) (Response, error) {
	msg := client.Command(
		"lockunspent",
		[]interface{}{
			unlock,
			unspentOutputs,
		},
	)

	return client.Post(msg)
}
