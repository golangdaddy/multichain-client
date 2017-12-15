package multichain

func (client *Client) GetTxOut(txid string, vout int) (Response, error) {

	msg := client.NodeMsg(
		"gettxout",
		[]interface{}{
            txid,
            vout,
        },
	)

	return client.post(msg)
}
