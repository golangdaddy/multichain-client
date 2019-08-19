package multichain

import (
	"fmt"
)

//Creates a new stream on the blockchain called name. For now, always pass the value "stream" in the type parameter – this is designed for future functionality. If open is true then anyone with global send permissions can publish to the stream, otherwise publishers must be explicitly granted per-stream write permissions.
func (client *Client) Publish(typeName, name string, open bool) (Response, error) {

	msg := client.Command(
		"create",
		[]interface{}{
			"type="+typeName,
			name,
			fmt.Sprintf("open=%v", open),
		},
	)

	return client.Post(msg)
}
