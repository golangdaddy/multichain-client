package signature

import (
	"fmt"
	"bytes"
//	"strconv"
	"crypto/sha256"
	"encoding/hex"
	"encoding/binary"
	//
	"github.com/btcsuite/btcd/btcec"
	//
	"gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/multichain/opcodes"
)

type SigReq struct {
	transaction *Transaction
	publicKey []byte
	privateKey []byte
}

/*

def makeSignedTransaction(privateKey, outputTransactionHash, sourceIndex, scriptPubKey, outputs):
    myTxn_forSig = (makeRawTransaction(outputTransactionHash, sourceIndex, scriptPubKey, outputs)
         + "01000000") # hash code

    s256 = hashlib.sha256(hashlib.sha256(myTxn_forSig.decode('hex')).digest()).digest()
    sk = ecdsa.SigningKey.from_string(privateKey.decode('hex'), curve=ecdsa.SECP256k1)
    sig = sk.sign_digest(s256, sigencode=ecdsa.util.sigencode_der) + '\01' # 01 is hashtype
    pubKey = keyUtils.privateKeyToPublicKey(privateKey)
    scriptSig = utils.varstr(sig).encode('hex') + utils.varstr(pubKey.decode('hex')).encode('hex')
    signed_txn = makeRawTransaction(outputTransactionHash, sourceIndex, scriptSig, outputs)
    verifyTxnSignature(signed_txn)
    return signed_txn

*/

func (self *SigReq) Sign(index int) error {

	hashType := 0x01

	buf, err := toBuffer(self.transaction, 0)
	if err != nil {
		return err
	}

	hashForSignature := sha256.Sum256(
		bytes.Join(
			[][]byte{
				buf,
				uint32Buffer(uint32(hashType)),
			},
			nil,
		),
	)

	privkey, _ := btcec.PrivKeyFromBytes(btcec.S256(), self.privateKey)
	sig, err := privkey.Sign(hashForSignature[:])
	if err != nil {
		return err
	}

	scriptSignature := bytes.Join(
		[][]byte{
			sig.Serialize(),
			uint8Buffer(uint8(hashType)),
		},
		nil,
	)

	self.transaction.Vin[index].ScriptSig.Hex = hex.EncodeToString(
		bytes.Join(
			[][]byte{
				pushDataIntBuffer(len(scriptSignature)),
				scriptSignature,
				pushDataIntBuffer(len(self.publicKey)),
				self.publicKey,
			},
			nil,
		),
	)

	return nil
}

func toBuffer(tx *Transaction, index int) ([]byte, error) {

	chunks := [][]byte{}

	for _, txIn := range tx.Vin {

		txid, _ := hex.DecodeString(txIn.Txid)
		chunks = append(
			chunks,
			txid,
			uint8Buffer(uint8(txIn.Vout)),
		)

		script := txIn.ScriptSig.Hex

		if len(script) > 0 {

			chunks = append(
				chunks,
				varIntBuffer(uint64(len(script))),
			)

			chunks = append(
				chunks,
				[]byte(script),
			)

		} else {

			chunks = append(
				chunks,
				varIntBuffer(0),
			)

		}

		chunks = append(
			chunks,
			uint32Buffer(uint32(txIn.Sequence)),
		)

	}

	for _, txOut := range tx.Vout {

		script, _ := hex.DecodeString(txOut.ScriptPubKey.Hex)
		chunks = append(
			chunks,
			uint64Buffer(uint64(txOut.Value)),
			varIntBuffer(uint64(len(script))),
			script,
		)

	}

	return bytes.Join(chunks, nil), nil
}

func pushDataIntBuffer(num int) []byte {

	var chunks [][]byte

	if num < opcodes.OP_PUSHDATA1 {

		chunks = append(
			chunks,
			uint8Buffer(uint8(num)),
		)

	} else if num < 0xff {

		chunks = append(
			chunks,
			uint8Buffer(opcodes.OP_PUSHDATA1),
			uint8Buffer(uint8(num)),
		)

	} else if num < 0xffff {

		chunks = append(
			chunks,
			uint8Buffer(opcodes.OP_PUSHDATA2),
			uint16Buffer(uint16(num)),
		)

	} else {

		chunks = append(
			chunks,
			uint8Buffer(opcodes.OP_PUSHDATA4),
			uint32Buffer(uint32(num)),
		)

	}

	return bytes.Join(chunks, nil)
}

func varIntBuffer(num uint64) []byte {

	var chunks [][]byte

	if num < 253 {

		chunks = append(
			chunks,
			uint8Buffer(uint8(num)),
		)

	} else if num < 0x10000 {

		chunks = append(
			chunks,
			uint16Buffer(uint16(num)),
		)

	} else if num < 0x100000000 {

		chunks = append(
			chunks,
			uint32Buffer(uint32(num)),
		)

	} else {

		chunks = append(
			chunks,
			uint64Buffer(uint64(num)),
		)

	}

	return bytes.Join(chunks, nil)
}

/*
var uint8Buffer = function(number) {
	var buffer = new Buffer(1);
	buffer.writeUInt8(number, 0);

	return buffer;
};
*/

func uint8Buffer(num uint8) []byte {

	return []byte{byte(num)}
}

/*
var uint16Buffer = function(number) {
	var buffer = new Buffer(2);
	buffer.writeUInt16LE(number, 0);

	return buffer;
};
*/

func uint16Buffer(num uint16) []byte {

	b := make([]byte, 3)
	binary.LittleEndian.PutUint16(b[0:], 0xFD)
	binary.LittleEndian.PutUint16(b[1:], num)

	return b
}

/*
var uint32Buffer = function(number) {
	var buffer = new Buffer(4);
	buffer.writeUInt32LE(number, 0);
	return buffer;
};
*/

func uint32Buffer(num uint32) []byte {

	b := make([]byte, 5)
	binary.LittleEndian.PutUint16(b[0:], 0xFE)
	binary.LittleEndian.PutUint32(b[1:], num)

	return b
}

/*
var uint64Buffer = function(number) {
	var buffer = new Buffer(8);
	buffer.writeInt32LE(number & -1, 0)
	buffer.writeUInt32LE(Math.floor(number / 0x100000000), 4)

	return buffer;
};
*/

func uint64Buffer(num uint64) []byte {

	b := make([]byte, 9)
	binary.LittleEndian.PutUint16(b[0:], 0xFF)
	binary.LittleEndian.PutUint64(b[1:], num)

	return b
}
