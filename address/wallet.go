package address

import (
//    "fmt"
//    "time"
    "github.com/tyler-smith/go-bip32"
    "golang.org/x/crypto/bcrypt"
)

type KeyPair struct {
    Type string
    Index int
    Public string
    Private string
}

func KeyFromSeed(input []byte, index int) (*bip32.Key, *bip32.Key, error) {

//  t := time.Now()

    // 90 ms of protection
    seed, err := bcrypt.GenerateFromPassword(input, 10)
    if err != nil {
        return nil, nil, err
    }

//  fmt.Println(time.Since(t))

    masterKey, err := bip32.NewMasterKey(seed)
    if err != nil {
        return nil, nil, err
    }

    childKey, err := masterKey.NewChildKey(uint32(index))
    if err != nil {
        return nil, nil, err
    }

    return masterKey, childKey, nil
}

func MultiChainWallet(seed []byte, index int) (*KeyPair, error) {

    _, key, err := KeyFromSeed(seed, index)
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

    _, key, err := KeyFromSeed(seed, index)
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
