package params

import (
    "testing"
)

func TestPackage(t *testing.T) {

    params, _, err := Open("params.dat")
    if err != nil {
        panic(err)
    }

    params.Bool("only-accept-std-txs")

    params.Int("max-std-op-return-size")

    params.String("chain-name")

    params.Float64("test-float")
}
