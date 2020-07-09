package multichain

// PurgeAllStreamItems purges ALL selected off-chain items retrieved from stream (specified by name, ref or creation txid)
// from local storage and prevents their future automatic retrieval.
//
// Returns statistics on how many items were matched and how many chunks were purged and their total size.
// This does not purge items published by this node – see purgepublisheditems. Note also that if another stream
// contains an identical item, and neither item is salted, then the purged item may still be available in this
// stream. To avoid these collisions, create streams with {"salted":true}.
//
// HINT: Available in MultiChain Enterprise only.
//
func (client *Client) PurgeAllStreamItems(stream string) (Response, error) {
	return client.purgePublishedItems([]interface{}{
		stream,
		"all",
	})
}
// PurgeStreamItemsByTxid purges selected off-chain items retrieved from stream (specified by name, ref or creation txid)
// which got created within transactions defined by txids from local storage and prevents their future automatic retrieval.
//
// The txids parameter accepts a slice of transaction IDs.
//
// Returns statistics on how many items were matched and how many chunks were purged and their total size.
// This does not purge items published by this node – see purgepublisheditems. Note also that if another stream
// contains an identical item, and neither item is salted, then the purged item may still be available in this
// stream. To avoid these collisions, create streams with {"salted":true}.
//
// HINT: Available in MultiChain Enterprise only.
//
func (client *Client) PurgeStreamItemsByTxids(stream string, txids []string) (Response, error) {
	return client.purgePublishedItems([]interface{}{
		stream,
		txids,
	})
}
// PurgeStreamItemsByTxOutputs purges selected off-chain items retrieved from stream (specified by name, ref or creation txid)
// which got created within transactions of given txOuts from local storage and prevents their future automatic retrieval.
//
// The txOuts parameter accepts a slice of {"txid":"id","vout":n} transaction outputs.
//
// Returns statistics on how many items were matched and how many chunks were purged and their total size.
// This does not purge items published by this node – see purgepublisheditems. Note also that if another stream
// contains an identical item, and neither item is salted, then the purged item may still be available in this
// stream. To avoid these collisions, create streams with {"salted":true}.
//
// HINT: Available in MultiChain Enterprise only.
//
func (client *Client) PurgeStreamItemsByTxOutputs(stream string, txOuts []TransactionOutput) (Response, error) {
	return client.purgePublishedItems([]interface{}{
		stream,
		txOuts,
	})
}
// PurgeAllStreamItems purges selected off-chain items retrieved from stream (specified by name, ref or creation txid)
// which got created within blocks defined by given blockSet from local storage and prevents their future automatic retrieval.
//
// The blockSet parameter accepts a {"blocks":blocks-set} set of blocks (where blocks-set
// takes any format accepted by listblocks).
//
// Returns statistics on how many items were matched and how many chunks were purged and their total size.
// This does not purge items published by this node – see purgepublisheditems. Note also that if another stream
// contains an identical item, and neither item is salted, then the purged item may still be available in this
// stream. To avoid these collisions, create streams with {"salted":true}.
//
// HINT: Available in MultiChain Enterprise only.
//
func (client *Client) PurgeStreamItemsByBlockset(stream string, blockSet string) (Response, error) {
	return client.purgePublishedItems([]interface{}{
		stream,
		blocks{Blocks: blockSet},
	})
}
// PurgeAllStreamItems purges selected off-chain items retrieved from stream (specified by name, ref or creation txid)
// which match given keys and publishers from local storage and prevents their future automatic retrieval.
//
// The keys and publishers parameters define a query such as {"keys": keys, "publishers": publishers} accepted
// by liststreamqueryitems.
//
// Returns statistics on how many items were matched and how many chunks were purged and their total size.
// This does not purge items published by this node – see purgepublisheditems. Note also that if another stream
// contains an identical item, and neither item is salted, then the purged item may still be available in this
// stream. To avoid these collisions, create streams with {"salted":true}.
//
// HINT: Available in MultiChain Enterprise only.
//
func (client *Client) PurgeStreamItemsByQuery(stream string, keys, publishers []string) (Response, error) {
	return client.purgePublishedItems([]interface{}{
		stream,
		getQuery(keys, publishers),
	})
}

func (client *Client) purgeStreamItems(params []interface{}) (Response, error) {

	return client.Post(
		map[string]interface{}{
			"jsonrpc": "1.0",
			"id": CONST_ID,
			"method": "purgestreamitems",
			"params": params,
		},
	)
}
