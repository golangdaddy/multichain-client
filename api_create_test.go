package multichain

import (
	"testing"
)

func TestCreate(t *testing.T) {

	_, err := client.Create("stream", "testStream", true)
	if err != nil {
		t.Fail()
	}

}
