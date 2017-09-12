package multichain

import (
	"fmt"
)

func (client *Client) ListStreams(streams string, start, count int, verbose bool) (Response, error) {

	msg := map[string]interface{}{
		"jsonrpc": "1.0",
		"id": CONST_ID,
		"method": "liststreams",
		"params": []interface{}{
			fmt.Sprintf("streams=%v", streams),
			fmt.Sprintf("start=%v", start),
			fmt.Sprintf("count=%v", count),
			fmt.Sprintf("verbose=%v", verbose),
		},
	}

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
