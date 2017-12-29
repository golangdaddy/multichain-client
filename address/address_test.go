package address

import (
    "fmt"
    "testing"
)

const (
    CONST_BCRYPT_DIFF = 2000
)

func TestAddress(t *testing.T) {

    Configure(&Config{
        PrivateKeyVersion: "8025B89E",
        AddressPubkeyhashVersion: "00AFEA21",
        AddressChecksumValue: "953ABC69",
    })

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

        if keyPair.Public != "1Yu2BuptuZSiBWfr2Qy4aic6qEVnwPWrjddvYh" {
            t.Error("INVALID PUBLIC ADDRESSS GENERATED")
        }

    })

    Configure(&Config{
        PrivateKeyVersion: "8025B89E",
        AddressPubkeyhashVersion: "00AFEA21",
        AddressChecksumValue: "7B7AEF76",
    })

    t.Run("Test private key wif generation", func (t *testing.T) {

        keyPair := DebugKeyPair()

        if keyPair.Private != "VEEWgYhDhqWnNnDCXXjirJYXGDFPjH1B8v6hmcnj1kLXrkpxArmz7xXw" {
            t.Error("INVALID PRIVATE ADDRESSS GENERATED")
        }

    })

}
