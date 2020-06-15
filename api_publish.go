package multichain

// Publishes an item in a stream (passed as a stream name, ref or creation txid). Data is held inside blockchain.
//
// Pass a single textual key or an array thereof in the key(s) parameter.
// The data parameter can accept raw hexadecimal data like a1b2c3d4, a reference to the binary cache
// {"cache":"Ev1HQV1aUCY"}, textual data {"text":"hello world"} or JSON data {"json":{"i":[1,2],"j":"yes"}}.
// Returns the txid of the transaction sent.
// To easily publish multiple items in a single transaction, see the publishmulti(from) command.
func (client *Client) PublishOnchain(stream string, keys []string, data interface{}) (Response, error) {
	return client.publish(stream, keys, data, false)
}

// Publishes an item in a stream (passed as a stream name, ref or creation txid). Data is held outside blockchain
// (offchain).
//
// Pass a single textual key or an array thereof in the key(s) parameter.
// The data parameter can accept raw hexadecimal data like a1b2c3d4, a reference to the binary cache
// {"cache":"Ev1HQV1aUCY"}, textual data {"text":"hello world"} or JSON data {"json":{"i":[1,2],"j":"yes"}}.
// Returns the txid of the transaction sent.
// To easily publish multiple items in a single transaction, see the publishmulti(from) command.
func (client *Client) PublishOffchain(stream string, keys []string, data interface{}) (Response, error) {
	return client.publish(stream, keys, data, true)
}

func (client *Client) publish(stream string, keys []string, data interface{}, offchain bool) (Response, error) {

	params := []interface{}{
		stream,
		keys,
		data,
	}
	if offchain {
		params = append(params, "offchain")
	}
	msg := client.Command(
		"publish",
		params,
	)

	return client.Post(msg)
}
