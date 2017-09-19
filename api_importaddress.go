package multichain

import (
//	"fmt"
)

// Adds address (or a full public key, or an array of either) to the wallet, without an associated private key. This creates one or more watch-only addresses, whose activity and balance can be retrieved via various APIs (e.g. with the includeWatchOnly parameter), but whose funds cannot be spent by this node. If rescan is true, the entire blockchain is checked for transactions relating to all addresses in the wallet, including the added ones. Returns null if successful.
func (client *Client) ImportAddress(pubKey, label string, rescan bool) (Response, error) {

	msg := map[string]interface{}{
		"jsonrpc": "1.0",
		"id": CONST_ID,
		"method": "importaddress",
		"params": []interface{}{pubKey, label, rescan},
	}

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
