package multichain

import (
	"errors"
)

// This works like liststreamitems, but listing items in stream which match all of the specified keys and/or
// publishers in query.
// The query is an object with a key or keys field, and/or a publisher or publishers field. If present, key and
// publisher should specify a single key or publisher respectively, whereas keys and publishers should specify
// arrays thereof.
// Note that, unlike other stream retrieval APIs, liststreamqueryitems cannot rely completely on prior indexing,
// so the maxqueryscanitems runtime parameter limits how many items will be scanned after using the best index.
// If more than this is needed, an error will be returned.
func (client *Client) ListStreamQueryItems(stream, keys, publishers []string, verbose bool) (Response, error) {

	q, err := getQuery(keys, publishers)
	if err != nil {
		return nil, err
	}

	return client.Post(
		map[string]interface{}{
			"jsonrpc": "1.0",
			"id":      CONST_ID,
			"method":  "liststreamqueryitems",
			"params": []interface{}{
				stream,
				q,
				verbose,
			},
		},
	)
}

func getQuery(keys, publishers []string) (map[string]interface{}, error) {
	lenKeys := len(keys)
	lenPublishers := len(publishers)
	if lenKeys == 0 && lenPublishers == 0 {
		return nil,  errors.New("either keys or publishers slices must contain at least one entry")
	}
	q := make(map[string]interface{})
	if lenKeys == 1 {
		q["key"] = keys[0]
	} else if lenKeys > 1 {
		q["keys"] = keys
	}
	if lenPublishers == 1 {
		q["publisher"] = publishers[0]
	} else if lenPublishers > 1 {
		q["publishers"] = publishers
	}
	return q, nil
}
