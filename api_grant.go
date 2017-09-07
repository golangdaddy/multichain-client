package multichain

import (
    "strings"
)

func (client *Client) Grant(addresses, permissions []string) (Response, error) {

	msg := map[string]interface{}{
		"jsonrpc": "1.0",
		"id": CONST_ID,
		"method": "grant",
		"params": []interface{}{strings.Join(addresses, ","), strings.Join(permissions, ",")},
	}

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
