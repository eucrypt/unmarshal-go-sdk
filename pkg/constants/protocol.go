package constants

import "errors"

type Protocol string

const (
	PancakeswapV1 Protocol = "pancakeswap_v1"
	PancakeswapV2 Protocol = "pancakeswap_v2"
	UniswapV2     Protocol = "uniswap_v2"
)

var allProtocols = map[Protocol]bool{
	PancakeswapV1: true,
	PancakeswapV2: true,
	UniswapV2:     true,
}
var allowedCallersByProtocol = map[APIName]map[Protocol]bool{
	PROTO_GetPairs:     allProtocols,
	PROTO_GetPositions: allProtocols,
}

//String returns the string specific version of the protocol
func (c Protocol) String() string {
	return string(c)
}

var UnsupportedProtocolError = errors.New("unsupported protocol for API")
