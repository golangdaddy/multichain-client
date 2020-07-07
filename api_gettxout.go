package multichain

func (client *Client) GetTxOut(txid string, vout int) (Response, error) {

	msg := client.Command(
		"gettxoutdata",
		[]interface{}{
            txid,
            vout,
        },
	)

	return client.Post(msg)
}
