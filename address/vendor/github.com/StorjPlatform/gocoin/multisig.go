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
 * See LICENSE file for the original license:
 *
 * This file also includes codes from https://github.com/soroushjp/go-bitcoin-multisig
 * copyrighted by Soroush Pour.
 */

package gocoin

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"

	"golang.org/x/crypto/ripemd160"

	"github.com/TankerApp/gocoin/base58check"
	"github.com/TankerApp/gocoin/btcec"
)

//RedeemScript represents Redeem script for M of N multisig transaction.
type RedeemScript struct {
	M          int
	PublicKeys []*PublicKey
	Script     []byte
}

// GetHash creates hash of redeem script.
func (rs *RedeemScript) getHash() []byte {
	shadPublicKeyBytes := sha256.Sum256(rs.Script)
	ripeHash := ripemd160.New()
	ripeHash.Write(shadPublicKeyBytes[:])
	return ripeHash.Sum(nil)
}

// GetAddress creates P2SH multisig addresses.
func (rs *RedeemScript) GetAddress() string {
	//Get P2SH address by base58 encoding with P2SH prefix 0x05
	var prefix byte

	if rs.PublicKeys[0].isTestnet {
		prefix = 0xc4
	} else {
		prefix = 0x5
	}
	return base58check.Encode(prefix, rs.getHash())
}

// createSriptPubkey creates a scriptPubKey for a P2SH transaction given the redeemScript struct.
func (rs *RedeemScript) createSriptPubkey() []byte {
	redeemScriptHash := rs.getHash()
	//P2SH scriptSig format:
	//<OP_HASH160> <Hash160(redeemScript)> <OP_EQUAL>
	var scriptPubKey bytes.Buffer
	scriptPubKey.WriteByte(opHASH160)
	scriptPubKey.WriteByte(byte(len(redeemScriptHash))) //PUSH
	scriptPubKey.Write(redeemScriptHash)
	scriptPubKey.WriteByte(opEQUAL)
	return scriptPubKey.Bytes()
}

// NewRedeemScript creates a M-of-N Multisig redeem script given m, n and n public keys
//and return RedeemScript struct.
func NewRedeemScript(m int, publicKeys []*PublicKey) (*RedeemScript, error) {
	rs := RedeemScript{}
	rs.M = m
	rs.PublicKeys = publicKeys
	n := len(publicKeys)
	//Check we have valid numbers for M and N
	if n < 1 || n > 7 {
		return nil, errors.New("N must be between 1 and 7 (inclusive) for valid, standard P2SH multisig transaction as per Bitcoin protocol.")
	}
	if m < 1 || m > n {
		return nil, errors.New("M must be between 1 and N (inclusive).")
	}
	isTestnet := publicKeys[0].isTestnet
	for _, pk := range publicKeys {
		if pk.isTestnet != isTestnet {
			return nil, errors.New("every public keys must be on same net(testnet or not)")
		}
	}
	//Get OP Code for m and n.
	//81 is OP_1, 82 is OP_2 etc.
	//80 is not a valid OP_Code, so we floor at 81
	mOPCode := op1 + (byte(m) - 1)
	nOPCode := op1 + (byte(n) - 1)
	//Multisig redeemScript format:
	//<OP_m> <A pubkey> <B pubkey> <C pubkey>... <OP_n> OP_CHECKMULTISIG
	var redeemScript bytes.Buffer
	redeemScript.WriteByte(byte(mOPCode)) //m
	for _, pk := range publicKeys {
		publicKey := pk.key.SerializeUncompressed()
		redeemScript.WriteByte(byte(len(publicKey))) //PUSH
		redeemScript.Write(publicKey)                //<pubkey>
	}
	redeemScript.WriteByte(byte(nOPCode)) //n
	redeemScript.WriteByte(byte(opCHECKMULTISIG))
	rs.Script = redeemScript.Bytes()
	return &rs, nil
}

