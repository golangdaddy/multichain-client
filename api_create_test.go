package multichain

import (
	"testing"
)

func TestCreate(t *testing.T) {

	client = NewClient(
		"",
		"localhost",
		"multichainrpc",
		"12345678",
		80,
	)
	_, err := client.Create("stream", "testStream", nil)
	if err != nil {
		t.Fail()
	}

}
