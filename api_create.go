package multichain

import (
	"fmt"
)

//Creates a new stream on the blockchain called name. For now, always pass the value "stream" in the type parameter – this is designed for future functionality. If open is true then anyone with global send permissions can publish to the stream, otherwise publishers must be explicitly granted per-stream write permissions.
func (client *Client) Create(typeName, name string, open bool) (Response, error) {

	msg := client.Msg(
		"create",
		[]interface{}{
			"type="+typeName,
			name,
			fmt.Sprintf("open=%v", open),
		},
	)

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
