package multichain

func (client *Client) CreateRawExchange(txid string, vout int, assets map[string]float64) (Response, error) {

	msg := client.Command(
		"createrawexchange",
		[]interface{}{
			txid,
			vout,
			assets,
		},
	)

	return client.Post(msg)
}
