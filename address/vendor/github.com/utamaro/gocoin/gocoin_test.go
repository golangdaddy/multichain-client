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
 */

package gocoin

import (
	"bytes"
	"encoding/hex"
	"sort"
	"testing"
	"time"
)

func TestKeys2(t *testing.T) {
	key, err := GenerateKey(true)
	if err != nil {
		t.Errorf(err.Error())
	}
	adr, _ := key.Pub.GetAddress()
	logging.Println("address=", adr)
	wif := key.Priv.GetWIFAddress()
	logging.Println("wif=", wif)

	key2, err := GetKeyFromWIF(wif)
	if err != nil {
		t.Errorf(err.Error())
	}
	adr2, _ := key2.Pub.GetAddress()
	logging.Println("address2=", adr2)

	if adr != adr2 {
		t.Errorf("key unmatched")
	}
}

func TestKeys(t *testing.T) {
	key, err := GenerateKey(false)
	if err != nil {
		t.Errorf(err.Error())
	}
	adr, _ := key.Pub.GetAddress()
	logging.Println("address=", adr)
	wif := key.Priv.GetWIFAddress()
	logging.Println("wif=", wif)

	key2, err := GetKeyFromWIF(wif)
	if err != nil {
		t.Errorf(err.Error())
	}
	adr2, _ := key2.Pub.GetAddress()
	logging.Println("address2=", adr2)

	if adr != adr2 {
		t.Errorf("key unmatched")
	}

}

func TestTX(t *testing.T) {
	wif := "928Qr9J5oAC6AYieWJ3fG3dZDjuC7BFVUqgu4GsvRVpoXiTaJJf"
	txKey, err := GetKeyFromWIF(wif)
	if err != nil {
		t.Errorf(err.Error())
	}
	adr, _ := txKey.Pub.GetAddress()
	logging.Println("address for tx=", adr)
	if adr != "n3Bp1hbgtmwDtjQTpa6BnPPCA8fTymsiZy" {
		t.Errorf("invalid address")
	}

	txin := TXin{}
	txin.Hash, err = hex.DecodeString("1a103718e2e0462c50cb057a0f39d7c6cbf960276452d07dc4a50ddca725949c")
	if err != nil {
		t.Errorf(err.Error())
	}
	txin.Index = 1
	txin.Sequence = uint32(0xffffffff)
	txin.PrevScriptPubkey, err = createP2PKHScriptPubkey(adr)
	if err != nil {
		t.Errorf(err.Error())
	}
	txin.CreateScriptSig = func(rawTransaction []byte) ([]byte, error) {
		return createP2PKHScriptSig(rawTransaction, txKey)
	}
	txout := TXout{}
	txout.Value = 68000000
	txout.ScriptPubkey, err = createP2PKHScriptPubkey("n2eMqTT929pb1RDNuqEnxdaLau1rxy3efi")
	if err != nil {
		t.Errorf(err.Error())
	}
	tx := TX{}
	tx.Txin = []*TXin{&txin}
	tx.Txout = []*TXout{&txout}
	tx.Locktime = 0

	rawtx, err := tx.MakeTX()
	if err != nil {
		t.Errorf(err.Error())
	}
	ok := "01000000019c9425a7dc0da5c47dd052642760f9cbc6d7390f7a05cb502c46e0e21837101a010000008a473044022030ebb89d54e76b9e14b8eb21aa30055eb54289dcd3aad9b415ebcc153b211eee0220720fa77cfc2c25da52899f3bf9a947869bc89d26066c02a1c428e9530a3f49b10141049f160b18fa4acedccdc063961d63b3a23385b1e67159d07521cb46d4e7209ecd443e473796e7ace130164c660fbcfb7dcac8437cc55f3ceafb546054c8d8cbdfffffffff0100990d04000000001976a914e7c1345fc8f87c68170b3aa798a956c2fe6a9eff88ac00000000"
	if hex.EncodeToString(rawtx) != ok {
		t.Errorf("invalid tx")
	}
}

