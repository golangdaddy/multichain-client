package address

import (
    "bytes"
    "github.com/mr-tron/base58/base58"
)

func MultiChainWIF(key []byte) string {

    stage1 := bytes.Join(
        [][]byte{
            key,
            []byte{0x01},
        },
        nil,
    )

    x := int(33 / len(private_key_version))

    var index int
    extended := []byte{}
    for index = 0; index < len(private_key_version); index++ {
        extended = append(extended, private_key_version[ index : (index + 1) ]...)
        extended = append(extended, stage1[ (x * index) : (index * x) + x ]...)
    }

    extended = append(extended, stage1[ index * x : ]...)

    stage2 := sha(
        sha(
            extended,
        ),
    )[:4]

    checksum := make([]byte, 4)
    safeXORBytes(checksum, stage2, address_checksum_value)

    stage3 := bytes.Join(
        [][]byte{
            extended,
            checksum,
        },
        nil,
    )

    return base58.FastBase58Encoding(stage3)
}

func BitcoinWIF(key []byte) string {

    stage1 := bytes.Join(
		[][]byte{
			[]byte{private_key_version[0]},
			key,
		},
		nil,
	)

	stage2 := sha(
		sha(
			stage1,
		),
	)[:4]

	stage3 := bytes.Join(
		[][]byte{
			stage1,
			stage2,
		},
		nil,
	)

    return base58.FastBase58Encoding(stage3)
}
