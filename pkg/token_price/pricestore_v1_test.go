package token_price

import (
	"github.com/eucrypt/unmarshal-go-sdk/pkg"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	httpclient "github.com/eucrypt/unmarshal-go-sdk/pkg/http"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestPriceStoreV1_GetPriceAtInstant(t1 *testing.T) {
	ps := getTestPriceStore()
	resp, err := ps.GetPriceAtInstant("0x41ab1b6fcbb2fa9dced81acbdec13ea6315f2bf2", constants.ETH,
		1600173203)
	assert.NoError(t1, err, "There should be no error for a valid call")
	assert.NotEmpty(t1, resp, "The response should not be empty")

}

func getTestPriceStore() PriceStoreV1 {
	httpClient := httpclient.NewHttpJSONClient(constants.Environment.GetEndpoint("prod"))
	authKey := os.Getenv("API_KEY")
	httpClient.DefaultQuery = map[string]string{"auth_key": authKey}
	PsV1 := New(pkg.Session{Config: struct {
		AuthKey     string
		HttpClient  *http.Client
		Environment constants.Environment
	}{AuthKey: authKey, HttpClient: nil, Environment: constants.Prod}, Client: httpClient})
	return PsV1
}
