package multichain

import (
	"fmt"
	"errors"
)

const (
	CONST_ID = "multichain-client"
)

type Response map[string]interface{}

func (r Response) Result() interface{} {
	return r["result"]
}

func (client *Client) IsDebugMode() bool {
	return client.debug
}

func (client *Client) DebugMode() *Client {
	client.debug = true
	return client
}

func (client *Client) msg(params []interface{}) map[string]interface{} {
	return map[string]interface{}{
		"jsonrpc": "1.0",
		"id": CONST_ID,
		"params": params,
	}
}

func (client *Client) Command(method string, params []interface{}) map[string]interface{} {

	msg := client.msg(params)
	msg["method"] = fmt.Sprintf("%s", method)

	return msg
}

func (client *Client) Post(msg interface{}) (Response, error) {

	fmt.Println("POSTING: "+client.host, msg)

	obj := make(Response)
	if _, err := client.Resty.R().SetBody(msg).SetResult(&obj).SetBasicAuth(client.username, client.password).Post(client.host); err != nil {
		return nil, err
	}

	if obj == nil || obj["error"] != nil {
		e := obj["error"].(map[string]interface{})
		var s string
		m, ok := msg.(map[string]interface{})
		if ok {
			s = fmt.Sprintf("multichaind - '%s': %s", m["method"], e["message"].(string))
		} else {
			s = fmt.Sprintf("multichaind - %s", e["message"].(string))
		}
		return nil, errors.New(s)
	}

	return obj, nil
}
