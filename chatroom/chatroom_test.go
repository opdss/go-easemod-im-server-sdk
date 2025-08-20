package chatroom_test

import (
	"github.com/opdss/go-easemod-im-server-sdk/easemod"
	"github.com/opdss/go-easemod-im-server-sdk/request"
)

var em *easemod.EaseMod

func newEaseMod() *easemod.EaseMod {
	if em != nil {
		return em
	}
	em = easemod.NewEaseMod(request.Config{
		Endpoint:     "https://a71.easemob.com",
		OrgName:      "",
		AppName:      "",
		AppKey:       "",
		ClientId:     "",
		ClientSecret: "",
	})
	return em
}
