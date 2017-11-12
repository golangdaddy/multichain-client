package address

import (
    "fmt"
    "testing"
    "encoding/hex"
)

const (
    CONST_BCRYPT_DIFF = 10
)

func TestAddress(t *testing.T) {

    Configure("8025B89E", "00AFEA21", "953ABC69")

    t.Run("Test MultiChain wallet generation", func (t *testing.T) {

        seed := []byte("seed")

        keyPair, err := MultiChainWallet(seed, CONST_BCRYPT_DIFF, 0)
        if err != nil {
            t.Error(err)
        }

        fmt.Println(keyPair)

    })

    t.Run("Test public address generation", func (t *testing.T) {

        b, _ := hex.DecodeString("0284E5235E299AF81EBE1653AC5F06B60E13A3A81F918018CBD10CE695095B3E24")

        pubAddress, err := MultiChainAddress(b)
        if err != nil {
            t.Error(err)
        }

        fmt.Println(pubAddress)

        if pubAddress != "1Yu2BuptuZSiBWfr2Qy4aic6qEVnwPWrdkHPEc" {
            t.Error("INVALID PUBLIC ADDRESSS GENERATED")
        }

    })

    Configure("8025B89E", "00AFEA21", "7B7AEF76")

    t.Run("Test private key wif generation", func (t *testing.T) {

        b, _ := hex.DecodeString("B69CA8FFAE36F11AD445625E35BF6AC57D6642DDBE470DD3E7934291B2000D78")

        wif := MultiChainWIF(b)

        fmt.Println(wif)

        if wif != "VEEWgYhDhqWnNnDCXXjirJYXGDFPjH1B8v6hmcnj1kLXrkpxArmz7xXw" {
            t.Error("INVALID PRIVATE ADDRESSS GENERATED")
        }

    })

}
