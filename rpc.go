package multichain

import (
	"fmt"
	//"strings"
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
