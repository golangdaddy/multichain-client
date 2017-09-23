package multichain

import "fmt"

func (client *Client) ListAddressTransactions(address string) (Response, error) {

	msg := map[string]interface{}{
		"jsonrpc": "1.0",
		"id": CONST_ID,
		"method": "listaddresstransactions",
		"params": []interface{}{
			address,
			0,
			0,
			true,
		},
	}

	obj, err := client.post(msg)
	if err != nil {
		fmt.Println(msg)
		return nil, err
	}

	return obj, nil
}
