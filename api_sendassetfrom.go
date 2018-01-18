package multichain

func (client *Client) SendAssetFrom(fromAddress, toAddress, assetType string, value float64) (Response, error) {

	msg := client.Command(
		"sendassetfrom",
		[]interface{}{
			fromAddress,
			toAddress,
			assetType,
			value,
		},
	)

	return client.Post(msg)
}
