package multichain

// Returns a base64-encoded digital signature which proves that message was approved by the owner of address (which must belong to this wallet) or any other private key given in privkey. The signature can be verified by any node using the verifymessage command.
func (client *Client) SignMessage(addressOrPrivKey, message string) (Response, error) {

	msg := client.NodeMsg(
		"signmessage",
		[]interface{}{
			addressOrPrivKey,
			message,
		},
	)

	return client.post(msg)
}
