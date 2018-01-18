package multichain

import (
    "strings"
)

// Grants permissions to addresses, a comma-separated list of addresses. For global permissions, set permissions to one of connect, send, receive, create, issue, mine, activate, admin, or a comma-separated list thereof. For per-asset or per-stream permissions, use the form entity.issue or entity.write,admin where entity is an asset or stream name, ref or creation txid. If the chain uses a native currency, you can send some to each recipient using the native-amount parameter. Returns the txid of the transaction granting the permissions.
func (client *Client) Grant(addresses, permissions []string) (Response, error) {

	msg := client.Command(
        "grant",
        []interface{}{
            strings.Join(addresses, ","),
            strings.Join(permissions, ","),
        },
	)

	return client.Post(msg)
}
