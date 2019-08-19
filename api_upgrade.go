package multichain

//When using create upgrade, the following blockchain parameters can now be included: target-block-time maximum-block-size max-std-tx-size max-std-op-returns-count max-std-op-return-size max-std-op-drops-count max-std-element-size. Note that these upgrade parameters can only be applied to chains running multichain protocol 20002 or later. In addition, to prevent abrupt changes in blockchain capacity and performance, the following constraints apply: The target-block-time parameter cannot be changed more than once per 100 blocks. All seven capacity-related parameters cannot be changed to less than half, or more than double, their previous size. e.g. create upgrade upgrade1 false '{"target-block-time":10,"maximum-block-size":16777216}'
func (client *Client) Upgrade(name string, params map[string]int) (Response, error) {

	msg := client.Command(
		"create",
		[]interface{}{
			"upgrade",
			name,
			false,
			params,
		},
	)
	return client.Post(msg)
}
