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
 */

package gocoin

import (
	"encoding/hex"
	"errors"
	"time"
)

type micropayment struct {
	key         *Key
	rs          *RedeemScript
	service     Service
	bondHash    []byte
	TotalAmount uint64
	Paied       uint64
	Locktime    *time.Time
}

//Micropayer represents payer of Micropayment.
type Micropayer struct {
	micropayment
	bond   *TX
	refund *TX
}

//Micropayee represents payee of Micropayment.
type Micropayee struct {
	micropayment
	lastTX *TX
}

//NewMicropayer creates new MicroPayer struct.
func NewMicropayer(key *Key, publicKey *PublicKey, service Service) (*Micropayer, error) {
	var err error
	m := Micropayer{}
	m.key = key
	m.service = service
	m.rs, err = NewRedeemScript(2, []*PublicKey{key.Pub, publicKey})
	if err != nil {
		return nil, err
	}
	return &m, nil
}

//NewMicropayee creates new MicroPayee struct.
func NewMicropayee(key *Key, publicKey *PublicKey, service Service) (*Micropayee, error) {
	var err error
	m := Micropayee{}
	m.key = key
	m.service = service
	m.rs, err = NewRedeemScript(2, []*PublicKey{publicKey, key.Pub})
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (m *micropayment) createRefund(amount []*Amounts, lockTime *time.Time, payerSign []byte, payeeSign []byte) ([]byte, *TX, error) {
	var err error
	refund := &TX{}
	txin := TXin{}
	txin.Hash = m.bondHash
	txin.Index = 1
	if lockTime == nil {
		refund.Locktime = 0
		txin.Sequence = uint32(0xffffffff)
	} else {
		refund.Locktime = uint32(lockTime.Unix())
		txin.Sequence = 0
	}
	txin.PrevScriptPubkey = m.rs.Script
	refund.Txin = []*TXin{&txin}

	refund.Txout, err = setupP2PKHTXout(amount)
	if err != nil {
		return nil, nil, err
	}
	rawtx := refund.getRawTransactionHash(0)
	mySign, err := m.key.Priv.Sign(rawtx)
	if err != nil {
		return nil, nil, err
	}
	if payerSign == nil && payeeSign == nil {
		return mySign, nil, nil
	}
	var signPayee, signPayer []byte
	if payerSign == nil {
		signPayer = mySign
		signPayee = payeeSign
	}
	if payeeSign == nil {
		signPayer = payerSign
		signPayee = mySign
	}
	refund.Txin[0].CreateScriptSig = func(rawTransaction []byte) ([]byte, error) {
		return m.rs.createScriptSig(rawTransaction, [][]byte{signPayer, signPayee})
	}
	return mySign, refund, nil
}

//CreateBond creates bond and return its hash.
func (m *Micropayer) CreateBond(keys []*Key, amount uint64) ([]byte, error) {
	var err error

	m.TotalAmount = amount - DefaultFee
	m.bond, err = m.rs.getMultisigTX(keys, amount, m.service)
	if err != nil {
		return nil, err
	}
	m.bondHash, err = m.bond.getTransactionHash()
	if err != nil {
		return nil, err
	}
	logging.Println("hash of bond", hex.EncodeToString(m.bondHash))
	return m.bondHash, nil
}

//SignToRefund create signature of refund tx.
func (m *Micropayee) SignToRefund(txHash []byte, amount uint64, lockTime *time.Time) ([]byte, error) {
	m.Locktime = lockTime
	m.bondHash = txHash
	m.TotalAmount = amount
	addr, _ := m.rs.PublicKeys[0].GetAddress()
	amounts := []*Amounts{&Amounts{addr, amount}}
	sign, _, err := m.createRefund(amounts, lockTime, nil, nil)
	return sign, err
}

//SendBond send bond and refunds and returns bond tx hash.
func (m *Micropayer) SendBond(lockTime *time.Time, sign []byte) ([]byte, error) {
	var err error
	if m.bond == nil {
		return nil, errors.New("create bond first")
	}
	m.Locktime = lockTime
	addr, _ := m.key.Pub.GetAddress()
	amounts := []*Amounts{&Amounts{addr, m.TotalAmount}}
	_, m.refund, err = m.createRefund(amounts, lockTime, nil, sign)
	txRefund, err := m.refund.MakeTX()
	if err != nil {
		return nil, err
	}
	hash, err := m.refund.getTransactionHash()
	if err != nil {
		return nil, err
	}
	logging.Println("refund tx", hex.EncodeToString(txRefund))
	logging.Println("hash of refund tx", hex.EncodeToString(hash))

	rawtx, err := m.bond.MakeTX()
	if err != nil {
		return nil, err
	}
	txHashBond, err := m.service.SendTX(rawtx)
	if err != nil {
		return nil, err
	}
	logging.Println("tx hash of bond", hex.EncodeToString(txHashBond))

	return txHashBond, nil
}

//SignToIncrementedPayment creates signature for incremented payment tx.
func (m *Micropayer) SignToIncrementedPayment(increment uint64) ([]byte, error) {
	if m.refund == nil {
		return nil, errors.New("send bond first")
	}
	if time.Now().After(*m.Locktime) {
		return nil, errors.New("already passed locktime")
	}
	m.Paied += increment
	payer, _ := m.rs.PublicKeys[0].GetAddress()
	payee, _ := m.rs.PublicKeys[1].GetAddress()
	amounts := []*Amounts{&Amounts{payer, m.TotalAmount - m.Paied}, &Amounts{payee, m.Paied}}
	sign, _, err := m.createRefund(amounts, nil, nil, nil)
	logging.Println("amount paied", m.Paied)
	return sign, err
}

//IncrementPayment increment payment ,creates transaction and store its tx into struct.
func (m *Micropayee) IncrementPayment(increment uint64, sign []byte) error {
	var err error
	if m.bondHash == nil {
		return errors.New("receive refund tx first")
	}
	if time.Now().After(*m.Locktime) {
		return errors.New("already passed locktime")
	}

	payer, _ := m.rs.PublicKeys[0].GetAddress()
	payee, _ := m.rs.PublicKeys[1].GetAddress()
	amounts := []*Amounts{&Amounts{payer, m.TotalAmount - m.Paied - increment}, &Amounts{payee, m.Paied + increment}}
	_, lastTX, err := m.createRefund(amounts, nil, sign, nil)
	if err != nil {
		return err
	}
	_, err = lastTX.MakeTX()
	if err != nil {
		logging.Println(err, hex.EncodeToString(sign))
		return err
	}
	m.lastTX = lastTX
	m.Paied += increment
	logging.Println("amount paied", m.Paied)
	return nil
}

//SendLastPayment sends last incremented tx and returns tx hash.
func (m *Micropayee) SendLastPayment() ([]byte, error) {
	if m.lastTX == nil {
		return nil, errors.New("receive incremented refund tx first")
	}
	rawtx, err := m.lastTX.MakeTX()
	if err != nil {
		return nil, err
	}
	txHash, err := m.service.SendTX(rawtx)
	if err != nil {
		return nil, err
	}
	m.lastTX = nil
	m.bondHash = nil
	m.TotalAmount = 0
	m.Paied = 0
	m.Locktime = nil
	return txHash, nil
}

//SendRefund sends refunt tx and returns tx hash.
//bond tx will be locked after sending thye refund (?).
func (m *Micropayer) SendRefund() ([]byte, error) {
	if m.refund == nil {
		return nil, errors.New("receive refund signature first")
	}
	if time.Now().Before(*m.Locktime) {
		return nil, errors.New("not passed locktime yet")
	}

	rawtx, err := m.refund.MakeTX()
	if err != nil {
		return nil, err
	}
	txHash, err := m.service.SendTX(rawtx)
	if err != nil {
		return nil, err
	}
	m.bond = nil
	m.Paied = 0
	m.refund = nil
	m.bondHash = nil
	m.TotalAmount = 0
	m.Locktime = nil
	return txHash, nil
}
