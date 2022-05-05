package token_price

import (
	"fmt"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	httpclient "github.com/eucrypt/unmarshal-go-sdk/pkg/http"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/session"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/token_price/types"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"strconv"
	"testing"
)

//@dev before a run set your AUTH_KEY env
func TestPriceStoreV1_GetPrice(t1 *testing.T) {
	ps := getTestPriceStore()
	ast := assert.New(t1)
	validAddr := "0x41ab1b6fcbb2fa9dced81acbdec13ea6315f2bf2"
	chain := constants.ETH
	var time uint64 = 1600173203

	t1.Run("Evaluate Get Price at instant", func(t *testing.T) {
		resp, err := ps.GetTokenPrice(chain, validAddr, &types.GetPriceOptions{Timestamp: time})

		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")

		resp, _ = ps.GetTokenPrice(chain, "invalidAddr", nil)
		ast.Empty(resp, "should have an empty response for an invalid call")

		//@dev The chain below is currently unsupported. Test will be deprecated if the chain is ever supported
		_, err = ps.GetTokenPrice(constants.HUOBI, "", nil)
		ast.Equal(constants.UnsupportedChainError, err, "Call should result in an unsupported chain error")
	})

	t1.Run("Evaluate Get Current Price", func(t *testing.T) {
		resp, err := ps.GetTokenPrice(chain, validAddr, nil)

		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")
		fmt.Printf("response: %#v", resp)
	})
	t1.Run("Evaluate Get Price with options", func(t *testing.T) {
		resp, err := ps.GetTokenPrice(chain, validAddr, &types.GetPriceOptions{
			Timestamp:            time,
			TwentyFourHourChange: true,
			AlternateChain:       true,
		})

		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")
		ast.NotEmpty(resp.PriceChange, "The response should have price change data")
	})

}

func getTestPriceStore() PriceStoreImpl {
	httpClient := httpclient.NewHttpJSONClient(constants.Environment.GetEndpoint("prod"))
	authKey := os.Getenv("API_KEY")
	httpClient.DefaultQuery = map[string]string{"auth_key": authKey}
	PsV1 := New(session.Session{Config: struct {
		AuthKey     string
		HttpClient  *http.Client
		Environment constants.Environment
	}{AuthKey: authKey, HttpClient: nil, Environment: constants.Prod}, Client: httpClient})
	return PsV1
}

func TestPriceStoreV1_GetLPTokens(t *testing.T) {
	ps := getTestPriceStore()
	ast := assert.New(t)
	validAddr := "0x222F93187f15F354D41Ff6a7703eF7e18cdD5103"
	chain := constants.BSC
	t.Run("Evaluate GetLPToken", func(t *testing.T) {
		resp, err := ps.GetLPTokens(chain, validAddr)

		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")

		resp, _ = ps.GetLPTokens(chain, "")
		ast.Empty(resp, "should have an empty response for an invalid call")

		//@dev The chain below is currently unsupported. Test will be deprecated if the chain is ever supported
		_, err = ps.GetLPTokens(constants.HUOBI, "")
		ast.Equal(constants.UnsupportedChainError, err, "Call should result in an unsupported chain error")
	})
}

func TestPriceStoreV1_GetTokensPrice(t *testing.T) {
	ps := getTestPriceStore()
	ast := assert.New(t)
	validAddr := []string{"0x2fa5daf6fe0708fbd63b1a7d1592577284f52256", "0xad29abb318791d579433d831ed122afeaf29dcfe"}
	chain := constants.BSC
	t.Run("Evaluate PS_GetTokensPrice", func(t *testing.T) {
		resp, err := ps.GetMultipleTokenPrice(chain, validAddr)

		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")

		resp, _ = ps.GetMultipleTokenPrice(chain, nil)
		ast.Empty(resp, "should have an empty response for an invalid call")
	})
}

