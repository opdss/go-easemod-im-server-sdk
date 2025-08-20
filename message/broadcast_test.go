package message_test

import (
	"context"
	"fmt"
	"github.com/opdss/go-easemod-im-server-sdk/message"
	"testing"
)

func TestMessage_Broadcast(t *testing.T) {

	data, err := newEaseMod().Message().BroadcastTxt(context.Background(), &message.BroadcastMsgReq[message.TxtMsg]{
		TargetType: "users",
		Appkey:     "aaaa",
		Msg: message.BroadcastMsg[message.TxtMsg]{
			Type: message.Txt,
			Msg: message.TxtMsg{
				Msg: "test111",
			},
		},
	})

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)

}
