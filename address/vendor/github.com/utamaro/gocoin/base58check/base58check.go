package base58check

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"math/big"

	"github.com/TankerApp/gocoin/base58check/base58"
)

func Encode(prefix byte, byteData []byte) string {
	length := len(byteData) + 1
	encoded := make([]byte, length)
	encoded[0] = prefix
	copy(encoded[1:], byteData)

	//Perform SHA-256 twice
	hash := sha256.Sum256(encoded)
	hash2 := sha256.Sum256(hash[:])

	//First 4 bytes if this double-sha'd byte array is the checksum
	checksum := hash2[0:4]

	//Append this checksum to the input bytes
	encodedChecksum := append(encoded, checksum...)

	//Convert this checksum'd version to a big Int
	bigIntEncodedChecksum := big.NewInt(0)
	bigIntEncodedChecksum.SetBytes(encodedChecksum)

	//Encode the big int checksum'd version into a Base58Checked string
	base58EncodedChecksum := string(base58.EncodeBig(nil, bigIntEncodedChecksum))

	//Now for each zero byte we counted above we need to prepend a 1 to our
	//base58 encoded string. The rational behind this is that base58 removes 0's (0x00).
	//So bitcoin demands we add leading 0s back on as 1s.
	var buffer bytes.Buffer

	//base58 alone is not enough. We need to first count each of the zero bytes
	//which are at the beginning of the encodedCheckSum

	for _, v := range encodedChecksum {
		if v == 0 {
			buffer.WriteByte('1')
		} else {
			break
		}
	}

	buffer.WriteString(base58EncodedChecksum)

	return buffer.String()
}

func Decode(value string) ([]byte, bool, error) {
	publicKeyInt, err := base58.DecodeToBig([]byte(value))
	if err != nil {
		return nil, false, err
	}

	encodedChecksum := publicKeyInt.Bytes()

	encoded := encodedChecksum[:len(encodedChecksum)-4]
	cksum := encodedChecksum[len(encodedChecksum)-4:]

	var buffer bytes.Buffer
	for _, v := range value {
		if v == '1' {
			buffer.WriteByte(0)
		} else {
			break
		}
	}

	buffer.Write(encoded)

	result := buffer.Bytes()

	//Perform SHA-256 twice
	hash := sha256.Sum256(result)
	hash2 := sha256.Sum256(hash[:])

	if !bytes.Equal(hash2[:4], cksum) {
		log.Println(value, "warn:", "checksum not matched", "embeded cksum:", hex.EncodeToString(cksum), "cksum:", hex.EncodeToString(hash2[:4]))
	}

	if value[0] == 'K' || value[0] == 'L' || value[0] == 'c' {
		return result[0 : len(result)-1], true, err
	}
	return result, false, err
}
