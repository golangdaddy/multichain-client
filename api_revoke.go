package multichain

import (
    "strings"
)

// Revokes permissions from addresses, a comma-separated list of addresses. The permissions parameter works the same as for grant. This is equivalent to calling grant with start-block=0 and end-block=0. Returns the txid of transaction revoking the permissions. For more information, see permissions management.
func (client *Client) Revoke(addresses, permissions []string) (Response, error) {

	msg := client.Command(
        "revoke",
        []interface{}{
            strings.Join(addresses, ","),
            strings.Join(permissions, ","),
        },
	)

	return client.Post(msg)
}
