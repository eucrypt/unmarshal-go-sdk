package protocol_details

import (
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	httpclient "github.com/eucrypt/unmarshal-go-sdk/pkg/http"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/session"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestProtocolDetailsImpl_GetPairs(t *testing.T) {
	protoObj := getTestProtocolDetailsImpl()
	ast := assert.New(t)

	t.Run("Evaluating Get Pairs with valid data", func(t *testing.T) {
		validProtocol := constants.PancakeswapV2

		resp, err := protoObj.GetPairs(validProtocol)

		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp.ProtocolPairs, "The response should not be empty")
	})
}

func TestProtocolDetailsImpl_GetPositions(t *testing.T) {
	protoObj := getTestProtocolDetailsImpl()
	ast := assert.New(t)

	t.Run("Evaluating Get Pairs with valid data", func(t *testing.T) {
		validProtocol := constants.UniswapV2
		validAddress := "demo.eth"

		resp, err := protoObj.GetPositions(validProtocol, validAddress)

		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp.Positions, "The response should have valid positions")
	})
	t.Run("Evaluating Get Pairs with invalid data", func(t *testing.T) {
		validProtocol := constants.PancakeswapV2
		validAddress := "demo.eth"

		resp, _ := protoObj.GetPositions(validProtocol, validAddress)

		//@dev at the time of writing, the address held 0 positions. This could change leading to a failed test
		ast.Len(resp.Positions, 0, "The number of valid positions must be 0")
	})

}
func getTestProtocolDetailsImpl() ProtocolDetailsImpl {
	httpClient := httpclient.NewHttpJSONClient(constants.Environment.GetEndpoint("prod"))
	authKey := os.Getenv("API_KEY")
	httpClient.DefaultQuery = map[string]string{"auth_key": authKey}
	details := New(session.Session{Config: struct {
		AuthKey     string
		HttpClient  *http.Client
		Environment constants.Environment
	}{AuthKey: authKey, HttpClient: nil, Environment: constants.Prod}, Client: httpClient})
	return details
}
