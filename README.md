# unmarshal-go-sdk

A Golang SDK to simplify access to Unmarshal APIs

**This project is currently in *BETA* and has limited support**

## Usage Guide

The first step involved is to `go get` this project.

```shell
go get github.com/eucrypt/unmarshal-go-sdk
```

Create an SDK type and pass it your auth key. (To generate an Auth key you will need to create a user account at (
console.unmarshal.io))

```go
package main

import (
  unmarshal "github.com/eucrypt/unmarshal-go-sdk/pkg
  conf "github.com/eucrypt/unmarshal-go-sdk/pkg/config"
)

func main() {
  sdk := unmarshal.NewWithConfig(conf.Config{
    AuthKey:     "<auth key>",
    Environment: constants.Prod,
  })
}
```

There are other options to use your own `http` client if you prefer as well. The SDK is all you now need to query the
unmarshal APIs

```go
//For example to get the current Price of Marsh, you can now:
package main

import (
  unmarshal "github.com/eucrypt/unmarshal-go-sdk/pkg"
  conf "github.com/eucrypt/unmarshal-go-sdk/pkg/config"
)

func main() {
  sdk := unmarshal.NewWithConfig(conf.Config{
    AuthKey:     "<auth key>",
    Environment: constants.Prod,
  })

  resp, err := sdk.GetPriceWithSymbol("marsh")
}

```

Some functions are chain specific and the chain can be passed in via the `constants.Chain` package

```go
package main

import (
  unmarshal "github.com/eucrypt/unmarshal-go-sdk/pkg"
  conf "github.com/eucrypt/unmarshal-go-sdk/pkg/config"
  "github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
)

func main() {
  sdk := unmarshal.NewWithConfig(conf.Config{
    AuthKey:     "<auth key>",
    Environment: constants.Prod,
  })

  resp, err := sdk.GetCurrentPrice(constants.BSC, "0x222F93187f15F354D41Ff6a7703eF7e18cdD5103")
}

```

## The SDK has support for the following unmarshal API:

### Price Store ([Docs](https://docs.unmarshal.io))

- Get Price
  - (`GetCurrentPrice`)
- Get Price at Instant
  - (`GetPriceAtInstant`)
- Get Price With Symbol
  - (`GetPriceWithSymbol`)
- Get Gainers
  - (`GetGainers`)
- Get Losers
  - (`GetLosers`)
- Get LpTokens
  - (`GetLPTokens`)
- Get Price of List of tokens
  - (`GetTokensPrice`)

### Token Details ([Docs](https://docs.unmarshal.io/token-store))

- Get Token With Contract
  - `v1/tokenstore/token/address/:address` (`GetDetailsWithContract`)
- Get Token With Symbol
  - `v1/tokenstore/token/symbol/:symbol` (`GetTokenWithSymbol`)
- Get Paginated List Of Tokens
  - `v1/tokenstore/token/all` (`GetTokenList`)

### Assets API ([Docs](https://docs.unmarshal.io/unmarshal-apis/token-balance-apis))

- Get List of Assets - `v1/:chain/address/:address/assets` (`GetAssets`)

### NFT APIs ([Docs](https://docs.unmarshal.io/nft-apis))

- Get NFT assets for an address
  - `v1/:chain/address/:address/nft-assets` (`GetAssetsByAddress`)
- Get NFT Transactions by Address
  - `v1/:chain/address/:address/nft-transactions` (`GetTransactionsByAddress`)
- Get NFT Metadata
  - `v1/:chain/address/:address/details?tokenId=` (`GetDetailsByID`)
- Get NFT Holders using the NFT's Token ID
  - `v1/:chain/address/:address/nftholders?tokenId=` (`GetHolderByID`)

### Transaction APIs ([Docs](https://docs.unmarshal.io/supported-networks))

- Get Transaction data for an address
  - `v1/:chain/address/:address/transactions?contract=&page=&pageSize=` (`GetTokenTxns`)
- Get Transactions data for an address V2/Get Transactions with Price data
  - `v2/:chain/address/:address/transactions?contract=&page=&pageSize=` (`GetTokenTxnsV2`)
- Get Transaction details by Transaction ID
  - `v1/:chain/transactions/:txID` (`GetTxnDetails`)


