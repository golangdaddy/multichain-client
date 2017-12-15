package multichain

// This works like preparelockunspent, but with control over the from-address whose funds are used to prepare the unspent transaction output. Any change from the transaction is send back to from-address.
func (client *Client) PrepareLockUnspentFrom(address, asset string, quantity float64) (Response, error) {

	msg := client.NodeMsg(
		"preparelockunspentfrom",
		[]interface{}{
			address,
			map[string]float64{
				asset: quantity,
			},
		},
	)

	return client.post(msg)
}
