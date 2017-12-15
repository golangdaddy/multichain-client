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

	return client.post(msg)
}
