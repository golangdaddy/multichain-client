package multichain

// Returns a base64-encoded digital signature which proves that message was approved by the owner of address (which must belong to this wallet) or any other private key given in privkey. The signature can be verified by any node using the verifymessage command.
func (client *Client) SignMessage(addressOrPrivKey, message string) (Response, error) {

	msg := map[string]interface{}{
		"jsonrpc": "1.0",
		"id": CONST_ID,
		"method": "signmessage",
		"params": []interface{}{
			addressOrPrivKey,
			message,
		},
	}

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
