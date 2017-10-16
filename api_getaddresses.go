package multichain

import "fmt"

func (client *Client) GetAddresses(verbose bool) (Response, error) {

	msg := map[string]interface{}{
		"jsonrpc": "1.0",
		"id": CONST_ID,
		"method": "getaddresses",
		"params": []interface{}{
			fmt.Sprintf("verbose=%v", verbose),
		},
	}

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
