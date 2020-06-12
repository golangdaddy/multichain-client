package multichain

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	Resty *resty.Client
	chainName string
	host string
	username string
	password string
	timeout []int
	debug bool
}

func NewClient(chainName, host, username, password string, port int) *Client {
	return &Client{
		Resty: resty.New(),
		chainName: chainName,
		host: fmt.Sprintf("http://%s:%d", host, port),
		username: username,
		password: password,
	}
}
