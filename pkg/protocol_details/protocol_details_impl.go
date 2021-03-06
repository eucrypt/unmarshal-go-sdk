package protocol_details

import (
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/protocol_details/types"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/session"
	"strings"
)

type ProtocolDetailsImpl struct {
	sess session.Session
}

func New(sess session.Session) ProtocolDetailsImpl {
	return ProtocolDetailsImpl{sess: sess}
}

//GetPairs fetches all the available tracks pairs of a given protocol.
//It accepts constants.Protocol as the input and returns the valid and tracked pairs
func (protoImpl ProtocolDetailsImpl) GetPairs(protocol constants.Protocol) (resp types.GetPairsResp, err error) {
	if !constants.PROTO_GetPairs.SupportsProtocol(protocol) {
		return types.GetPairsResp{}, constants.UnsupportedProtocolError
	}

	path := strings.Replace(constants.PROTO_GetPairs.GetURI(), ":protocol", protocol.String(), 1)
	err = protoImpl.sess.Client.Get(&resp, path, nil)

	return
}

//GetPositions fetches the positions an address holds wrt a particular protocol (constants.Protocol).
//It accepts an address and the protocol.
func (protoImpl ProtocolDetailsImpl) GetPositions(protocol constants.Protocol, address string) (
	resp types.GetPositionsResp, err error) {
	if !constants.PROTO_GetPositions.SupportsProtocol(protocol) {
		return types.GetPositionsResp{}, constants.UnsupportedProtocolError
	}

	path := strings.Replace(constants.PROTO_GetPositions.GetURI(), ":protocol", protocol.String(), 1)
	path = strings.Replace(path, ":address", address, 1)
	err = protoImpl.sess.Client.Get(&resp, path, nil)

	return
}
