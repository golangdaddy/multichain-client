package multichain

import (
	"fmt"
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

func (client *Client) post(msg, dst interface{}) error {

	request, err := sling.New().Post(client.endpoint).BodyJSON(msg).Request()

	request.Header.Add("Authorization", "Basic " + client.credentials)

	resp, err := client.httpClient.Do(request)
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, dst)
	if err != nil {
		fmt.Println(string(b))
		return err
	}

	return nil
}

func (client *Client) GetInfo() (map[string]interface{}, error) {

	msg := map[string]interface{}{
		"jsonrpc": "1.0",
		"id": CONST_ID,
		"method": "getinfo",
		"params": []interface{}{},
	}

	obj := map[string]interface{}{}

	if err := client.post(msg, &obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (client *Client) GetNewAddress() (map[string]interface{}, error) {

	msg := map[string]interface{}{
		"jsonrpc": "1.0",
		"id": CONST_ID,
		"method": "getnewaddress",
		"params": []interface{}{},
	}

	obj := map[string]interface{}{}

	if err := client.post(msg, &obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (client *Client) CreateKeypairs() (*AddressKeyPair, error) {

	msg := map[string]interface{}{
		"jsonrpc": "1.0",
		"id": CONST_ID,
		"method": "createkeypairs",
		"params": []interface{}{},
	}

	obj := &AddressKeyPair{}

	if err := client.post(msg, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (client *Client) SendAssetToAddress(accountAddress, assetName string, value float64) (map[string]interface{}, error) {

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

	obj := map[string]interface{}{}

	if err := client.post(msg, &obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (client *Client) IssueMore(accountAddress, assetName string, value float64) (map[string]interface{}, error) {

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

	obj := map[string]interface{}{}

	if err := client.post(msg, &obj); err != nil {
		return nil, err
	}

	return obj, nil
}
