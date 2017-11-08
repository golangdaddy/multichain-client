package address

import (
    "github.com/tyler-smith/go-bip32"
)

type KeyPair struct {
    Type string
    Index int
    Public string
    Private string
}

func MultiChainWallet(seed []byte, index int) (*KeyPair, error) {

    masterKey, err := bip32.NewMasterKey(seed)
    if err != nil {
        return nil, err
    }

    key, err := masterKey.NewChildKey(uint32(index))
    if err != nil {
        return nil, err
    }

    publicKey, err := MultiChainAddress(key.PublicKey().Key)
    if err != nil {
        return nil, err
    }

    keyPair := &KeyPair{
        Type: "MultiChain",
        Index: index,
        Public: publicKey,
        Private: wif(key.Key),
    }

    return keyPair, nil
}

func BitcoinWallet(seed []byte, index int) (*KeyPair, error) {

    masterKey, err := bip32.NewMasterKey(seed)
    if err != nil {
        return nil, err
    }

    key, err := masterKey.NewChildKey(uint32(index))
    if err != nil {
        return nil, err
    }

    publicKey, err := BitcoinAddress(key.PublicKey().Key)
    if err != nil {
        return nil, err
    }

    keyPair := &KeyPair{
        Type: "Bitcoin",
        Index: index,
        Public: publicKey,
        Private: wif(key.Key),
    }

    return keyPair, nil
}
