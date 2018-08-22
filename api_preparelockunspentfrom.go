package multichain

import "fmt"

// This works like preparelockunspent, but with control over the from-address whose funds are used to prepare the unspent transaction output. Any change from the transaction is send back to from-address.
func (client *Client) PrepareLockUnspentFrom(address, asset string, quantity float64, lock bool) (Response, error) {

	msg := client.Command(
		"preparelockunspentfrom",
		[]interface{}{
			address,
			map[string]float64{
				asset: quantity,
			},
			fmt.Sprintf("lock=%v", lock),
		},
	)

	return client.Post(msg)
}
 