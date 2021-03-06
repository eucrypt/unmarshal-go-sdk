package nft_details

import (
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	httpclient "github.com/eucrypt/unmarshal-go-sdk/pkg/http"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/session"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestNFTDetailsImpl_GetAssetsByAddress(t *testing.T) {
	nftObj := getTestNFTObj()
	ast := assert.New(t)
	t.Run("Evaluating Get Assets by Address", func(t *testing.T) {
		validAddr := "demo.eth"
		chain := constants.ETH

		resp, err := nftObj.GetNFTAssetsByAddress(chain, validAddr)

		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")

		resp, _ = nftObj.GetNFTAssetsByAddress(chain, "")
		ast.Empty(resp, "should have an empty response for an invalid call")

		//@dev The chain below is currently unsupported. Test will be deprecated if the chain is ever supported
		_, err = nftObj.GetNFTAssetsByAddress(constants.HUOBI, "")
		ast.Equal(constants.UnsupportedChainError, err, "Call should result in an unsupported chain error")

	})
	t.Run("Evaluating Get Assets by Address on Avalanche", func(t *testing.T) {
		validAddr := "0x59DD353A69e2370ca98C1ff32587131b779a587a"
		chain := constants.AVALANCHE

		resp, err := nftObj.GetNFTAssetsByAddress(chain, validAddr)

		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")

	})
}

func getTestNFTObj() NFTDetailsImpl {
	httpClient := httpclient.NewHttpJSONClient(constants.Environment.GetEndpoint("prod"))
	authKey := os.Getenv("API_KEY")
	httpClient.DefaultQuery = map[string]string{"auth_key": authKey}
	NftImpl := New(session.Session{Config: struct {
		AuthKey     string
		HttpClient  *http.Client
		Environment constants.Environment
	}{AuthKey: authKey, HttpClient: nil, Environment: constants.Prod}, Client: httpClient})
	return NftImpl
}

func TestNFTDetailsImpl_GetDetailsByID(t *testing.T) {
	nftObj := getTestNFTObj()
	ast := assert.New(t)

	t.Run("Evaluating Get NFT Details by ID", func(t *testing.T) {
		validChain := constants.ETH
		validTokenId := "61"
		validNFTAddr := "0x3cf8695c5cb6caa78d9c7fc9fa34bc8271483a1a"

		resp, err := nftObj.GetNFTDetailsByID(validChain, validTokenId, validNFTAddr)

		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")

		resp, _ = nftObj.GetNFTDetailsByID(validChain, "", "")
		ast.Empty(resp, "should have an empty response for an invalid call")

		//@dev The chain below is currently unsupported. Test will be deprecated if the chain is ever supported
		_, err = nftObj.GetNFTDetailsByID(constants.HUOBI, "", "")
		ast.Equal(constants.UnsupportedChainError, err, "Call should result in an unsupported chain error")
	})
	t.Run("Evaluating Get NFT Details by ID 2", func(t *testing.T) {
		validChain := constants.ETH
		validTokenId := "1300020038"
		validNFTAddr := "0x4629122c04eacc2ca48bda4a92aadcaee5d15389"

		resp, err := nftObj.GetNFTDetailsByID(validChain, validTokenId, validNFTAddr)

		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")

	})
}

func TestNFTDetailsImpl_GetHolderByID(t *testing.T) {
	nftObj := getTestNFTObj()
	ast := assert.New(t)

	t.Run("Evaluating Get NFT Holders by Token ID", func(t *testing.T) {
		validChain := constants.ETH
		validTokenId := "61"
		validNFTAddr := "0x3cf8695c5cb6caa78d9c7fc9fa34bc8271483a1a"

		resp, err := nftObj.GetNFTHolderByID(validChain, validTokenId, validNFTAddr)

		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")

		resp, _ = nftObj.GetNFTHolderByID(validChain, "", "")
		ast.Empty(resp, "should have an empty response for an invalid call")

		//@dev The chain below is currently unsupported. Test will be deprecated if the chain is ever supported
		_, err = nftObj.GetNFTHolderByID(constants.HUOBI, "", "")
		ast.Equal(constants.UnsupportedChainError, err, "Call should result in an unsupported chain error")
	})
	t.Run("Evaluating Get NFT Holders by Token ID 2", func(t *testing.T) {
		validChain := constants.ETH
		validTokenId := "1300020038"
		validNFTAddr := "0x4629122c04eacc2ca48bda4a92aadcaee5d15389"

		resp, err := nftObj.GetNFTHolderByID(validChain, validTokenId, validNFTAddr)

		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")

	})
}

func TestNFTDetailsImpl_GetTransactionsByAddress(t *testing.T) {
	nftObj := getTestNFTObj()
	ast := assert.New(t)

	t.Run("Evaluating Get NFT Transactions By Address", func(t *testing.T) {
		validAddr := "0x39fcb954d0535befe1b0f52aea79ca2ee1ddf54e"
		validChain := constants.ETH
		pageNumber := 1
		pageSize := 5

		resp, err := nftObj.GetNFTTransactionsByAddress(validChain, validAddr, pageNumber, pageSize)
		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")
		ast.Len(resp, 5, "Exactly 5 objects should be a part of the response")

		//@dev The chain below is currently unsupported. Test will be deprecated if the chain is ever supported
		_, err = nftObj.GetNFTTransactionsByAddress(constants.HUOBI, "", 0, 0)
		ast.Equal(constants.UnsupportedChainError, err, "Call should result in an unsupported chain error")
	})
	t.Run("Evaluating Get NFT Transactions By Address 2", func(t *testing.T) {
		validAddr := "demo.eth"
		validChain := constants.ETH
		pageNumber := 1
		pageSize := 10

		resp, err := nftObj.GetNFTTransactionsByAddress(validChain, validAddr, pageNumber, pageSize)
		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")
		ast.Len(resp, 10, "Exactly 10 objects should be a part of the response")

	})

	t.Run("Evaluating Get NFT Transactions By Address 3", func(t *testing.T) {
		validAddr := "0x59DD353A69e2370ca98C1ff32587131b779a587a"
		validChain := constants.AVALANCHE
		pageNumber := 1
		pageSize := 10

		resp, err := nftObj.GetNFTTransactionsByAddress(validChain, validAddr, pageNumber, pageSize)
		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")
		ast.Len(resp, 10, "Exactly 10 objects should be a part of the response")

	})
}
