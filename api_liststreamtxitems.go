package multichain

// This works like liststreamitems, but listing items in stream within the given txid only.
// It should be used as a replacement for getstreamitem if multiple items are being published to a single stream
// in a single transaction.
func (client *Client) ListStreamTxItems(stream, txid string, verbose bool) (Response, error) {

	return client.Post(
		map[string]interface{}{
			"jsonrpc": "1.0",
			"id": CONST_ID,
			"method": "liststreamtxitems",
			"params": []interface{}{
				stream,
				txid,
				verbose,
			},
		},
	)
}