func TestSend(t *testing.T) {
	wif := "928Qr9J5oAC6AYieWJ3fG3dZDjuC7BFVUqgu4GsvRVpoXiTaJJf"
	txKey, err := GetKeyFromWIF(wif)
	if err != nil {
		t.Errorf(err.Error())
	}
	adr, _ := txKey.Pub.GetAddress()
	logging.Println("address for tx=", adr)
	if adr != "n3Bp1hbgtmwDtjQTpa6BnPPCA8fTymsiZy" {
		t.Errorf("invalid address")
	}
	service, err := SelectService(true)
	if err != nil {
		t.Errorf(err.Error())
	}
	txs, err := service.GetUTXO(adr, nil)
	if err != nil {
		t.Errorf(err.Error())
	}
	logging.Println("UTXO:")
	for _, tx := range txs {
		logging.Println("hash", hex.EncodeToString(tx.Hash))
		logging.Println("amount", tx.Amount)
		logging.Println("index", tx.Index)
		logging.Println("script", hex.EncodeToString(tx.Script))
	}
	amounts := []*Amounts{&Amounts{"n3Bp1hbgtmwDtjQTpa6BnPPCA8fTymsiZy", 0.05 * BTC}, &Amounts{"n2eMqTT929pb1RDNuqEnxdaLau1rxy3efi", 0.01 * BTC}}
	_, err = Pay([]*Key{txKey}, amounts, service)
	if err != nil {
		t.Errorf(err.Error())
	}

}

