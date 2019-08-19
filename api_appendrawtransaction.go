package multichain

func (client *Client) AppendRawTransaction(tx string, inputs []*Unspent, outputs map[string]map[string]float64) (Response, error) {

	return client.Post(
		client.Command(
			"appendrawtransaction",
			[]interface{}{
				tx,
				inputs,
				outputs,
			},
		),
	)
}
