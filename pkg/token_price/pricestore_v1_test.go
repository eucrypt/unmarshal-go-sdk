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

func TestPriceStoreV1_GetPrice(t1 *testing.T) {
	ps := getTestPriceStore()
	ast := assert.New(t1)
	validAddr := "0x41ab1b6fcbb2fa9dced81acbdec13ea6315f2bf2"
	chain := constants.ETH
	var time int64 = 1600173203

	t1.Run("Evaluate Get Price at instant", func(t *testing.T) {
		resp, err := ps.GetPriceAtInstant(validAddr, chain, time)

		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")

		resp, _ = ps.GetPriceAtInstant("invalidAddr", chain, time)
		ast.Empty(resp, "should have an empty response for an invalid call")
	})

	t1.Run("Evaluate Get  Current Price", func(t *testing.T) {
		resp, err := ps.GetCurrentPrice(validAddr, chain)

		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")

		resp, _ = ps.GetCurrentPrice("invalidAddr", chain)
		ast.Empty(resp, "should have an empty response for an invalid call")
	})

}

func getTestPriceStore() PriceStoreV1 {
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
