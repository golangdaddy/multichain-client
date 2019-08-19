package multichain

import (
	"fmt"
)

// Returns information about streams created on the blockchain. Pass a stream name, ref or creation txid in streams to retrieve information about one stream only, an array thereof for multiple streams, or * for all streams. Use count and start to retrieve part of the list only, with negative start values (like the default) indicating the most recently created streams. Extra fields are shown for streams to which this node has subscribed.
func (client *Client) ListStreams(streams string, start, count int, verbose bool) (Response, error) {

	if len(streams) == 0 {
		streams = "*"
	}

	return client.Post(
		map[string]interface{}{
			"jsonrpc": "1.0",
			"id": CONST_ID,
			"method": "liststreams",
			"params": []interface{}{
				fmt.Sprintf("%s", streams),
			},
		},
	)
}
