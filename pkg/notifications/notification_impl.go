package notifications

import (
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/notifications/types"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/session"
)

type NotificationImpl struct {
	sess session.Session
}

func (n NotificationImpl) TriggerNotification(req types.CommunicationReq) error {
	var resp interface{}
	return n.sess.Client.Post(&resp, constants.NOTIFY_TriggerNotification.GetURI(), req)
}
