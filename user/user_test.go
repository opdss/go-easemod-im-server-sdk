package user_test

import (
	"github.com/opdss/go-easemod-im-server-sdk/easemod"
	"github.com/opdss/go-easemod-im-server-sdk/request"
)

var em *easemod.EaseMod

func newEaseMod() *easemod.EaseMod {
	if em != nil {
		return em
	}
	var err error
	em, err = easemod.NewEaseMod(request.Config{
		Endpoints:    []string{"http://baiaaaadu.com/cff", "https://a71.easemob.com"},
		OrgName:      "",
		AppName:      "",
		ClientId:     "",
		ClientSecret: "",
	})
	if err != nil {
		panic(err)
	}
	return em
}
