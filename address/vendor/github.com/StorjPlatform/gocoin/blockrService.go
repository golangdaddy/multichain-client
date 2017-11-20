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
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type unspent struct {
	Status string
	Data   struct {
		Address string
		Unspent []struct {
			Tx            string
			Amount        string
			N             int
			Confirmations int
			Script        string
		}
	}
	Code    int
	Message string
}

type sendtx struct {
	Status  string
	Data    string
	Code    int
	Message string
}

//BlockrService is a service using Blockr.io.
type BlockrService struct {
	isTestnet bool
}

//NewBlockrServiceForTest creates BlockrService struct for test.
func NewBlockrServiceForTest() (Service, error) {
	b := &BlockrService{isTestnet: true}
	return b, nil
}

//NewBlockrService creates BlockrService struct for not test.
func NewBlockrService() (Service, error) {
	return &BlockrService{isTestnet: false}, nil
}

//GetServiceName return service name.
func (b *BlockrService) GetServiceName() string {
	return "BlockrService"
}

//SendTX send a transaction using Blockr.io.
func (b *BlockrService) SendTX(data []byte) ([]byte, error) {
	var btc string

	if b.isTestnet {
		btc = "tbtc"
	} else {
		btc = "btc"
	}

	resp, err := http.PostForm("http://"+btc+".blockr.io/api/v1/tx/push",
		url.Values{"hex": {hex.EncodeToString(data)}})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	logging.Println(string(body))
	var u sendtx
	err = json.Unmarshal(body, &u)
	if err != nil {
		return nil, err
	}
	if u.Status != "success" {
		return nil, errors.New("blockr returns " + u.Message)

	}
	return hex.DecodeString(u.Data)
}

//GetUTXO gets unspent transaction outputs by using Blockr.io.
func (b *BlockrService) GetUTXO(addr string, key *Key) (UTXOs, error) {
	if cacheUTXO[addr] != nil {
		return cacheUTXO[addr], nil
	}
	var btc string

	if b.isTestnet {
		btc = "tbtc"
	} else {
		btc = "btc"
	}

	resp, err := http.Get("http://" + btc + ".blockr.io/api/v1/address/unspent/" + addr)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var u unspent
	err = json.Unmarshal(body, &u)
	if err != nil {
		return nil, err
	}
	if u.Status != "success" {
		return nil, errors.New("blockr returns " + u.Message)
	}

	utxos := make(UTXOs, 0, len(u.Data.Unspent))
	for _, tx := range u.Data.Unspent {
		utxo := UTXO{}
		utxo.Addr = addr
		amount, err := strconv.ParseFloat(tx.Amount, 64)
		if err != nil {
			return nil, err
		}
		utxo.Amount = uint64(amount * BTC)
		utxo.Hash, err = hex.DecodeString(tx.Tx)
		if err != nil {
			return nil, err
		}
		utxo.Index = uint32(tx.N)
		utxo.Script, err = hex.DecodeString(tx.Script)
		if err != nil {
			return nil, err
		}
		utxo.Age = uint64(tx.Confirmations)
		utxo.Key = key
		utxos = append(utxos, &utxo)
	}
	if key != nil {
		cacheUTXO[addr] = utxos
	}
	return utxos, nil
}
