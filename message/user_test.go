package message_test

import (
	"context"
	"fmt"
	"github.com/opdss/go-easemod-im-server-sdk/message"
	"testing"
)

func TestMessage_SendTxtToUser(t *testing.T) {

	data, err := newEaseMod().Message().SendTxtToUser(context.Background(), &message.SendToUserMsgReq[message.TxtMsg]{
		UserMsgReq: message.UserMsgReq{
			To:   []string{"888888"},
			Type: message.Txt,
		},
		Body: message.TxtMsg{
			Msg: "xxxxx",
		},
	})

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)

}
