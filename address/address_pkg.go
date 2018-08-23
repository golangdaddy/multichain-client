package address

import (
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/crypto/ripemd160"
)

const (
	CONST_UNCONFIGURED = "CALL THE CONFIGURE METHOD WITH YOUR BLOCKCHAIN PARAMS FIRST"
)

var configued bool

var private_key_version []byte
var address_pubkeyhash_version []byte
var address_checksum_value []byte

type Config struct {
	PrivateKeyVersion        string
	AddressPubkeyhashVersion string
	AddressChecksumValue     string
}

func Configure(config *Config) {
	private_key_version, _ = hex.DecodeString(config.PrivateKeyVersion)
	address_pubkeyhash_version, _ = hex.DecodeString(config.AddressPubkeyhashVersion)
	address_checksum_value, _ = hex.DecodeString(config.AddressChecksumValue)
	configued = true
}

func ripemd(b []byte) []byte {
	h := ripemd160.New()
	h.Write(b)
	return h.Sum(nil)
}

func sha(b []byte) []byte {
	c := sha256.Sum256(b)
	return c[:]
}

func safeXORBytes(dst, a, b []byte) int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	for i := 0; i < n; i++ {
		dst[i] = a[i] ^ b[i]
	}
	return n
}

func DebugKeyPair() *KeyPair {

	Configure(&Config{
		PrivateKeyVersion:        "8025B89E",
		AddressPubkeyhashVersion: "00AFEA21",
		AddressChecksumValue:     "7B7AEF76",
	})

	b, _ := hex.DecodeString("0284E5235E299AF81EBE1653AC5F06B60E13A3A81F918018CBD10CE695095B3E24")

	pubAddress, err := MultiChainAddress(b)
	if err != nil {
		panic(err)
	}

	b, _ = hex.DecodeString("B69CA8FFAE36F11AD445625E35BF6AC57D6642DDBE470DD3E7934291B2000D78")

	return &KeyPair{
		Type:    "MultiChainDebug",
		Private: MultiChainWIF(b),
		Public:  pubAddress,
	}
}
