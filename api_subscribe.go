package multichain

import (
	"errors"
	"strings"
)

type IndexType string

const (
	IndexItems           IndexType = "items"
	IndexKeys            IndexType = "keys"
	IndexPublishers      IndexType = "publishers"
	IndexItemsLocal      IndexType = "items-local"
	IndexKeysLocal       IndexType = "keys-local"
	IndexPublishersLocal IndexType = "publishers-local"
)

func (sp IndexType) IsValid() error {
	switch sp {
	case IndexItems, IndexKeys, IndexPublishers, IndexItemsLocal, IndexKeysLocal, IndexPublishersLocal:
		return nil
	}
	return errors.New("Invalid IndexType.")
}

// Instructs the node to start tracking one or more asset(s) or stream(s).
//
// These are specified using a name, ref or creation/issuance txid, or for multiple items, an array thereof.
// If rescan is true, the node will reindex all items from when the assets and/or streams were created, as well as
// those in other subscribed entities.
// If retrieveAllOffchain is true, all offchain items are retrieved.
// List of IndexTypes are built for stream.
// Returns null if successful.
// See also the autosubscribe runtime parameter.
func (client *Client) Subscribe(streamUid string, rescan bool, retrieveAllOffchain bool, indexTypes []IndexType) (Response, error) {

	params := []interface{}{
		streamUid,
		rescan,
	}
	params, err := appendInnerParams(indexTypes, retrieveAllOffchain, params)
	if err != nil {
		return nil, err
	}
	msg := client.Command(
		"subscribe",
		params,
	)

	return client.Post(msg)
}

func appendInnerParams(indexTypes []IndexType, retrieveAllOffchain bool, params []interface{}) ([]interface{}, error) {
	var innerParams []string
	for _, indexType := range indexTypes {
		if err := indexType.IsValid(); err != nil {
			return nil, err
		}
		innerParams = append(innerParams, string(indexType))
	}
	if retrieveAllOffchain {
		innerParams = append(innerParams, "retrieve")
	}
	if len(innerParams) > 0 {
		params = append(params, strings.Join(innerParams, ","))
	}
	return params, nil
}
