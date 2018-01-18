package multichain

func (client *Client) IssueMore(accountAddress, assetName string, value float64) (Response, error) {

	msg := client.Command(
		"issuemore",
		[]interface{}{
			accountAddress,
			assetName,
			value,
		},
	)

	return client.Post(msg)
}
