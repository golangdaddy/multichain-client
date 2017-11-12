package address

import (
    "fmt"
    "testing"
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

        keyPair := DebugKeyPair()

        if keyPair.Public != "1Yu2BuptuZSiBWfr2Qy4aic6qEVnwPWrdkHPEc" {
            t.Error("INVALID PUBLIC ADDRESSS GENERATED")
        }

    })

    Configure("8025B89E", "00AFEA21", "7B7AEF76")

    t.Run("Test private key wif generation", func (t *testing.T) {

        keyPair := DebugKeyPair()

        if keyPair.Private != "VEEWgYhDhqWnNnDCXXjirJYXGDFPjH1B8v6hmcnj1kLXrkpxArmz7xXw" {
            t.Error("INVALID PRIVATE ADDRESSS GENERATED")
        }

    })

}
