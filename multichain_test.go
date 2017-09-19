package multichain

import (
	"fmt"
	"flag"
)

var client *Client

func init() {

	host := flag.String("host", "localhost", "is a string")
	port := flag.String("port", "80", "is a string")
	username := flag.String("username", "multichainrpc", "is a string")
	password := flag.String("password", "12345678", "is a string")

	flag.Parse()

	client = NewClient(
		*host,
		*port,
		*username,
		*password,
	)

	fmt.Println(client.debug())
}
