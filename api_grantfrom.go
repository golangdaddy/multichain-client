package multichain

import (
    "strings"
)

// This works like grant, but with control over the from-address used to grant the permissions. It is useful if the node has multiple addresses with administrator permissions.
func (client *Client) GrantFrom(fromAddress string, addresses, permissions []string) (Response, error) {

	msg := client.Command(
        "grantfrom",
		[]interface{}{
            fromAddress,
            strings.Join(addresses, ","),
            strings.Join(permissions, ","),
        },
	)

	return client.Post(msg)
}
 