package multichain

// This works like liststreamitems, but listing items with the given key only.
func (client *Client) ListStreamKeyItems(stream, key string, count int, verbose bool) (Response, error) {

	return client.Post(
		map[string]interface{}{
			"jsonrpc": "1.0",
			"id": CONST_ID,
			"method": "liststreamkeyitems",
			"params": []interface{}{
				stream,
				key,
				verbose,
				count,
			},
		},
	)
}
