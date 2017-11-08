package address

import (
    "fmt"
    "bytes"
    "github.com/mr-tron/base58/base58"
    "github.com/tyler-smith/go-bip32"
)

type KeyPair struct {
    Public string
    Private string
}

func NewWallet(seed []byte) (*KeyPair, error) {

    masterKey, err := bip32.NewMasterKey(seed)
    if err != nil {
        return nil, err
    }

    key, err := masterKey.NewChildKey(0)
    if err != nil {
        return nil, err
    }

	stage1 := bytes.Join(
		[][]byte{
			[]byte{byte(0x80)},
			key.Key,
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

    publicKey, err := MultiChain(fmt.Sprintf("%X", key.PublicKey().Key))
    if err != nil {
        return nil, err
    }

    keyPair := &KeyPair{
        Public: publicKey,
        Private: base58.FastBase58Encoding(stage3),
    }

    return keyPair, nil
}
