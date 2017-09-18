package multichain

import (
	"testing"
)

func TestCreate(t *testing.T) {

	_, err := client.GetInfo()
	if err != nil {
		t.Fail()
	}

}
