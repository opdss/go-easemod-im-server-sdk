package chatroom_test

import (
	"context"
	"fmt"
	"github.com/opdss/go-easemod-im-server-sdk/chatroom"
	"testing"
)

func TestEaseModChatRoom_GetAnnouncement(t *testing.T) {
	data, err := newEaseMod().Chatroom().GetAnnouncement(context.Background(), "289695344951298")

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEaseModChatRoom_SetAnnouncement(t *testing.T) {
	data, err := newEaseMod().Chatroom().SetAnnouncement(context.Background(), "289695344951298", "888888的地盘")

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEaseModChatRoom_SetMetadata(t *testing.T) {
	data, err := newEaseMod().Chatroom().SetMetadata(context.Background(), "289695344951298", "888888", &chatroom.SetMetaDataReq{
		MetaData: map[string]string{
			"key":  "value",
			"key2": "value2",
		},
		AutoDelete: chatroom.NoDelete,
	})

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEaseModChatRoom_GetMetadata(t *testing.T) {
	data, err := newEaseMod().Chatroom().GetMetadata(context.Background(), "289695344951298", []string{"key", "key2", "key3"})

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEaseModChatRoom_DelMetadata(t *testing.T) {
	data, err := newEaseMod().Chatroom().DelMetadata(context.Background(), "289695344951298", "888888", []string{"key2", "key3"})

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEaseModChatRoom_ForceSetMetadata(t *testing.T) {
	data, err := newEaseMod().Chatroom().ForceSetMetadata(context.Background(), "289695344951298", "8888881", &chatroom.SetMetaDataReq{
		MetaData: map[string]string{
			"key":  "value",
			"key2": "value2",
		},
		AutoDelete: chatroom.NoDelete,
	})

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEaseModChatRoom_ForceDelMetadata(t *testing.T) {
	data, err := newEaseMod().Chatroom().DelMetadata(context.Background(), "289695344951298", "8888881", []string{"key2", "key3"})

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}
