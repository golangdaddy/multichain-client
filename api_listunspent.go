package multichain

// Returns a list of unspent transaction outputs in the wallet, with between minconf and maxconf confirmations. For a MultiChain blockchain, each transaction output includes assets and permissions fields listing any assets or permission changes encoded within that output. If the third parameter is provided, only outputs which pay an address in this array will be included.
func (client *Client) ListUnspent(address string) (Response, error) {

	msg := map[string]interface{}{
		"jsonrpc": "1.0",
		"id": CONST_ID,
		"method": "listunspent",
		"params": []interface{}{
			0,
			999999,
			[]string{
				address,
			},
		},
	}

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
