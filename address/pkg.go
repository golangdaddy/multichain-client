package address

import (
    "crypto/sha256"
    "encoding/hex"
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
