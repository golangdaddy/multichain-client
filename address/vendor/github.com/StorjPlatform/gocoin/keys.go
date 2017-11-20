/*
 * Copyright (c) 2015, Shinya Yagyu
 * All rights reserved.
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *
 * 1. Redistributions of source code must retain the above copyright notice,
 *    this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright notice,
 *    this list of conditions and the following disclaimer in the documentation
 *    and/or other materials provided with the distribution.
 * 3. Neither the name of the copyright holder nor the names of its
 *    contributors may be used to endorse or promote products derived from this
 *    software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
 * AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
 * LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
 * CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
 * SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
 * INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
 * CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
 * ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 * POSSIBILITY OF SUCH DAMAGE.
 *
 * see LICENSE file for the original license:
 */

package gocoin

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"log"
	"os"

	"github.com/TankerApp/gocoin/base58check"
	"github.com/TankerApp/gocoin/btcec"
	"golang.org/x/crypto/ripemd160"
)

var logging *log.Logger

//initialize logger.
func init() {
	logging = log.New(os.Stdout, "",
		log.Ldate|log.Ltime|log.Lshortfile)
}

//SetLogger sets logger this module uses.
func SetLogger(logger *log.Logger) {
	logging = logger
}

var flagTestnet bool

//PublicKey represents public key for bitcoin
type PublicKey struct {
	key          *btcec.PublicKey
	isTestnet    bool
	isCompressed bool
}

//PrivateKey represents private key for bitcoin
type PrivateKey struct {
	key       *btcec.PrivateKey
	isTestnet bool
}

//Key includes PublicKey and PrivateKey.
type Key struct {
	Pub  *PublicKey
	Priv *PrivateKey
}

//Multiple steps follow. I've encapsulated this functionality into
//the base58CheckEncode method because a similar process is used to generate
//a readable public key as well. Here are the steps for the private key.

//First generate "extended" private key from private key
//The difference between a private key and an extended
//private key is this prefix, which determines the
//network the key belongs to (real btc network, or test network)

//EF is the testnet prefix
//80 is the mainnet prefix

//Perform SHA-256 on the extended key twice
//First 4 bytes if this double-sha'd byte array are the checksum
//Append this checksum to the extended private key
//Convert the extended private key to a big Int
//Encoded the big int extended private key into a Base58Checked string

//There is also a prefix on the public key
//This is known as the Network ID Byte, or the version byte
//6f is the testnet prefix
//00 is the mainnet prefix

//GetPublicKey returns PublicKey struct using public key hex string.
func GetPublicKey(pubKeyByte []byte, isTestnet bool) (*PublicKey, error) {
	secp256k1 := btcec.S256()
	key, err := btcec.ParsePubKey(pubKeyByte, secp256k1)
	if err != nil {
		return nil, err
	}
	return &PublicKey{key: key, isTestnet: isTestnet}, nil
}

//GetKeyFromWIF gets PublicKey and PrivateKey from private key of WIF format.
func GetKeyFromWIF(wif string) (*Key, error) {
	secp256k1 := btcec.S256()
	privateKeyBytes, isCmpressed, err := base58check.Decode(wif)
	if err != nil {
		return nil, err
	}

	pub := PublicKey{}
	priv := PrivateKey{}
	key := Key{
		Pub:  &pub,
		Priv: &priv,
	}
	switch privateKeyBytes[0] {
	case 0xef:
		pub.isTestnet = true
		priv.isTestnet = true
	case 0x80:
		pub.isTestnet = false
		priv.isTestnet = false
	default:
		return nil, errors.New("cannot determin net param from private key")
	}
	pub.isCompressed = isCmpressed

	//Get the raw public
	priv.key, pub.key = btcec.PrivKeyFromBytes(secp256k1, privateKeyBytes[1:])

	return &key, nil

}

//GenerateKey generates random PublicKey and PrivateKey.
func GenerateKey(flagTestnet bool) (*Key, error) {
	seed := make([]byte, 32)
	_, err := rand.Read(seed)
	if err != nil {
		return nil, err
	}
	s256 := btcec.S256()

	private := PrivateKey{}
	private.isTestnet = flagTestnet
	public := PublicKey{}
	public.isTestnet = flagTestnet
	private.key, public.key = btcec.PrivKeyFromBytes(s256, seed)
	key := Key{
		Pub:  &public,
		Priv: &private,
	}

	//Print the keys
	logging.Println("Your private key in WIF is")
	logging.Println(private.GetWIFAddress())

	logging.Println("Your address is")
	logging.Println(public.GetAddress())

	return &key, nil
}

//Sign sign data.
func (priv *PrivateKey) Sign(hash []byte) ([]byte, error) {
	sig, err := priv.key.Sign(hash)
	if err != nil {
		return nil, err
	}
	return sig.Serialize(), nil
}

//SignMessage sign using bitcoin sign struct
func (key *Key) SignMessage(hash []byte) ([]byte, error) {
	msg := make([]byte, 0)
	msg = append(msg, []byte("\x18Bitcoin Signed Message:\n")...)
	msg = append(msg, []byte{byte(len(hash))}...)
	msg = append(msg, hash...)
	h := sha256.Sum256(msg)
	hh := sha256.Sum256(h[:])
	s256 := btcec.S256()

	sig, err := btcec.SignCompact(s256, key.Priv.key, hh[:], key.Pub.isCompressed)
	if err != nil {
		return nil, err
	}
	return sig, nil
}

//GetWIFAddress returns WIF format string from PrivateKey
func (priv *PrivateKey) GetWIFAddress() string {
	var privateKeyPrefix byte
	if priv.isTestnet {
		privateKeyPrefix = 0xEF
	} else {
		privateKeyPrefix = 0x80
	}

	return base58check.Encode(privateKeyPrefix, priv.key.Serialize())
}

//GetAddress returns bitcoin address from PublicKey
func (pub *PublicKey) GetAddress() (string, []byte) {
	var publicKeyPrefix byte

	if pub.isTestnet {
		publicKeyPrefix = 0x6F
	} else {
		publicKeyPrefix = 0x00
	}

	//Next we get a sha256 hash of the public key generated
	//via ECDSA, and then get a ripemd160 hash of the sha256 hash.
	var shadPublicKeyBytes [32]byte
	if pub.isCompressed {
		shadPublicKeyBytes = sha256.Sum256(pub.key.SerializeCompressed())

	} else {
		shadPublicKeyBytes = sha256.Sum256(pub.key.SerializeUncompressed())
	}

	ripeHash := ripemd160.New()
	ripeHash.Write(shadPublicKeyBytes[:])
	ripeHashedBytes := ripeHash.Sum(nil)

	publicKeyEncoded := base58check.Encode(publicKeyPrefix, ripeHashedBytes)
	return publicKeyEncoded, ripeHashedBytes
}

//IsTestnet returns true if addr is for testnet.
func IsTestnet(addr string) (bool, error) {
	bytes, _, err := base58check.Decode(addr)
	if err != nil {
		return false, err
	}

	switch bytes[0] {
	case 0x6f:
		return true, nil
	case 0x00:
		return false, nil
	default:
		return false, errors.New("invalid address")
	}

}
