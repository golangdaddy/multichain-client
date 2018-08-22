package multichain

func (client *Client) SendAssetToAddress(accountAddress, assetName string, value float64) (Response, error) {

	msg := client.Command(
		"sendassettoaddress",
		[]interface{}{
			accountAddress,
			assetName,
			value,
		},
	)

	return client.Post(msg)
}
 