// CreateScriptSig signs a raw transaction with keys.
func (rs *RedeemScript) createScriptSig(rawTransactionHashed []byte, signs [][]byte) ([]byte, error) {

	//Verify that it worked.
	secp256k1 := btcec.S256()
	count := 0
	for i, signature := range signs {
		if signature == nil {
			continue
		}
		count++
		sig, err := btcec.ParseSignature(signature, secp256k1)
		if err != nil {
			return nil, err
		}
		valid := sig.Verify(rawTransactionHashed, rs.PublicKeys[i].key)
		if !valid {
			return nil, fmt.Errorf("number %d of signature is invalid", i)
		}
	}

	if count != rs.M {
		return nil, fmt.Errorf("number of signatures %d must be %d", count, rs.M)
	}

	//redeemScript length. To allow redeemScript > 255 bytes, we use OP_PUSHDATA2 and use two bytes to specify length
	var redeemScriptLengthBytes []byte
	var requiredPUSHDATA byte
	if len(rs.Script) < 255 {
		requiredPUSHDATA = opPUSHDATA1 //OP_PUSHDATA1 specifies next *one byte* will be length to be pushed to stack
		redeemScriptLengthBytes = []byte{byte(len(rs.Script))}
	} else {
		requiredPUSHDATA = opPUSHDATA2 //OP_PUSHDATA2 specifies next *two bytes* will be length to be pushed to stack
		redeemScriptLengthBytes = make([]byte, 2)
		binary.LittleEndian.PutUint16(redeemScriptLengthBytes, uint16(len(rs.Script)))
	}
	//Create scriptSig
	var buffer bytes.Buffer
	buffer.WriteByte(op0) //OP_0 for Multisig off-by-one error
	for _, signature := range signs {
		if signature == nil {
			continue
		}
		buffer.WriteByte(byte(len(signature) + 1)) //PUSH each signature. Add one for hash type byte
		buffer.Write(signature)                    // Signature bytes
		buffer.WriteByte(0x1)                      //hash type
	}
	buffer.WriteByte(requiredPUSHDATA)    //OP_PUSHDATA1 or OP_PUSHDATA2 depending on size of redeemScript
	buffer.Write(redeemScriptLengthBytes) //PUSH redeemScript
	buffer.Write(rs.Script)               //redeemScript
	return buffer.Bytes(), nil
}

func (rs *RedeemScript) getMultisigTX(keys []*Key, amount uint64, service Service) (*TX, error) {
	var err error
	tx := TX{}
	tx.Locktime = 0
	var remain uint64
	tx.Txin, remain, err = setupP2PKHTXin(keys, amount+DefaultFee, service)
	if err != nil {
		return nil, err
	}

	adr, _ := keys[0].Pub.GetAddress()
	tx.Txout, err = setupP2PKHTXout([]*Amounts{&Amounts{adr, remain}})
	if err != nil {
		return nil, err
	}

	txout := TXout{}
	txout.Value = amount
	txout.ScriptPubkey = rs.createSriptPubkey()
	tx.Txout = append(tx.Txout, &txout)

	return &tx, nil
}

//Pay pays to a fund.
func (rs *RedeemScript) Pay(keys []*Key, amount uint64, service Service) ([]byte, error) {
	tx, err := rs.getMultisigTX(keys, amount, service)
	rawtx, err := tx.MakeTX()
	if err != nil {
		return nil, err
	}
	txHash, err := service.SendTX(rawtx)
	if err != nil {
		return nil, err
	}
	for _, txin := range tx.Txin {
		SetUTXOSpent(txin.Hash)
	}
	logging.Println("tx hash", hex.EncodeToString(txHash))
	return txHash, nil
}

//CreateRawTransactionHashed returns a hash of raw transaction for signing.
func (rs *RedeemScript) CreateRawTransactionHashed(addresses []*Amounts, service Service) ([]byte, *TX, error) {
	tx := TX{}
	tx.Locktime = 0

	var totalAmount uint64
	for _, amount := range addresses {
		totalAmount += amount.Amount
	}

	var utxo *UTXO
	txs, err := service.GetUTXO(rs.GetAddress(), nil)
	for _, tx := range txs {
		if tx.Amount >= totalAmount+DefaultFee {
			utxo = tx
			logging.Println("using tx:", hex.EncodeToString(utxo.Hash))
			break
		}
	}
	if utxo == nil {
		return nil, nil, errors.New("no utxo contains sufficient coin")
	}
	txin := TXin{}
	txin.Hash = utxo.Hash
	txin.Index = utxo.Index
	txin.Sequence = uint32(0xffffffff)
	txin.PrevScriptPubkey = rs.Script
	tx.Txin = []*TXin{&txin}

	tx.Txout, err = setupP2PKHTXout(addresses)
	if err != nil {
		return nil, nil, err
	}

	return tx.getRawTransactionHash(0), &tx, nil
}

//Spend spend the fund.
func (rs *RedeemScript) Spend(tx *TX, signs [][]byte, service Service) ([]byte, error) {
	tx.Txin[0].CreateScriptSig = func(rawTransaction []byte) ([]byte, error) {
		return rs.createScriptSig(rawTransaction, signs)
	}
	rawtx, err := tx.MakeTX()
	if err != nil {
		return nil, err
	}
	txHash, err := service.SendTX(rawtx)
	if err != nil {
		return nil, err
	}
	logging.Println("tx hash", hex.EncodeToString(txHash))
	return txHash, nil
}
