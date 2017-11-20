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
 */

package gocoin

import (
	"bytes"
	"math/rand"
)

//UTXO represents unspent transaction outputs.
type UTXO struct {
	Addr   string
	Hash   []byte
	Amount uint64
	Index  uint32
	Script []byte
	Age    uint64
	Key    *Key
}

//UTXOs is for sorting UTXO
type UTXOs []*UTXO

var cacheUTXO = make(map[string]UTXOs)

//Service is for getting UTXO or sending transactions , basically by using WEB API.
type Service interface {
	GetServiceName() string
	GetUTXO(string, *Key) (UTXOs, error)
	SendTX([]byte) ([]byte, error)
}

//TestServices is an array containing generator of Services for testnet
var TestServices = []func() (Service, error){
	NewBlockrServiceForTest,
}

//TestServices is an array containing generator of Services
var Services = []func() (Service, error){
	NewBlockrService,
}

//to sort UTXO

//Len returns length of UTXO
func (us UTXOs) Len() int {
	return len(us)
}

//Swap swaps UTXO
func (us UTXOs) Swap(i, j int) {
	us[i], us[j] = us[j], us[i]
}

//Less returns true is age is smaller.
func (us UTXOs) Less(i, j int) bool {
	return us[i].Amount < us[j].Amount
}

//SelectService returns a service randomly.
func SelectService(isTestnet bool) (Service, error) {
	n := rand.Int() % len(Services)
	if isTestnet {
		return TestServices[n]()
	}
	return Services[n]()
}

//SetTXSpent sets  tx hash is already spent.
func SetUTXOSpent(hash []byte) {
	for k, v := range cacheUTXO {
		for i, utxo := range v {
			if bytes.Compare(hash, utxo.Hash) == 0 {
				v = append(v[0:i], v[i+1:]...)
				cacheUTXO[k] = v
				return
			}
		}
	}
}
