package multichain

// Returns a list of assets in multichain
func (client *Client) ListAssets(assets string, verbose bool, count int, start int) (Response, error) {

	msg := client.Command(
		"listassets",
		[]interface{}{
			assets,
			verbose,
			count,
			start,
		},
	)

	return client.Post(msg)
}
