package multichain

import (
	"fmt"
	"testing"
)

func TestListStreams(t *testing.T) {

	x, err := client.ListStreams("", 0, 0, true)
	if err != nil {
		t.Fail()
	}

	fmt.Println(x)
}
