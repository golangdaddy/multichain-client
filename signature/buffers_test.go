package signature

import (
    "fmt"
    "testing"
    "encoding/json"
    //
    "gitlab.com/golangdaddy/multichain-client/address"
)

const (
    CONST_TX_TEST = "01000000013c17411f8b705c4fe6b9cebf8501d3b6e59aaa2052f376962eea5283b6fde2270000000000ffffffff0200000000000000003776a91421ce0c375fca548b4c39847c5a0df06c8d1bd0df88ac1c73706b71e59aaa2052f376962eea5283b6fde227e8030000000000007500000000000000003776a914fca3548c1a26d433964f713bada36ee66e66836a88ac1c73706b71e59aaa2052f376962eea5283b6fde227b8820100000000007500000000"
    CONST_TX_DECODED = `{"txid":"bdb94203dce2fc8f804157c7ec0fd1a84247f87f197a15d5ec5edc33a5e6f996","version":1,"locktime":0,"vin":[{"txid":"27e2fdb68352ea2e9676f35220aa9ae5b6d30185bfceb9e64f5c708b1f41173c","vout":0,"scriptSig":{"asm":"","hex":""},"sequence":4294967295}],"vout":[{"value":0,"n":0,"scriptPubKey":{"asm":"OP_DUP OP_HASH160 21ce0c375fca548b4c39847c5a0df06c8d1bd0df OP_EQUALVERIFY OP_CHECKSIG 73706b71e59aaa2052f376962eea5283b6fde227e803000000000000 OP_DROP","hex":"76a91421ce0c375fca548b4c39847c5a0df06c8d1bd0df88ac1c73706b71e59aaa2052f376962eea5283b6fde227e80300000000000075","reqSigs":1,"type":"pubkeyhash","addresses":["15ZzkADmZxWf2fHK6xgkZgjGU9Ynj5JPDQ45Pc"]},"assets":[{"name":"4e9d9c77b69fabfd2c385afffabb1273","issuetxid":"27e2fdb68352ea2e9676f35220aa9ae5b6d30185bfceb9e64f5c708b1f41173c","assetref":null,"qty":10,"raw":1000,"type":"transfer"}]},{"value":0,"n":1,"scriptPubKey":{"asm":"OP_DUP OP_HASH160 fca3548c1a26d433964f713bada36ee66e66836a OP_EQUALVERIFY OP_CHECKSIG 73706b71e59aaa2052f376962eea5283b6fde227b882010000000000 OP_DROP","hex":"76a914fca3548c1a26d433964f713bada36ee66e66836a88ac1c73706b71e59aaa2052f376962eea5283b6fde227b88201000000000075","reqSigs":1,"type":"pubkeyhash","addresses":["1b9RYmRa1pXSWSPEycZEV3S8zUkCcpRxJUp2HY"]},"assets":[{"name":"4e9d9c77b69fabfd2c385afffabb1273","issuetxid":"27e2fdb68352ea2e9676f35220aa9ae5b6d30185bfceb9e64f5c708b1f41173c","assetref":null,"qty":990,"raw":99000,"type":"transfer"}]}]}`
)

func TestBuffers(t *testing.T) {

    t.Run(
        "test 7000000",
        func (t *testing.T) {

            x := varIntBuffer(70000000000)

            fmt.Printf("%x\n", x)

        },
    )

    t.Run(
        "test 70000",
        func (t *testing.T) {

            x := varIntBuffer(998000)

            fmt.Printf("%x\n", x)

        },
    )

    t.Run(
        "test 515",
        func (t *testing.T) {

            x := varIntBuffer(515)

            fmt.Printf("%x\n", x)

        },
    )

    t.Run(
        "test 128",
        func (t *testing.T) {

            x := varIntBuffer(106)

            fmt.Printf("%x\n", x)

        },
    )

    address.Configure(
        &address.Config{
            PrivateKeyVersion: "8097af59",
            AddressPubkeyhashVersion: "00ddc3a9",
            AddressChecksumValue: "8fdcaf22",
        },
    )

    t.Run(
        "test sig",
        func (t *testing.T) {
            seed := []byte("seed")

            keyPair, err := address.MultiChainWallet(seed, 2000, 0)
            if err != nil {
                t.Error(err)
                return
            }

            sigReq := &SigReq{
                transaction: &Transaction{},
                privateKey: keyPair.PrivateKey,
                publicKey: keyPair.PublicKey,
            }

            err = json.Unmarshal([]byte(CONST_TX_DECODED), sigReq.transaction)
            if err != nil {
                t.Error(err)
                return
            }

            b, _ := json.Marshal(sigReq.transaction)
            if err != nil {
                t.Error(err)
                return
            }

            fmt.Println(string(b))

            err = sigReq.Sign(0)
            if err != nil {
                t.Error(err)
                return
            }

            b, _ = json.Marshal(sigReq.transaction)
            if err != nil {
                t.Error(err)
                return
            }

            fmt.Println(string(b))
        },
    )
}
