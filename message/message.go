package message

import "github.com/opdss/go-easemod-im-server-sdk/request"

type Message interface {
}

type message struct {
	client *request.Client
}

func NewMessage(client *request.Client) Message {
	return &message{
		client: client,
	}
}
