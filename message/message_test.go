package message_test

import (
	"github.com/opdss/go-easemod-im-server-sdk/easemod"
	"github.com/opdss/go-easemod-im-server-sdk/request"
)

var em *easemod.EaseMod

func newEaseMod() *easemod.EaseMod {
	if em != nil {
		return em
	}
	em, _ = easemod.NewEaseMod(request.Config{
		Endpoints:    []string{"https://a71.easemob.com"},
		OrgName:      "",
		AppName:      "",
		ClientId:     "",
		ClientSecret: "",
	})
	return em
}
