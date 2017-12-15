package multichain

func (client *Client) SignRawTransaction(rawTransaction string, txDataArray []*TxData, privateKey string) (Response, error) {

	msg := client.NodeMsg(
		"signrawtransaction",
		[]interface{}{
			rawTransaction,
			txDataArray,
			[]string{privateKey},
		},
	)

	return client.post(msg)
}
