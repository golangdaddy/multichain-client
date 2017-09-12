package multichain

import (
	"fmt"
)

func (client *Client) Create(typeName, name string, open bool) (Response, error) {

	msg := map[string]interface{}{
		"jsonrpc": "1.0",
		"id": CONST_ID,
		"method": "create",
		"params": []interface{}{"type="+typeName, name, fmt.Sprintf("open=%v", open)},
	}

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
