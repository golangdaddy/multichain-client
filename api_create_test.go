package multichain

import (
	"testing"
)

func TestCreate(t *testing.T) {

	_, err := client.Create("stream", "testStream", nil)
	if err != nil {
		t.Fail()
	}

}
