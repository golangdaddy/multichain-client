package multichain

func (client *Client) SignRawTransaction(rawTransaction string, txDataArray []*TxData, privateKey string, args ...string) (Response, error) {

	if txDataArray == nil {
		txDataArray = []*TxData{}
	}

	params := []interface{}{
		rawTransaction,
		txDataArray,
		[]string{privateKey},
	}
	for _, arg := range args {
		params = append(params, arg)
	}

	msg := client.NodeMsg(
		"signrawtransaction",
		params,
	)

	return client.post(msg)
}
