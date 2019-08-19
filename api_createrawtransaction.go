package multichain

func (client *Client) CreateRawTransaction(inputs []*Unspent, outputs map[string]map[string]float64) (Response, error) {

	msg := client.Command(
		"createrawtransaction",
		[]interface{}{
			inputs,
			outputs,
			[]string{},
			"lock",
		},
	)

	return client.Post(msg)
}
