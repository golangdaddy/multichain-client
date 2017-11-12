package address

import (
//    "fmt"
    "github.com/utamaro/gocoin"
    "github.com/mr-tron/base58/base58"
)

func BitcoinAddress(input []byte) (string, error) {

    if !configued { panic(CONST_UNCONFIGURED) }

    privateKey := BitcoinWIF(input)

    k, err := gocoin.GetKeyFromWIF(privateKey)
    if err != nil {
        return "", err
    }

    publicKey, _ := k.Pub.GetAddress()

    return publicKey, nil
}

func MultiChainAddress(input []byte) (string, error) {

    if !configued { panic(CONST_UNCONFIGURED) }

    x := int(20 / len(address_pubkeyhash_version))

    b := ripemd(sha(input))

    extended := []byte{}
    for index := 0; index < len(address_pubkeyhash_version); index++ {
        extended = append(extended, address_pubkeyhash_version[ index : (index + 1) ]...)
        extended = append(extended, b[ (x * index) : (index * x) + x ]...)
    }

    b = make([]byte, 4)
    safeXORBytes(b, address_checksum_value, sha(sha(extended))[:4])

    b = append(extended, b...)

    return string(base58.FastBase58Encoding(b)), nil
}
