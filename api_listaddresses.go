package multichain

import "fmt"

func (client *Client) ListAddresses(verbose bool, addresses ...string) (Response, error) {

	v := fmt.Sprintf("verbose=%v", verbose)

	var params []interface{}

	if len(addresses) > 0 {
		params = []interface{}{
			addresses,
			v,
		}
	} else {
		params = []interface{}{
			v,
		}
	}

	msg := client.NodeMsg(
		"listaddresses",
		params,
	)

	return client.post(msg)
}
