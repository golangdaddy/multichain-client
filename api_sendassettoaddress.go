package multichain

func (client *Client) SendAssetToAddress(accountAddress, assetName string, value float64) (Response, error) {

	msg := client.Msg(
		"sendassettoaddress",
		[]interface{}{
			accountAddress,
			assetName,
			value,
		},
	)

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
