package multichain

import "fmt"

func (client *Client) ListAddressTransactions(address string, count, skip int, verbose bool) (Response, error) {

	msg := client.Command(
		"listaddresstransactions",
		[]interface{}{
			address,
			fmt.Sprintf("count=%v", count),
			fmt.Sprintf("skip=%v", skip),
			fmt.Sprintf("verbose=%v", verbose),
		},
	)

	return client.Post(msg)
}
