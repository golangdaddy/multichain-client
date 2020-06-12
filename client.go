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

func (c *Client) String() string {
	return fmt.Sprintf("{chainName: %s, host: %s, username: ***, password: ***, timeout: %p, debug: %t}",
		c.chainName, c.host, c.timeout, c.debug)
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
