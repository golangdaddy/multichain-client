package multichain

func (client *Client) SignRawTransaction(rawTransaction string, txDataArray []*Unspent, privateKey string, args ...string) (Response, error) {

	if txDataArray == nil {
		txDataArray = []*Unspent{}
	}

	params := []interface{}{
		rawTransaction,
		txDataArray,
		[]string{privateKey},
	}
	for _, arg := range args {
		params = append(params, arg)
	}

	msg := client.Command(
		"signrawtransaction",
		params,
	)

	return client.Post(msg)
}
