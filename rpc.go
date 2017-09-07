package multichain

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"encoding/base64"
	//
	"github.com/dghubble/sling"
)

const (
	CONST_ID = "multichain-client"
)

type Response map[string]interface{}

func (r Response) Result() interface{} {
	return r["result"]
}

type Client struct {
	httpClient *http.Client
	endpoint string
	credentials string
}

func NewClient(host, port, username, password string) *Client {

	credentials := username + ":" + password

	return &Client{
		httpClient: &http.Client{},
		endpoint: fmt.Sprintf("http://%s:%s", host, port),
		credentials: base64.StdEncoding.EncodeToString([]byte(credentials)),
	}
}

func (client *Client) post(msg interface{}) (Response, error) {

	request, err := sling.New().Post(client.endpoint).BodyJSON(msg).Request()

	request.Header.Add("Authorization", "Basic " + client.credentials)

	resp, err := client.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	obj := make(Response)

	if err := json.Unmarshal(b, &obj); err != nil {
		fmt.Println(string(b))
		return nil, err
	}

	return obj, nil
}

func (client *Client) GetInfo() (Response, error) {

	msg := map[string]interface{}{
		"jsonrpc": "1.0",
		"id": CONST_ID,
		"method": "getinfo",
		"params": []interface{}{},
	}

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (client *Client) GetNewAddress() (Response, error) {

	msg := map[string]interface{}{
		"jsonrpc": "1.0",
		"id": CONST_ID,
		"method": "getnewaddress",
		"params": []interface{}{},
	}

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (client *Client) Grant(addresses, permissions []string) (Response, error) {

	msg := map[string]interface{}{
		"jsonrpc": "1.0",
		"id": CONST_ID,
		"method": "grant",
		"params": []interface{}{strings.Join(addresses, ","), strings.Join(permissions, ",")},
	}

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	return obj, nil
}


func (client *Client) GetAddressBalances(address string) (Response, error) {

	msg := map[string]interface{}{
		"jsonrpc": "1.0",
		"id": CONST_ID,
		"method": "getaddressbalances",
		"params": []interface{}{address},
	}

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (client *Client) CreateKeypair() (*AddressKeyPair, error) {

	msg := map[string]interface{}{
		"jsonrpc": "1.0",
		"id": CONST_ID,
		"method": "createkeypairs",
		"params": []interface{}{},
	}

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	array := obj["result"].([]interface{})

	result := array[0].(map[string]interface{})

	addressKeyPair := &AddressKeyPair{
		Address: result["address"].(string),
		PubKey: result["pubkey"].(string),
		PrivKey: result["privkey"].(string),
	}

	return addressKeyPair, nil
}

func (client *Client) SendAssetToAddress(accountAddress, assetName string, value float64) (Response, error) {

	msg := map[string]interface{}{
		"jsonrpc": "1.0",
		"id": CONST_ID,
		"method": "sendassettoaddress",
		"params": []interface{}{
			accountAddress,
			assetName,
			value,
		},
	}

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (client *Client) IssueMore(accountAddress, assetName string, value float64) (Response, error) {

	msg := map[string]interface{}{
		"jsonrpc": "1.0",
		"id": CONST_ID,
		"method": "sendassettoaddress",
		"params": []interface{}{
			accountAddress,
			assetName,
			value,
		},
	}

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
