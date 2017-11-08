# Golang client library for MultiChain blockchain

This library will allow you to complete a basic set of functions with either the hot or cold MultiChain nodes.

You should be able to issue, and send assets between addresses.

## Usage

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

Using the address package within this repo, you can create a deterministic keypair with WIF encoded private-key, and Bitcoin or MultiChain encoded public key.
