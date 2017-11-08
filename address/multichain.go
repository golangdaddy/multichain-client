package address

import (
//    "fmt"
    "encoding/hex"
    "github.com/mr-tron/base58/base58"
)

func MultiChain(one string) (string, error) {

    x := int(20 / len(address_pubkeyhash_version))
//    fmt.Println("--", len(address_pubkeyhash_version), x)

//    fmt.Println("1:", one)

    two, err := hex.DecodeString(one)
    if err != nil {
        panic(err)
    }

    three := sha(two)
//    fmt.Printf("3: %X\n", three)

    four := ripemd(three)
//    fmt.Printf("4: %X\n", four)

    five := []byte{}
    for index := 0; index < len(address_pubkeyhash_version); index++ {
        five = append(five, address_pubkeyhash_version[ index : (index + 1) ]...)
        five = append(five, four[ (x * index) : (index * x) + x ]...)
    }
//    fmt.Printf("5: %X\n", five)

    six := sha(five)
//    fmt.Printf("6: %X\n", six)

    seven := sha(six)
//    fmt.Printf("7: %X\n", seven)

    eight := seven[:4]
//    fmt.Printf("8: %X\n", eight)

    nine := make([]byte, 4)
    safeXORBytes(nine, address_checksum_value, eight)
//    fmt.Printf("9: %X\n", nine)

    ten := append(five, nine...)
//    fmt.Printf("10: %X\n", ten)

    encoded := base58.FastBase58Encoding(ten)

    address := string(encoded)

//    fmt.Println("11:", address)

    return address, nil
}
