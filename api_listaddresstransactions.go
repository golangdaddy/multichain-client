package multichain

func (client *Client) ListAddressTransactions(address string, count int, skip int, verbose bool) (Response, error) {

	msg := client.Command(
		"listaddresstransactions",
		[]interface{}{
			address,
			count,
			skip,
			verbose,
		},
	)

	return client.Post(msg)
}
