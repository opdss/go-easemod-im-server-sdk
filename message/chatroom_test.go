package message_test

import (
	"context"
	"fmt"
	"github.com/opdss/go-easemod-im-server-sdk/message"
	"testing"
)

func TestMessage_SendTxtToChatroom(t *testing.T) {

	data, err := newEaseMod().Message().SendTxtToChatroom(context.Background(), &message.SendToChatroomMsgReq[message.TxtMsg]{
		ChatroomMsgReq: message.ChatroomMsgReq{
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
