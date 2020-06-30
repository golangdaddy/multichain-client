package multichain

// Retrieves a specific item with txid from stream, passed as a stream name, ref or creation txid, to which the
// node must be subscribed.
// Set verbose to true for additional information about the item’s transaction.
// If an item’s data as stored in the transaction is larger than the maxshowndata runtime parameter, it will be
// returned as an object with parameters for gettxoutdata.
// If the transaction contains more than one item for the stream, this will return an error –
//	use liststreamtxitems instead.
func (client *Client) GetStreamItem(stream, txid string, verbose bool) (Response, error) {

	return client.Post(
		map[string]interface{}{
			"jsonrpc": "1.0",
			"id": CONST_ID,
			"method": "getstreamitem",
			"params": []interface{}{
				stream,
				txid,
				verbose,
			},
		},
	)
}
