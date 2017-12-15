package multichain

func (client *Client) SendAssetFrom(fromAddress, toAddress, assetType string, value float64) (Response, error) {

	msg := client.NodeMsg(
		"sendassetfrom",
		[]interface{}{
			fromAddress,
			toAddress,
			assetType,
			value,
		},
	)

	return client.post(msg)
}