func TestRedeemScript(t *testing.T) {
	pks := []string{
		"04a882d414e478039cd5b52a92ffb13dd5e6bd4515497439dffd691a0f12af9575fa349b5694ed3155b136f09e63975a1700c9f4d4df849323dac06cf3bd6458cd",
		"046ce31db9bdd543e72fe3039a1f1c047dab87037c36a669ff90e28da1848f640de68c2fe913d363a51154a0c62d7adea1b822d05035077418267b1a1379790187",
		"0411ffd36c70776538d079fbae117dc38effafb33304af83ce4894589747aee1ef992f63280567f52f5ba870678b4ab4ff6c8ea600bd217870a8b4f1f09f3a8e83",
	}
	publicKeys := make([]*PublicKey, 3)
	for i, pk := range pks {
		pkb, err := hex.DecodeString(pk)
		if err != nil {
			t.Errorf(err.Error())
		}
		publicKeys[i], err = GetPublicKey(pkb, false)
		if err != nil {
			t.Errorf(err.Error())
		}
		logging.Println(publicKeys[i].GetAddress())
	}
	rs, err := NewRedeemScript(2, publicKeys)
	if err != nil {
		t.Errorf(err.Error())
	}
	scriptStr := "524104a882d414e478039cd5b52a92ffb13dd5e6bd4515497439dffd691a0f12af9575fa349b5694ed3155b136f09e63975a1700c9f4d4df849323dac06cf3bd6458cd41046ce31db9bdd543e72fe3039a1f1c047dab87037c36a669ff90e28da1848f640de68c2fe913d363a51154a0c62d7adea1b822d05035077418267b1a1379790187410411ffd36c70776538d079fbae117dc38effafb33304af83ce4894589747aee1ef992f63280567f52f5ba870678b4ab4ff6c8ea600bd217870a8b4f1f09f3a8e8353ae"
	address := "347N1Thc213QqfYCz3PZkjoJpNv5b14kBd"
	script, err := hex.DecodeString(scriptStr)
	if err != nil {
		t.Errorf(err.Error())
	}
	if address != rs.GetAddress() {
		t.Errorf("address is bad")
	}
	if bytes.Compare(rs.Script, script) != 0 {
		t.Errorf("script is bad")
	}
	logging.Println("address", rs.GetAddress())
	logging.Println("script", hex.EncodeToString(rs.Script))

	inputTX := "3ad337270ac0ba14fbce812291b7d95338c878709ea8123a4d88c3c29efbc6ac"
	okTX := "0100000001acc6fb9ec2c3884d3a12a89e7078c83853d9b7912281cefb14bac00a2737d33a000000008b483045022100fb244ac83b257f4233920077819dfa5203a11cd330c58a37c984699bc8048e9102200caca5b3772022a5cb5ce8e31f644da4e27e2c4f121cfd9b5291e3bccf7017d701410431393af9984375830971ab5d3094c6a7d02db3568b2b06212a7090094549701bbb9e84d9477451acc42638963635899ce91bacb451a1bb6da73ddfbcf596bddfffffffff01400001000000000017a9141a8b0026343166625c7475f01e48b5ede8c0252e8700000000"
	oktx, err := hex.DecodeString(okTX)
	if err != nil {
		t.Errorf(err.Error())
	}

	txHash, err := hex.DecodeString(inputTX)
	if err != nil {
		t.Errorf(err.Error())
	}

	//Create a fund.
	txkey, err := GetKeyFromWIF("5JJyqG4bb15zqi7fTA4b227aUxQhBo1Ux6qX69ngeXYLr7fk2hs")
	if err != nil {
		t.Errorf(err.Error())
	}

	txin := TXin{}
	txin.Hash = txHash
	txin.Index = 0
	txin.Sequence = uint32(0xffffffff)
	adr, _ := txkey.Pub.GetAddress()
	logging.Println("adr", adr)
	txin.PrevScriptPubkey, err = createP2PKHScriptPubkey(adr)
	if err != nil {
		t.Errorf(err.Error())
	}
	txin.CreateScriptSig = func(rawTransaction []byte) ([]byte, error) {
		return createP2PKHScriptSig(rawTransaction, txkey)
	}

	txout := TXout{}
	txout.Value = 65600
	txout.ScriptPubkey = rs.createSriptPubkey()
	tx := TX{}
	tx.Txin = []*TXin{&txin}
	tx.Txout = []*TXout{&txout}
	tx.Locktime = 0

	rawtx, err := tx.MakeTX()
	if err != nil {
		t.Errorf(err.Error())
	}
	logging.Println(hex.EncodeToString(rawtx))
	if bytes.Compare(rawtx, oktx) != 0 {
		t.Errorf("transaction is bad")
	}

	//spend the fund
	inputTX = "02b082113e35d5386285094c2829e7e2963fa0b5369fb7f4b79c4c90877dcd3d"
	pks = []string{
		"5JruagvxNLXTnkksyLMfgFgf3CagJ3Ekxu5oGxpTm5mPfTAPez3",
		"5JjHVMwJdjPEPQhq34WMUhzLcEd4SD7HgZktEh8WHstWcCLRceV",
	}
	okTX = "01000000013dcd7d87904c9cb7f4b79f36b5a03f96e2e729284c09856238d5353e1182b00200000000fd5e0100483045022100af4831eb0cee1b9642ff8691acb53c2fd8e28bbd6c0389af69c132b342405c8502200627e67cbb0bee502421be4d2b8e370d24eb17ff7c3d77d604a5c07ee5934a74014830450221009bfee48a5ae99a8fc0f78c631b0b2beb6d96337fc9d1e95c333a08bdcecb3647022040a30b3528b136821077d0bb45c324d7eedcd0ac61af7019ba1286e9277b7c54014cc9524104a882d414e478039cd5b52a92ffb13dd5e6bd4515497439dffd691a0f12af9575fa349b5694ed3155b136f09e63975a1700c9f4d4df849323dac06cf3bd6458cd41046ce31db9bdd543e72fe3039a1f1c047dab87037c36a669ff90e28da1848f640de68c2fe913d363a51154a0c62d7adea1b822d05035077418267b1a1379790187410411ffd36c70776538d079fbae117dc38effafb33304af83ce4894589747aee1ef992f63280567f52f5ba870678b4ab4ff6c8ea600bd217870a8b4f1f09f3a8e8353aeffffffff0130d90000000000001976a914569076ba39fc4ff6a2291d9ea9196d8c08f9c7ab88ac00000000"
	oktx, err = hex.DecodeString(okTX)
	if err != nil {
		t.Errorf(err.Error())
	}

	txHash, err = hex.DecodeString(inputTX)
	if err != nil {
		t.Errorf(err.Error())
	}
	keys := make([]*Key, 2)
	for i, pk := range pks {
		keys[i], err = GetKeyFromWIF(pk)
		if err != nil {
			t.Errorf(err.Error())
		}
		logging.Println(keys[i].Pub.GetAddress())
	}

	txin2 := TXin{}
	txin2.Hash = txHash
	txin2.Index = 0
	txin2.Sequence = uint32(0xffffffff)
	txin2.PrevScriptPubkey = rs.Script

	txout2 := TXout{}
	txout2.Value = 55600
	txout2.ScriptPubkey, err = createP2PKHScriptPubkey("18tiB1yNTzJMCg6bQS1Eh29dvJngq8QTfx")
	if err != nil {
		t.Errorf(err.Error())
	}
	tx2 := TX{}
	tx2.Txin = []*TXin{&txin2}
	tx2.Txout = []*TXout{&txout2}
	tx2.Locktime = 0

	rawtx = tx2.getRawTransactionHash(0)

	sign1, err := keys[0].Priv.Sign(rawtx)
	if err != nil {
		t.Errorf(err.Error())
	}
	sign2, err := keys[1].Priv.Sign(rawtx)
	if err != nil {
		t.Errorf(err.Error())
	}

	txin2.CreateScriptSig = func(rawTransaction []byte) ([]byte, error) {
		return rs.createScriptSig(rawTransaction, [][]byte{sign1, nil, sign2})
	}

	rawtx, err = tx2.MakeTX()
	if err != nil {
		t.Errorf(err.Error())
	}
	if bytes.Compare(rawtx, oktx) != 0 {
		t.Errorf("transaction is bad")
	}
}

