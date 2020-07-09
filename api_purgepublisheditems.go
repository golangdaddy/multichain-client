package multichain

type TransactionOutput struct {
	Txid string
	Vout int
}
type blocks struct {
	Blocks string
}

// PurgeAllPublishedItems purges ALL off-chain items published by this node from local storage.
// Use with caution – this can lead to permanent data loss if no other node has retrieved the items.
//
// Returns statistics on how many chunks were purged and their total size. If this node is subscribed to the
// stream in which a purged item appears, the item will no longer be available. However this does not purge
// any item retrieved by this node from the network, unless that item is identical to a purged published item,
// and neither item is salted. To avoid these collisions, create streams with {"salted":true}. To explicitly
// purge retrieved items, use purgestreamitems.
//
// HINT: Available in MultiChain Enterprise only.
//
func (client *Client) PurgeAllPublishedItems() (Response, error) {
	return client.purgePublishedItems([]interface{}{
		"all",
	})
}

// PurgePublishedItemsByTxid purges off-chain items created within transactions of given txids,
// published by this node from local storage.
// Use with caution – this can lead to permanent data loss if no other node has retrieved the items.
//
// Returns statistics on how many chunks were purged and their total size. If this node is subscribed to the
// stream in which a purged item appears, the item will no longer be available. However this does not purge
// any item retrieved by this node from the network, unless that item is identical to a purged published item,
// and neither item is salted. To avoid these collisions, create streams with {"salted":true}. To explicitly
// purge retrieved items, use purgestreamitems.
//
// HINT: Available in MultiChain Enterprise only.
//
func (client *Client) PurgePublishedItemsByTxids(txids []string) (Response, error) {
	return client.purgePublishedItems([]interface{}{
		txids,
	})
}

// PurgePublishedItemsByTxOutputs purges off-chain items created within transactions of given txOuts, published
// by this node from local storage.
// Use with caution – this can lead to permanent data loss if no other node has retrieved the items.
//
// The txOuts parameter accepts a slice of {"txid":"id","vout":n} transaction outputs.
//
// Returns statistics on how many chunks were purged and their total size. If this node is subscribed to the
// stream in which a purged item appears, the item will no longer be available. However this does not purge
// any item retrieved by this node from the network, unless that item is identical to a purged published item,
// and neither item is salted. To avoid these collisions, create streams with {"salted":true}. To explicitly
// purge retrieved items, use purgestreamitems.
//
// HINT: Available in MultiChain Enterprise only.
//
func (client *Client) PurgePublishedItemsByTxOutputs(txOuts []TransactionOutput) (Response, error) {
	return client.purgePublishedItems([]interface{}{
		txOuts,
	})
}

// PurgePublishedItemsByBlockset purges off-chain items created within blocks defined by given blockSet and published
// by this node from local storage.
// Use with caution – this can lead to permanent data loss if no other node has retrieved the items.
//
// The blockSet parameter accepts  a {"blocks":blocks-set} set of blocks (where blocks-set
// takes any format accepted by listblocks).
//
// Returns statistics on how many chunks were purged and their total size. If this node is subscribed to the
// stream in which a purged item appears, the item will no longer be available. However this does not purge
// any item retrieved by this node from the network, unless that item is identical to a purged published item,
// and neither item is salted. To avoid these collisions, create streams with {"salted":true}. To explicitly
// purge retrieved items, use purgestreamitems.
//
// HINT: Available in MultiChain Enterprise only.
//
func (client *Client) PurgePublishedItemsByBlockset(blockSet string) (Response, error) {
	return client.purgePublishedItems([]interface{}{
		blocks{Blocks: blockSet},
	})
}

func (client *Client) purgePublishedItems(params []interface{}) (Response, error) {

	return client.Post(
		map[string]interface{}{
			"jsonrpc": "1.0",
			"id":      CONST_ID,
			"method":  "purgepublisheditems",
			"params":  params,
		},
	)
}
