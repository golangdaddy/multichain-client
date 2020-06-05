package multichain

import (
	"github.com/golangdaddy/multichain-client/address"
)

func (client *Client) CreateKeypair() ([]*address.KeyPair, error) {

	msg := client.Command(
		"createkeypairs",
		[]interface{}{},
	)

	obj, err := client.Post(msg)
	if err != nil {
		return nil, err
	}

	array := obj["result"].([]interface{})

	addresses := make([]*address.KeyPair, len(array))

	for i, v := range array {

		result := v.(map[string]interface{})

		addresses[i] = &address.KeyPair{
			Public: result["address"].(string),
			Private: result["privkey"].(string),
		}

	}

	return addresses, nil
}
