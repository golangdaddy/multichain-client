package multichain

func (client *Client) GetTxOut(txid string, vout int) (Response, error) {

	msg := client.Command(
		"gettxout",
		[]interface{}{
            txid,
            vout,
        },
	)

	return client.Post(msg)
}
