package multichain

// Returns a list of unspent transaction outputs in the wallet, with between minconf and maxconf confirmations. For a MultiChain blockchain, each transaction output includes assets and permissions fields listing any assets or permission changes encoded within that output. If the third parameter is provided, only outputs which pay an address in this array will be included.
func (client *Client) ListAssets(assets string, verbose bool, count int, start int) (Response, error) {

	msg := client.Command(
		"listassets",
		[]interface{}{
			assets,
			verbose,
			count,
			start,
		},
	)

	return client.Post(msg)
}
