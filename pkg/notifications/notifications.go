package notifications

import "github.com/eucrypt/unmarshal-go-sdk/pkg/notifications/types"

type Notification interface {
	TriggerNotification(req types.CommunicationReq) error
}
