package multichain

import (
	"fmt"
	"net/http"
	"encoding/base64"
	//
	"gitlab.com/TheDarkula/jsonrouter/http"
)

type Client struct {
	http *httpclient.Client
	host string
	credentials string
	timeout []int
	urlfetch bool
	debug bool
}

func NewClient(httpClient *http.Client, host, username, password string, port int) *Client {

	credentials := username + ":" + password

	return &Client{
		http: httpclient.NewClient(httpClient),
		host: fmt.Sprintf("http://%s:%d", host, port),
		credentials: base64.StdEncoding.EncodeToString([]byte(credentials)),
	}
}

func (client *Client) HttpClient() *httpclient.Client {
	return client.http
}
