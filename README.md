# Golang client library for MultiChain blockchain

This library will allow you to complete a basic set of functions with a MultiChain node.

You should be able to issue, and send assets between addresses.

If you wish to contribute to flesh out the remaining API calls, please make pull requests.

## Testing

To fully test this package it is neccesary to have a full hot node running at the given parameters.

```

  chain := flag.String("chain", "", "is the name of the chain")
  host := flag.String("host", "localhost", "is a string for the hostname")
  port := flag.String("port", "80", "is a string for the host port")
  username := flag.String("username", "multichainrpc", "is a string for the username")
  password := flag.String("password", "12345678", "is a string for the password")

  flag.Parse()

  client := multichain.NewClient(
      *chain,
      *host,
      *port,
      *username,
      *password,
  )

  obj, err := client.GetInfo()
  if err != nil {
      panic(err)
  }

  fmt.Println(obj)

```
## Deterministic Wallets

Using the address package within this repo, you can create a deterministic keypair with WIF-encoded private-key, and Bitcoin or MultiChain encoded public address.

```

type KeyPair struct {
    Type string
    Index int
    Public string
    Private string
}

```

Here is an example of making a multichain HD wallet.

```
package main

import (
    "fmt"
    "gitlab.com/golangdaddy/multichain-client/address"
    "gitlab.com/golangdaddy/multichain-client/params"
)

const (
    CONST_BCRYPT_DIFF = 10
)

func main() {

    // you need to refer to your params.dat file to get the needed config parameters
    cfg, err := params.Open("./multichain-cold/params.dat")
    if err != nil {
        panic(err)
    }

    // The address package handles the encoding of keys, so it needs to be configued.
    address.Configure(&address.Config{
        PrivateKeyVersion: cfg.String("private-key-version"),
        AddressPubkeyhashVersion: cfg.String("address-pubkeyhash-version"),
        AddressChecksumValue: cfg.String("address-checksum-value"),
    })

    seed := []byte("seed")
    keyChildIndex := 0

    // create a new wallet
    keyPair, err := address.MultiChainWallet(seed, CONST_BCRYPT_DIFF, keyChildIndex)
    if err != nil {
        panic(err)
    }

    fmt.Println(keyPair)
}
```

If you have an existing private key, you can export it as WIF (Wallet Import Format) which can be used with importprivkey API command with MultiChain.

```

    wif, err := address.MultiChainWIF(privKeyBytes)

```

...or it's MultiChain address from the public key with the MultiChainAddress function.

```

    addr, err := address.MultiChainAddress(pubKeyBytes)

```
