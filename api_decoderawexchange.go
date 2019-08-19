package multichain

// Decodes the raw exchange transaction in tx-hex, given by a previous call to createrawexchange or appendrawexchange. Returns details on the offer represented by the exchange and its present state. The offer field in the response lists the quantity of native currency and/or assets which are being offered for exchange. The ask field lists the native currency and/or assets which are being asked for. The candisable field specifies whether this wallet can disable the exchange transaction by double-spending against one of its inputs. The cancomplete field specifies whether this wallet has the assets required to complete the exchange. The complete field specifies whether the exchange is already complete (i.e. balanced) and ready for sending. If verbose is true then all of the individual stages in the exchange are listed. Other fields relating to fees are only relevant for blockchains which use a native currency.
func (client *Client) DecodeRawExchange(rawExchange string) (Response, error) {

	msg := client.Command(
		"decoderawexchange",
		[]interface{}{
			rawExchange,
		},
	)

	return client.Post(msg)
}
