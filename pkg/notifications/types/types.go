package types

type CommunicationReq struct {
	Recipient struct {
		Active     bool   `json:"active"`
		Channel    string `json:"channel"`
		FcmToken   string `json:"fcm_token"`
		WebhookUrl string `json:"webhook_url"`
	} `json:"recipient"`
	Data map[string]interface{} `json:"data"`
}
