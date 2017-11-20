[![Build Status](https://travis-ci.org/CoinStorage/gocoin.svg?branch=master)](https://travis-ci.org/StorjPlatform/gocoin)
[![GoDoc](https://godoc.org/github.com/StorjPlatform/gocoin?status.svg)](https://godoc.org/github.com/StorjPlatform/gocoin)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/StorjPlatform/gocoin/master/LICENSE)
[![Coverage Status](https://coveralls.io/repos/StorjPlatform/gocoin/badge.svg?branch=master)](https://coveralls.io/r/StorjPlatform/gocoin?branch=master)


# GOcoin 

## Overview

This is a library to make bitcoin address and transactions which was initially forked from [hellobitcoin](https://github.com/prettymuchbryce/hellobitcoin),
and added some useful features.

GOcoin uses btcec library in [btcd](https://github.com/btcsuite/btcd) instead of https://github.com/toxeus/go-secp256k1
to make it a pure GO program.


## Features 

1. Normaly Payment(P2PKH) supporting multi TxIns and multi TxOuts.
2. Gethering unspent transaction outputs(UTXO) and send transactions by using [Blockr.io](http://blockr.io) WEB API.
3. M of N multisig whose codes were partially ported from https://github.com/soroushjp/go-bitcoin-multisig.
4. Micropayment Channel


## Requirements

This requires

* git
* go 1.3+


## Installation

    $ mkdir tmp
    $ cd tmp
    $ mkdir src
    $ mkdir bin
    $ mkdir pkg
    $ exoprt GOPATH=`pwd`
    $ go get github.com/StorjPlatform/gocoin


## Example
(This example omits error handlings for simplicity.)

## Key Handling

```go

import gocoin

func main(){
	//make a public and private key pair.
	key, _ := gocoin.GenerateKey(true)
	adr, _ := key.Pub.GetAddress()
	fmt.Println("address=", adr)
	wif := key.Priv.GetWIFAddress()
	fmt.Println("wif=", wif)
	
	//get key from wif
	wif := "928Qr9J5oAC6AYieWJ3fG3dZDjuC7BFVUqgu4GsvRVpoXiTaJJf"
	txKey, _ := gocoin.GetKeyFromWIF(wif)
}
```

## Normal Payment

```go
import gocoin

func main(){
	key, _ := gocoin.GenerateKey(true)

	//get unspent transactions
	service := gocoin.NewBlockrService(true)
	txs, _ := service.GetUTXO(adr,nil)
	
	//Normal Payment
	gocoin.Pay([]*Key{txKey}, []*gocoin.Amounts{&{gocoin.Amounts{"n2eMqTT929pb1RDNuqEnxdaLau1rxy3efi", 0.01*gocoin.BTC}}, service)
}
```

## M of N Multisig

```go
import gocoin

func main(){
	key, _ := gocoin.GenerateKey(true)
	service := gocoin.NewBlockrService(true)

	//2 of 3 multisig
	key1, _ := gocoin.GenerateKey(true)
	key2, _ := gocoin.GenerateKey(true)
	key3, _ := gocoin.GenerateKey(true)
	rs, _:= gocoin.NewRedeemScript(2, []*PublicKey{key1.Pub, key2.Pub, key3.Pub})
	//make a fund
	rs.Pay([]*Key{txKey}, 0.05*gocoin.BTC, service)

    //get a raw transaction for signing.
	rawtx, tx, _:= rs.CreateRawTransactionHashed([]*gocoin.Amounts{&{gocoin.Amounts{"n3Bp1hbgtmwDtjQTpa6BnPPCA8fTymsiZy", 0.05*gocoin.BTC}}, service)

	//spend the fund
	sign1, _:= key2.Priv.Sign(rawtx)
	sign2, _:= key3.Priv.Sign(rawtx)
	rs.Spend(tx, [][]byte{nil, sign1, sign2}, service)
}
```


## Micropayment Channel

```go
import gocoin

func main(){
	service := gocoin.NewBlockrService(true)

	key1, _ := gocoin.GenerateKey(true) //payer
	key2, _ := gocoin.GenerateKey(true) //payee

	payer, _:= gocoin.NewMicropayer(key1, key2.Pub, service)
	payee, _:= gocoin.NewMicropayee(key2, key1.Pub, service)

	txHash, _:= payer.CreateBond([]*Key{key1}, 0.05*BTC)

	locktime := time.Now().Add(time.Hour)
	sign, _:= payee.SignToRefund(txHash, 0.05*gocoin.BTC-gocoin.Fee, uint32(locktime.Unix()))
	payer.SendBond(uint32(locktime.Unix()), sign) //return an error if payee's sig is invalid

	signIP, _:= payer.SignToIncrementedPayment(0.001 * gocoin.BTC)
	payee.IncrementPayment(0.001*gocoin.BTC, signIP) //return an error if payer's sig is invalid
	//more payments

	payee.SendLastPayment()
	//or
	//	payer.SendRefund() after locktime

}
```

Note:

payer.SendRefund() must be called after locktime.

http://chimera.labs.oreilly.com/books/1234000001802/ch05.html#tx_propagation

>Transactions with locktime specifying a future block or time must be held by the originating system
>and transmitted to the bitcoin network only after they become valid.


# Contribution
Improvements to the codebase and pull requests are encouraged.


