package protocol_details

import (
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/protocol_details/types"
)

type ProtocolDetails interface {
	GetPairs(protocol constants.Protocol) (resp types.GetPairsResp, err error)
	GetPositions(details constants.Protocol, address string) (resp types.GetPositionsResp, err error)
}
