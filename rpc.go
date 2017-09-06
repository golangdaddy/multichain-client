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

func (client *Client) post(msg interface{}) (map[string]interface{}, error) {

	request, err := sling.New().Post(client.endpoint).BodyJSON(msg).Request()

	resp, err := client.httpClient.Do(request)
	if err != nil {
		panic(err)
	}

	request.Header.Add("Authorization", "Basic " + client.credentials)

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	m := map[string]interface{}{}

	err = json.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (client *Client) GetInfo() (map[string]interface{}, error) {

	msg := map[string]interface{}{
		"jsonrpc": "1.0",
		"id": "curltest",
		"method": "getinfo",
		"params": []interface{}{},
	}

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	return obj, err
}

func (client *Client) SendAssetToAddress(accountAddress, assetName string, value float64) (map[string]interface{}, error) {

	msg := map[string]interface{}{
		"jsonrpc": "1.0",
		"id": "curltest",
		"method": "sendassettoaddress",
		"params": []interface{}{
			accountAddress,
			assetName,
			fmt.Sprintf("%d", value),
		},
	}

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	return obj, err
}
