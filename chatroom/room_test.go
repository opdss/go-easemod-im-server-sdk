package chatroom_test

import (
	"context"
	"fmt"
	"github.com/opdss/go-easemod-im-server-sdk/chatroom"
	"testing"
)

func TestEaseModChatRoom_GetInfo(t *testing.T) {
	data, err := newEaseMod().Chatroom().GetInfo(context.Background(), "289695344951298")

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEaseModChatRoom_Create(t *testing.T) {
	data, err := newEaseMod().Chatroom().Create(context.Background(), &chatroom.CreateReq{
		Name:        "chat room 888888",
		Description: "chatroom Description",
		MaxUsers:    10000,
		Owner:       "888888",
		Members:     []string{"8888881"},
		Custom:      "chatroom Custom",
	})

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEaseModChatRoom_Update(t *testing.T) {
	data, err := newEaseMod().Chatroom().Update(context.Background(), "289695344951298", &chatroom.UpdateReq{
		Name:        "chat room 289695344951298",
		Description: "chatroom Description",
		MaxUsers:    8888,
	})

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}
