package easemod

import (
	"github.com/opdss/go-easemod-im-server-sdk/chatroom"
	"github.com/opdss/go-easemod-im-server-sdk/message"
	"github.com/opdss/go-easemod-im-server-sdk/request"
	"github.com/opdss/go-easemod-im-server-sdk/user"
)

type EaseMod struct {
	*request.Client
	chatroom chatroom.Chatroom
	user     user.User
	message  message.Message
}

func NewEaseMod(conf request.Config) *EaseMod {
	c := request.NewClient(conf)
	em := &EaseMod{
		Client:   c,
		chatroom: chatroom.NewChatroom(c),
		user:     user.NewUser(c),
		message:  message.NewMessage(c),
	}
	return em
}

func (em *EaseMod) User() user.User {
	return em.user
}

func (em *EaseMod) Chatroom() chatroom.Chatroom {
	return em.chatroom
}

func (em *EaseMod) Message() message.Message {
	return em.message
}
