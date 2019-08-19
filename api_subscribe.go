package multichain

// Instructs the node to start tracking one or more asset(s) or stream(s). These are specified using a name, ref or creation/issuance txid, or for multiple items, an array thereof. If rescan is true, the node will reindex all items from when the assets and/or streams were created, as well as those in other subscribed entities. Returns null if successful. See also the autosubscribe runtime parameter.
func (client *Client) Subscribe(streamUid string, rescan bool) (Response, error) {

	msg := client.Command(
		"subscribe",
		[]interface{}{
			streamUid,
			rescan,
		},
	)

	return client.Post(msg)
}