func TestPriceStoreV1_GetPriceWithSymbol(t *testing.T) {
	ps := getTestPriceStore()
	ast := assert.New(t)
	symbol := "marsh"
	t.Run("Evaluate GetTokenPriceBySymbol", func(t *testing.T) {
		resp, err := ps.GetTokenPriceBySymbol(symbol, nil)
		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")
		resp, _ = ps.GetTokenPriceBySymbol("", nil)
		ast.Empty(resp, "should have an empty response for an invalid call")
	})
	t.Run("Evaluate GetTokenPriceBySymbol with Options", func(t *testing.T) {
		options := types.GetPriceWithSymbolOptions{Timestamp: 1644045522}
		resp, err := ps.GetTokenPriceBySymbol(symbol, &options)

		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")

		resp, _ = ps.GetTokenPriceBySymbol("", nil)
		ast.Empty(resp, "should have an empty response for an invalid call")
	})
}

func TestPriceStoreV1_GetLosers(t *testing.T) {
	ps := getTestPriceStore()
	ast := assert.New(t)
	chain := constants.ETH
	t.Run("GetLosers with no options", func(t *testing.T) {
		resp, err := ps.GetTopLosers(chain, nil)

		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")

		//@dev The chain below is currently unsupported. Test will be deprecated if the chain is ever supported
		_, err = ps.GetTopLosers(constants.HUOBI, nil)
		ast.Equal(constants.UnsupportedChainError, err, "Call should result in an unsupported chain error")
	})

	t.Run("GetLosers with options", func(t *testing.T) {
		expectedMinimumPrice := 10.0
		options := types.GetTopLosersOptions{MinimumPrice: uint64(expectedMinimumPrice)}

		resp, err := ps.GetTopLosers(chain, &options)

		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")
		for _, tokenDetails := range resp {
			currPrice, err := strconv.ParseFloat(tokenDetails.CurrentPrice, 0)
			ast.NoError(err, "There should be no error while converting the price")
			ast.GreaterOrEqual(currPrice, expectedMinimumPrice,
				"The token's price should be greater than or equal to the expected minimum price")
		}

		//@dev The chain below is currently unsupported. Test will be deprecated if the chain is ever supported
		_, err = ps.GetTopLosers(constants.HUOBI, &options)
		ast.Equal(constants.UnsupportedChainError, err, "Call should result in an unsupported chain error")
	})

}

func TestPriceStoreV1_GetGainers(t *testing.T) {
	ps := getTestPriceStore()
	ast := assert.New(t)
	chain := constants.ETH

	t.Run("GetGainers with no options", func(t *testing.T) {
		resp, err := ps.GetTopGainers(chain, nil)

		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")

		//@dev The chain below is currently unsupported. Test will be deprecated if the chain is ever supported
		_, err = ps.GetTopGainers(constants.HUOBI, nil)
		ast.Equal(constants.UnsupportedChainError, err, "Call should result in an unsupported chain error")
	})

	t.Run("GetGainers with options", func(t *testing.T) {
		expectedMinimumPrice := 10.0
		options := types.GetTopGainersOptions{MinimumPrice: uint64(expectedMinimumPrice)}

		resp, err := ps.GetTopGainers(chain, &options)

		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")
		for _, tokenDetails := range resp {
			currPrice, err := strconv.ParseFloat(tokenDetails.CurrentPrice, 0)
			ast.NoError(err, "There should be no error while converting the price")
			ast.GreaterOrEqual(currPrice, expectedMinimumPrice,
				"The token's price should be greater than or equal to the expected minimum price")
		}

		_, err = ps.GetTopGainers(constants.HUOBI, &options)
		ast.Equal(constants.UnsupportedChainError, err, "Call should result in an unsupported chain error")

	})
	t.Run("GetGainers with options", func(t *testing.T) {
		expectedMinimumPrice := 15.0
		options := types.GetTopGainersOptions{MinimumPrice: uint64(expectedMinimumPrice)}

		resp, err := ps.GetTopGainers(chain, &options)

		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")
		for _, tokenDetails := range resp {
			currPrice, err := strconv.ParseFloat(tokenDetails.CurrentPrice, 0)
			ast.NoError(err, "There should be no error while converting the price")
			ast.GreaterOrEqual(currPrice, expectedMinimumPrice,
				"The token's price should be greater than or equal to the expected minimum price")
		}

	})
}
