package multichain

func (client *Client) GetTxOut(txid string, vout int) (Response, error) {

	msg := client.NodeMsg(
		"gettxout",
		[]interface{}{
            txid,
            vout,
        },
	)

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
