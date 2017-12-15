package multichain

import (
    "strings"
)

// This works like grant, but with control over the from-address used to grant the permissions. It is useful if the node has multiple addresses with administrator permissions.
func (client *Client) GrantFrom(fromAddress string, addresses, permissions []string) (Response, error) {

	msg := client.NodeMsg(
        "grantfrom",
		[]interface{}{
            fromAddress,
            strings.Join(addresses, ","),
            strings.Join(permissions, ","),
        },
	)

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
