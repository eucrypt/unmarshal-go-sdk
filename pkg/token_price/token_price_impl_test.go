package token_price

import (
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	httpclient "github.com/eucrypt/unmarshal-go-sdk/pkg/http"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/session"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

//@dev before a run set your AUTH_KEY env
func TestPriceStoreV1_GetPrice(t1 *testing.T) {
	ps := getTestPriceStore()
	ast := assert.New(t1)
	validAddr := "0x41ab1b6fcbb2fa9dced81acbdec13ea6315f2bf2"
	chain := constants.ETH
	var time int64 = 1600173203

	t1.Run("Evaluate Get Price at instant", func(t *testing.T) {
		resp, err := ps.GetPriceAtInstant(chain, validAddr, time)

		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")

		resp, _ = ps.GetPriceAtInstant(chain, "invalidAddr", time)
		ast.Empty(resp, "should have an empty response for an invalid call")
	})

	t1.Run("Evaluate Get  Current Price", func(t *testing.T) {
		resp, err := ps.GetCurrentPrice(chain, validAddr)

		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")

		resp, _ = ps.GetCurrentPrice(chain, "invalidAddr")
		ast.Empty(resp, "should have an empty response for an invalid call")
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
	})
}

func TestPriceStoreV1_GetTokensPrice(t *testing.T) {
	ps := getTestPriceStore()
	ast := assert.New(t)
	validAddr := []string{"0x2fa5daf6fe0708fbd63b1a7d1592577284f52256", "0xad29abb318791d579433d831ed122afeaf29dcfe"}
	chain := constants.BSC
	t.Run("Evaluate PS_GetTokensPrice", func(t *testing.T) {
		resp, err := ps.GetTokensPrice(chain, validAddr)

		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")

		resp, _ = ps.GetTokensPrice(chain, nil)
		ast.Empty(resp, "should have an empty response for an invalid call")
	})
}

func TestPriceStoreV1_GetPriceWithSymbol(t *testing.T) {
	ps := getTestPriceStore()
	ast := assert.New(t)
	symbol := "marsh"
	t.Run("Evaluate GetPriceWithSymbol", func(t *testing.T) {
		resp, err := ps.GetPriceWithSymbol(symbol)

		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")

		resp, _ = ps.GetPriceWithSymbol("")
		ast.Empty(resp, "should have an empty response for an invalid call")
	})
}

func TestPriceStoreV1_GetLosers(t *testing.T) {
	ps := getTestPriceStore()
	ast := assert.New(t)
	chain := constants.ETH
	resp, err := ps.GetLosers(chain)

	ast.NoError(err, "There should be no error for a valid call")
	ast.NotEmpty(resp, "The response should not be empty")
}

func TestPriceStoreV1_GetGainers(t *testing.T) {
	ps := getTestPriceStore()
	ast := assert.New(t)
	chain := constants.ETH
	resp, err := ps.GetGainers(chain)

	ast.NoError(err, "There should be no error for a valid call")
	ast.NotEmpty(resp, "The response should not be empty")
}
