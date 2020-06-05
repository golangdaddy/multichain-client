package multichain

import (
	"fmt"
	"encoding/base64"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	Resty *resty.Client
	chainName string
	host string
	credentials string
	timeout []int
	debug bool
}

func NewClient(chainName, host, username, password string, port int) *Client {

	credentials := username + ":" + password

	return &Client{
		Resty: resty.New(),
		chainName: chainName,
		host: fmt.Sprintf("http://%s:%d", host, port),
		credentials: base64.StdEncoding.EncodeToString([]byte(credentials)),
	}
}
