package multichain

import (
	"gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models/network"
)

func (client *Client) CreateKeypair() ([]*models.AddressKeyPair, error) {

	msg := client.Command(
		"createkeypairs",
		[]interface{}{},
	)

	obj, err := client.Post(msg)
	if err != nil {
		return nil, err
	}

	array := obj["result"].([]interface{})

	addresses := make([]*models.AddressKeyPair, len(array))

	for i, v := range array {

		result := v.(map[string]interface{})

		addresses[i] = &models.AddressKeyPair{
			Address: result["address"].(string),
			PubKey: result["pubkey"].(string),
			PrivKey: result["privkey"].(string),
		}

	}

	return addresses, nil
}
