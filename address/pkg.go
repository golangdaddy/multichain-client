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

func Configure(privateKeyVersion, addressPubkeyhashVersion, addressChecksumValue string) {
    private_key_version, _ = hex.DecodeString(privateKeyVersion)
    address_pubkeyhash_version, _ = hex.DecodeString(addressPubkeyhashVersion)
    address_checksum_value, _ = hex.DecodeString(addressChecksumValue)
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
