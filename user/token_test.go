package user_test

import (
	"context"
	"fmt"
	"github.com/opdss/go-easemod-im-server-sdk/user"
	"testing"
)

func TestEaseModUser_Register(t *testing.T) {
	data, err := newEaseMod().User().Register(context.Background(), &user.RegisterReq{Username: "8888881", Password: "123456"})

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEaseModUser_BatchRegistry(t *testing.T) {
	data, err := newEaseMod().User().BatchRegistry(context.Background(), []*user.RegisterReq{
		&user.RegisterReq{Username: "8888881", Password: "123456"},
		&user.RegisterReq{Username: "9999994", Password: "123456"}})

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEaseModUser_GetToken(t *testing.T) {
	data, err := newEaseMod().User().GetToken(context.Background(), &user.TokenReq{
		Username:       "888888341",
		GrantType:      "inherit",
		AutoCreateUser: false,
		Ttl:            86400,
	})

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}
