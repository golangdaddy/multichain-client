package address

import (
    "bytes"
    "crypto/sha256"
    "encoding/hex"
    "github.com/mr-tron/base58/base58"
    "golang.org/x/crypto/ripemd160"
)

const (
    CONST_ADDRESS_PUBKEYHASH_VERSION = "00AFEA21"
    CONST_ADDRESS_CHECKSUM_VALUE = "953ABC69"
)

var address_pubkeyhash_version []byte
var address_checksum_value []byte

func init() {
    address_pubkeyhash_version, _ = hex.DecodeString(CONST_ADDRESS_PUBKEYHASH_VERSION)
    address_checksum_value, _ = hex.DecodeString(CONST_ADDRESS_CHECKSUM_VALUE)
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

func wif(key []byte) string {
    stage1 := bytes.Join(
		[][]byte{
			[]byte{byte(0x80)},
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
