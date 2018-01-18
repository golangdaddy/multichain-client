package multichain

func (client *Client) CreateRawExchange(txid string, vout int, asset string, quantity float64) (Response, error) {

	msg := client.Command(
		"createrawexchange",
		[]interface{}{
			txid,
			vout,
			map[string]interface{}{
				asset: quantity,
			},
		},
	)

	return client.Post(msg)
}
