package chatroom_test

import (
	"context"
	"fmt"
	"testing"
)

func TestEaseModChatRoom_GetMembers(t *testing.T) {
	data, err := newEaseMod().Chatroom().GetMembers(context.Background(), "289695344951298", 1, 220)
	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEaseModChatRoom_RemoveMember(t *testing.T) {
	data, err := newEaseMod().Chatroom().RemoveMember(context.Background(), "289695344951298", "8888881")
	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}
