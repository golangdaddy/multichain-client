package multichain

func (client *Client) Issue(isOpen bool, accountAddress, assetName string, quantity float64, units float64) (Response, error) {

	msg := client.ChainMsg(
		"issue",
		[]interface{}{
			accountAddress,
			map[string]interface{}{
				"name": assetName,
				"open": isOpen,
			},
			quantity,
			units,
		},
	)

	return client.post(msg)
}
