package params

import (
    "testing"
)

func TestPackage(t *testing.T) {

    _, err := Open("params.dat")
    if err != nil {
        panic(err)
    }
}
