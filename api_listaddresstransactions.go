package multichain

import "fmt"

func (client *Client) ListAddressTransactions(address string) (Response, error) {

	msg := client.NodeMsg(
		"listaddresstransactions",
		[]interface{}{
			address,
			0,
			0,
			true,
		},
	)

	obj, err := client.post(msg)
	if err != nil {
		fmt.Println(msg)
		return nil, err
	}

	return obj, nil
}
