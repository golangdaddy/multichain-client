package address

import (
	"crypto/sha512"
	"fmt"
	"time"

	"github.com/tyler-smith/go-bip32"
)

type KeyPair struct {
	Type    string
	Index   int
	Public  string
	Private string
}

func KeyFromSeed(input []byte, difficulty, index int) (*bip32.Key, *bip32.Key, error) {

	if !configued {
		panic(CONST_UNCONFIGURED)
	}

	t := time.Now()

	for x := 0; x < difficulty; x++ {

		h := sha512.New512_256()
		h.Write(input)
		input = append(input, h.Sum(nil)...)

	}

	h := sha512.New512_256()
	h.Write(input)
	input = h.Sum(nil)

	fmt.Println("INPUT:", input)

	fmt.Printf("key lengthening difficulty %v elaspsed: %v\n", difficulty, time.Since(t))

	masterKey, err := bip32.NewMasterKey(input)
	if err != nil {
		return nil, nil, err
	}

	childKey, err := masterKey.NewChildKey(uint32(index))
	if err != nil {
		return nil, nil, err
	}

	return masterKey, childKey, nil
}

func MultiChainWallet(seed []byte, difficulty, index int) (*KeyPair, error) {

	if !configued {
		panic(CONST_UNCONFIGURED)
	}

	_, key, err := KeyFromSeed(seed, difficulty, index)
	if err != nil {
		return nil, err
	}

	publicKey, err := MultiChainAddress(key.PublicKey().Key)
	if err != nil {
		return nil, err
	}

	keyPair := &KeyPair{
		Type:    "MultiChain",
		Index:   index,
		Public:  publicKey,
		Private: MultiChainWIF(key.Key),
	}

	return keyPair, nil
}
