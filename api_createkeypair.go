package multichain

func (client *Client) CreateKeypair() ([]*AddressKeyPair, error) {

	msg := client.Command(
		"createkeypairs",
		[]interface{}{},
	)

	obj, err := client.Post(msg)
	if err != nil {
		return nil, err
	}

	array := obj["result"].([]interface{})

	addresses := make([]*AddressKeyPair, len(array))

	for i, v := range array {

		result := v.(map[string]interface{})

		addresses[i] = &AddressKeyPair{
			Address: result["address"].(string),
			PubKey:  result["pubkey"].(string),
			PrivKey: result["privkey"].(string),
		}

	}

	return addresses, nil
}
