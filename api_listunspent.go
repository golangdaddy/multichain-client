package multichain

// Returns a list of unspent transaction outputs in the wallet, with between minconf and maxconf confirmations. For a MultiChain blockchain, each transaction output includes assets and permissions fields listing any assets or permission changes encoded within that output. If the third parameter is provided, only outputs which pay an address in this array will be included.
func (client *Client) ListUnspent(address string) (Response, error) {

	msg := client.Command(
		"listunspent",
		[]interface{}{
			0,
			999999,
			[]string{
				address,
			},
		},
	)

	return client.Post(msg)
}
