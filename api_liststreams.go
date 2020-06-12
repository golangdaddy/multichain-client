package multichain

import (
	"strings"
)

// Returns information about streams created on the blockchain.
// Pass a slice of stream names, refs or creation txids to retrieve information about respective streams.
// Extra fields are shown for streams to which this node has subscribed.
func (client *Client) ListStreams(streams []string, verbose bool) (Response, error) {
	return client.listStreams(strings.Join(streams, ","), verbose)
}

// Returns information about all streams created on the blockchain.
// Extra fields are shown for streams to which this node has subscribed.
func (client *Client) ListAllStreams(verbose bool) (Response, error) {
	return client.listStreams("*", verbose)
}

func (client *Client) listStreams(streams string, verbose bool) (Response, error) {

	params := getListStreamsParams(streams, verbose)

	return client.Post(
		map[string]interface{}{
			"jsonrpc": "1.0",
			"id":      CONST_ID,
			"method":  "liststreams",
			"params":  params,
		},
	)
}

func getListStreamsParams(streams string, verbose bool) []interface{} {
	if len(streams) == 0 {
		streams = "*"
	}
	params := []interface{}{}
	var streamsParam interface{}
	if streams == "*" {
		streamsParam = streams
	} else {
		streamsParam = strings.Split(streams, ",")
	}
	params = append(params, streamsParam)
	params = append(params, verbose)
	return params
}
