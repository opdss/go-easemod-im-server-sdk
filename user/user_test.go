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
	em = easemod.NewEaseMod(request.Config{
		Endpoint:     "https://a71.easemob.com",
		OrgName:      "1104241128170445",
		AppName:      "sout-test",
		AppKey:       "1104241128170445#sout-test",
		ClientId:     "YXA6OIEyvWabTTyhZBY8tlRRUg",
		ClientSecret: "YXA6iyTY12_6S1XehnmOjqF2DZBWCFg",
	})
	return em
}
