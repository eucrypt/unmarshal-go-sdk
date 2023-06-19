# unmarshal-go-sdk

A Golang SDK to simplify access to Unmarshal APIs

**This project is currently in *BETA***

## Usage Guide

The first step involved is to `go get` this project.

```shell
go get github.com/eucrypt/unmarshal-go-sdk
```

Create an SDK type and pass it your auth key. (To generate an Auth key you will need to create a user account
at [Unmarshal Console](https://console.unmarshal.io))

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
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
)

func main() {
	sdk := unmarshal.NewWithConfig(conf.Config{
		AuthKey:     "<auth key>",
		Environment: constants.Prod,
	})

	resp, err := sdk.GetTokenPriceBySymbol("marsh")
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

	resp, err := sdk.GetTokenCurrentPrice(constants.BSC, "0x2fa5daf6fe0708fbd63b1a7d1592577284f52256")
}

```

## The SDK has support for the following unmarshal API:

### Price Store ([Docs](https://docs.unmarshal.io/openapi/core/tag/Price-Store/))

- Get Price
    - `v1/pricestore/chain/:chain/:address` (`GetTokenCurrentPrice`)
- Get Price at Instant
    - `v1/pricestore/chain/:chain/:address?timestamp=` (`GetTokenPriceAtInstant`)
- Get Price With Symbol
    - `v1/pricestore/:symbol` (`GetTokenPriceBySymbol`)
- Get Gainers
    - `v1/pricestore/chain/:chain/gainers` (`GetTopGainers`)
- Get Losers
    - `v1/pricestore/chain/:chain/losers` (`GetTopLosers`)
- Get LpTokens
    - `v1/pricestore/chain/:chain/lptokens` (`GetLPTokens`)
- Get Price of List of tokens
    - `v1/tokenstore/token/all` (`GetMultipleTokenPrice`)

### Token Details ([Docs](https://docs.unmarshal.io/openapi/core/tag/Token-Store/))

- Get Token With Contract
    - `v1/tokenstore/token/address/:address` (`GetTokenDetailsByContract`)
- Get Token With Symbol
    - `v1/tokenstore/token/symbol/:symbol` (`GetTokenDetailsBySymbol`)
- Get Paginated List Of Tokens
    - `v1/tokenstore/token/all` (`GetTokenList`)

### Assets API ([Docs](https://docs.unmarshal.io/openapi/core/tag/Wallet-APIs/))

- Get List of Assets
    - `v1/:chain/address/:address/assets` (`GetTokenAssets`)
- Get Profit and loss
    - `v2/:chain/address/:address/userData?contract=` (`GetProfitAndLoss`)

### NFT APIs ([Docs](https://docs.unmarshal.io/openapi/core/tag/NFTs/))

- Get NFT assets for an address
    - `v1/:chain/address/:address/nft-assets` (`GetNFTAssetsByAddress`)
- Get NFT Transactions by Address
    - `v1/:chain/address/:address/nft-transactions` (`GetNFTTransactionsByAddress`)
- Get NFT Metadata
    - `v1/:chain/address/:address/details?tokenId=` (`GetNFTDetailsByID`)
- Get NFT Holders using the NFT's Token ID
    - `v1/:chain/address/:address/nftholders?tokenId=` (`GetNFTHolderByID`)

### Transaction APIs ([Docs](https://docs.unmarshal.io/openapi/core/tag/Wallet-APIs/#tag/Wallet-APIs/operation/transaction-history-v-1))

- Get Transaction data for an address
    - `v1/:chain/address/:address/transactions?contract=&page=&pageSize=` (`GetTokenTxns`)
- Get Transactions data for an address V2/Get Transactions with Price data
    - `v2/:chain/address/:address/transactions?contract=&page=&pageSize=` (`GetTokenTxnsV2`)
- Get Transaction details by Transaction ID
    - `v1/:chain/transactions/:txID` (`GetTxnDetails`)

### Protocol APIs ([Docs](https://docs.unmarshal.io/openapi/core/tag/Price-Store/#tag/Price-Store/operation/price-for-lp-tokens))

- Get Protocol Positions for an address
    - `v2/protocols/:protocol/address/:address/positions` (`GetPositions`)
- Get Protocol Pairs
    - `v2/protocols/:protocol/pairs` (`GetPairs`)

## Api and Supported Chains

<table>
  <thead>
    <tr>
      <th>API Name</th>
      <th>Ethereum</th>
      <th>BSC</th>
      <th>Matic</th>
      <th>Avalanche</th>
      <th>Solana</th>
      <th>XDC</th>
      <th>Zilliqa</th>
      <th>Huobi</th>
      <th>Arbitrum</th>
      <th>Celo</th>
      <th>Fantom</th>
      <th>Klaytn</th>
      <th>Fuse</th>
      <th>Cronos</th>
      <th>Velas</th>
      <th>Moonbeam</th>
      <th>Metis</th>
      <th>Aurora</th>
      <th>Matic Supernet</th>
      <th>zkEVM</th>
      <th>Mantle Testnet</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>GetTokenAssets</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
    </tr>
    <tr>
      <td>GetProfitAndLoss</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>✅</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
    </tr>
    <tr>
      <td>GetTokenTxns</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>❌</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
    </tr>
    <tr>
      <td>GetTxnDetails</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>❌</td>
      <td>❌</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
    </tr>
    <tr>
      <td>GetTxnDetailsV2</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>❌</td>
      <td>✅</td>
      <td>❌</td>
      <td>❌</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
    </tr>
    <tr>
      <td>GetNFTAssets</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
    </tr>
    <tr>
      <td>GetTxns - NFT</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>✅</td>
      <td>❌</td>
      <td>❌</td>
    </tr>
    <tr>
      <td>GetDetailsByID</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>✅</td>
      <td>❌</td>
      <td>❌</td>
    </tr>
    <tr>
      <td>GetHoldersByID</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
    </tr>
    <tr>
      <td>GetPriceWithAddress</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>✅</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
    </tr> 
    <tr>
      <td>GetTokensPrice</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>✅</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
    </tr>
    <tr>
      <td>GetLpTokenPrice</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>✅</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
    </tr>
    <tr>
      <td>GetLosers</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>✅</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
    </tr>
    <tr>
      <td>GetGainers</td>
      <td>✅</td>
      <td>✅</td>
      <td>✅</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>✅</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
      <td>❌</td>
    </tr>
  </tbody>
</table>

Additionally, support for raw transaction details include:

`GetRawTransactionsForAddress` : `Ethereum: Mainnet`, `Ethereum: Rinkeby`, `BSC`, `BSC: Testnet`, `Polygon`,
`Polygon: Mumbai`, `Avalanche: Mainnet`, `Arbitrum: Mainnet`, `Celo: Mainnet`, `Fantom: Mainnet`, `Klaytn: Mainnet`,
`Fuse: Mainnet`, `Cronos: Mainnet`, `Velas: Mainnet`, `Moonbeam: Mainnet`, `Metis: Mainnet`, `Aurora: Mainnet`, 
`zkEVM: Mainnet`,`Mantle: Testnet`

`GetBulkTxnDetails` : `Ethereum: Mainnet`, `Ethereum: Rinkeby`, `BSC`, `BSC: Testnet`, `Polygon`,
`Polygon: Mumbai`, `Avalanche: Mainnet`, `Arbitrum: Mainnet`, `Celo: Mainnet`, `Fantom: Mainnet`, `Klaytn: Mainnet`,
`Fuse: Mainnet`, `Cronos: Mainnet`, `Velas: Mainnet`, `zkEVM: Mainnet`, `Mantle: Testnet`

This API includes more details and abstains from including price at the time of the transaction