func TestMultisig(t *testing.T) {
	service, err := SelectService(true)
	if err != nil {
		t.Errorf(err.Error())
	}
	wif := "928Qr9J5oAC6AYieWJ3fG3dZDjuC7BFVUqgu4GsvRVpoXiTaJJf"
	//n3Bp1hbgtmwDtjQTpa6BnPPCA8fTymsiZy
	txKey, err := GetKeyFromWIF(wif)
	if err != nil {
		t.Errorf(err.Error())
	}
	adr, _ := txKey.Pub.GetAddress()
	logging.Println("address for tx=", adr)

	wif2 := "92DUfNPumHzpCkKjmeqiSEDB1PU67eWbyUgYHhK9ziM7NEbqjnK"
	//ms5repuZHtBrKRE93FdWqz8JEo6d8ikM3k
	txKey2, err := GetKeyFromWIF(wif2)
	if err != nil {
		t.Errorf(err.Error())
	}
	wif3 := "92QdHTuahkv52BteRvCidZmWU9mFaeSs6Key6Lfy6KxzCEKLcsa"
	//mhZobPCtc3NhwG2yfmGnKaohkpY5ZBNkC6
	txKey3, err := GetKeyFromWIF(wif3)
	if err != nil {
		t.Errorf(err.Error())
	}

	s, err := NewBlockrServiceForTest()
	if err != nil {
		t.Errorf(err.Error())
	}
	txs, err := s.GetUTXO(adr, nil)
	if err != nil {
		t.Errorf(err.Error())
	}
	logging.Println("UTXO of ", adr)
	for _, tx := range txs {
		logging.Println("hash", hex.EncodeToString(tx.Hash))
		logging.Println("amount", tx.Amount)
		logging.Println("index", tx.Index)
		logging.Println("script", hex.EncodeToString(tx.Script))
	}
	//Create a fund.
	rs, err := NewRedeemScript(2, []*PublicKey{txKey.Pub, txKey2.Pub, txKey3.Pub})
	if err != nil {
		t.Errorf(err.Error())
	}
	_, err = rs.Pay([]*Key{txKey}, 5000000, service)
	if err != nil {
		t.Errorf(err.Error())
	}

	txs, err = s.GetUTXO(rs.GetAddress(), nil)
	sort.Sort(UTXOs(txs))
	if err != nil {
		t.Errorf(err.Error())
	}
	logging.Println("UTXO of", rs.GetAddress())
	for _, tx := range txs {
		logging.Println("hash", hex.EncodeToString(tx.Hash))
		logging.Println("amount", tx.Amount)
		logging.Println("index", tx.Index)
		logging.Println("script", hex.EncodeToString(tx.Script))
	}

	//spend the fund
	rawtx, tx, err := rs.CreateRawTransactionHashed([]*Amounts{&Amounts{"n3Bp1hbgtmwDtjQTpa6BnPPCA8fTymsiZy", txs[0].Amount - DefaultFee}}, service)
	if err != nil {
		t.Errorf(err.Error())
	}
	sign1, err := txKey2.Priv.Sign(rawtx)
	if err != nil {
		t.Errorf(err.Error())
	}
	sign2, err := txKey3.Priv.Sign(rawtx)
	if err != nil {
		t.Errorf(err.Error())
	}
	_, err = rs.Spend(tx, [][]byte{nil, sign1, sign2}, service)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestMicro(t *testing.T) {
	service, err := SelectService(true)
	if err != nil {
		t.Errorf(err.Error())
	}
	wif := "928Qr9J5oAC6AYieWJ3fG3dZDjuC7BFVUqgu4GsvRVpoXiTaJJf"
	//n3Bp1hbgtmwDtjQTpa6BnPPCA8fTymsiZy
	txKey, err := GetKeyFromWIF(wif)
	if err != nil {
		t.Errorf(err.Error())
	}
	adr, _ := txKey.Pub.GetAddress()
	logging.Println("address for tx=", adr)

	wif2 := "92DUfNPumHzpCkKjmeqiSEDB1PU67eWbyUgYHhK9ziM7NEbqjnK"
	//ms5repuZHtBrKRE93FdWqz8JEo6d8ikM3k
	txKey2, err := GetKeyFromWIF(wif2)
	if err != nil {
		t.Errorf(err.Error())
	}

	txs, err := service.GetUTXO(adr, nil)
	if err != nil {
		t.Errorf(err.Error())
	}
	logging.Println("UTXO of ", adr)
	for _, tx := range txs {
		logging.Println("hash", hex.EncodeToString(tx.Hash))
		logging.Println("amount", tx.Amount)
		logging.Println("index", tx.Index)
		logging.Println("script", hex.EncodeToString(tx.Script))
	}

	payer, err := NewMicropayer(txKey, txKey2.Pub, service)
	if err != nil {
		t.Errorf(err.Error())
	}

	payee, err := NewMicropayee(txKey2, txKey.Pub, service)
	if err != nil {
		t.Errorf(err.Error())
	}

	txHash, err := payer.CreateBond([]*Key{txKey}, 0.05*BTC)
	if err != nil {
		t.Errorf(err.Error())
	}
	locktime := time.Now().Add(time.Hour)
	sign, err := payee.SignToRefund(txHash, 0.05*BTC-DefaultFee, &locktime)
	if err != nil {
		t.Errorf(err.Error())
	}
	_, err = payer.SendBond(&locktime, sign)
	if err != nil {
		t.Fatalf(err.Error())
	}
	signIP, err := payer.SignToIncrementedPayment(0.001 * BTC)
	if err != nil {
		t.Fatalf(err.Error())
	}
	logging.Println(hex.EncodeToString(signIP))
	err = payee.IncrementPayment(0.001*BTC, signIP)
	if err != nil {
		t.Fatalf(err.Error())
	}
	signIP, err = payer.SignToIncrementedPayment(0.001 * BTC)
	logging.Println(hex.EncodeToString(signIP))
	if err != nil {
		t.Fatalf(err.Error())
	}
	err = payee.IncrementPayment(0.001*BTC, signIP)
	if err != nil {
		t.Fatalf(err.Error())
	}
	_, err = payee.SendLastPayment()
	if err != nil {
		t.Fatalf(err.Error())
	}
	//	_, err = payer.SendRefund()
	//	if err != nil {
	//		t.Errorf(err.Error())
	//	}
}
