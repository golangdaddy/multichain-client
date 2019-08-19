package multichain

// Lists items in stream, passed as a stream name, ref or creation txid. Set verbose to true for additional information about each item’s transaction. Use count and start to retrieve part of the list only, with negative start values (like the default) indicating the most recent items. Set local-ordering to true to order items by when first seen by this node, rather than their order in the chain. If an item’s data is larger than the maxshowndata runtime parameter, it will be returned as an object whose fields can be used with gettxoutdata.
func (client *Client) ListStreamItems(stream string, start, count int, verbose bool) (Response, error) {

	return client.Post(
		map[string]interface{}{
			"jsonrpc": "1.0",
			"id": CONST_ID,
			"method": "liststreamitems",
			"params": []interface{}{
				stream,
				verbose,
				999999,
			},
		},
	)
}
