package address

import (
    "fmt"
    "testing"
    "encoding/hex"
)

func TestAddress(t *testing.T) {

    t.Run("Test address generation", func (t *testing.T) {

        b, _ := hex.DecodeString("0284E5235E299AF81EBE1653AC5F06B60E13A3A81F918018CBD10CE695095B3E24")

        address, err := MultiChainAddress(b)
        if err != nil {
            t.Error(err)
        }

        if address != "1Yu2BuptuZSiBWfr2Qy4aic6qEVnwPWrdkHPEc" {
            t.Error("INVALID PUBLIC ADDRESSS GENERATED")
        }

    })

    t.Run("Test wallet generation", func (t *testing.T) {

        keyPair, err := MultiChainWallet([]byte("seed"), 0)
        if err != nil {
            t.Error(err)
        }

        fmt.Println(keyPair)

        keyPair, err = BitcoinWallet([]byte("seed"), 0)
        if err != nil {
            t.Error(err)
        }

        fmt.Println(keyPair)
    })

}
