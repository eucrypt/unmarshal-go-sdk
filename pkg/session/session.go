package session

import (
	conf "github.com/eucrypt/unmarshal-go-sdk/pkg/config"
	httpclient "github.com/eucrypt/unmarshal-go-sdk/pkg/http"
)

type Session struct {
	Config conf.Config
	Client httpclient.Request
}
