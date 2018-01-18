package multichain

// Adds address (or a full public key, or an array of either) to the wallet, without an associated private key. This creates one or more watch-only addresses, whose activity and balance can be retrieved via various APIs (e.g. with the includeWatchOnly parameter), but whose funds cannot be spent by this node. If rescan is true, the entire blockchain is checked for transactions relating to all addresses in the wallet, including the added ones. Returns null if successful.
func (client *Client) ImportAddress(pubKey, label string, rescan bool) (Response, error) {

	msg := client.Command(
		"importaddress",
		[]interface{}{
			pubKey,
			label,
			rescan,
		},
	)

	return client.Post(msg)
}
