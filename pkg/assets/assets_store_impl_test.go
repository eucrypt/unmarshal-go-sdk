package assets

import (
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	httpclient "github.com/eucrypt/unmarshal-go-sdk/pkg/http"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/session"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestV1Store_GetAssets(t *testing.T) {
	ps := getTestPriceStore()
	ast := assert.New(t)
	validAddr := "demo.eth"
	chain := constants.ETH
	t.Run("Evaluate Get  Current Price", func(t *testing.T) {
		resp, err := ps.GetTokenAssets(chain, validAddr)

		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")

		resp, _ = ps.GetTokenAssets(chain, "invalidAddr")
		ast.Empty(resp, "should have an empty response for an invalid call")
	})
}

func getTestPriceStore() V1Store {
	httpClient := httpclient.NewHttpJSONClient(constants.Environment.GetEndpoint("prod"))
	authKey := os.Getenv("API_KEY")
	httpClient.DefaultQuery = map[string]string{"auth_key": authKey}
	assetsStore := New(session.Session{Config: struct {
		AuthKey     string
		HttpClient  *http.Client
		Environment constants.Environment
	}{AuthKey: authKey, HttpClient: nil, Environment: constants.Prod}, Client: httpClient})
	return assetsStore
}
