package multichain

func (client *Client) IssueMore(accountAddress, assetName string, value float64) (Response, error) {

	msg := client.ChainMsg(
		"issuemore",
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
