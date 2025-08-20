package chatroom

import (
	"context"
	"github.com/opdss/go-easemod-im-server-sdk/request"
)

type Info struct {
	Id                string              `json:"id"`
	Name              string              `json:"name"`
	Description       string              `json:"description"`
	MembersOnly       bool                `json:"membersonly"`
	AllowInvites      bool                `json:"allowinvites"`
	MaxUsers          int                 `json:"maxusers"`
	Owner             string              `json:"owner"`
	Created           int64               `json:"created"`
	Custom            string              `json:"custom"`
	AffiliationsCount int                 `json:"affiliations_count"`
	Affiliations      []map[string]string `json:"affiliations"`
	Public            bool                `json:"public"`
	Mute              bool                `json:"mute"`
}
type InfoResp struct {
	request.CommonResp
	Data []Info `json:"data"`
}

func (c *chatroom) GetInfo(ctx context.Context, chatroomId string) (*InfoResp, error) {
	resp := InfoResp{}
	err := c.client.Get(ctx, "/chatrooms/"+chatroomId, &resp)
	return &resp, err
}

type CreateReq struct {
	Name        string   `json:"name" validate:"required,max=128"`
	Description string   `json:"description" validate:"required,max=512"`
	MaxUsers    int      `json:"maxusers,omitempty" validate:"omitempty,min=1,max=10000"`
	Owner       string   `json:"owner" validate:"required"`
	Members     []string `json:"members,omitempty" validate:"omitempty,min=1,dive,required"`
	Custom      string   `json:"custom,omitempty" validate:"omitempty,max=8192"`
}

type CreateResp struct {
	request.CommonResp
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}

// Create 创建聊天室
// @document https://doc.easemob.com/document/server-side/chatroom_manage.html#%E5%88%9B%E5%BB%BA%E8%81%8A%E5%A4%A9%E5%AE%A4
func (c *chatroom) Create(ctx context.Context, req *CreateReq) (*CreateResp, error) {
	resp := CreateResp{}
	err := c.client.Post(ctx, "/chatrooms", req, &resp)
	return &resp, err
}

type UpdateReq struct {
	Name        string `json:"name" validate:"required,max=128"`
	Description string `json:"description" validate:"required,max=512"`
	MaxUsers    int    `json:"maxusers,omitempty" validate:"omitempty,min=1,max=10000"`
}

type UpdateResp struct {
	request.CommonResp
	Data struct {
		Description bool `json:"description"`
		MaxUsers    bool `json:"maxusers"`
		GroupName   bool `json:"groupname"`
	} `json:"data"`
}

func (c *chatroom) Update(ctx context.Context, chatroomId string, req *UpdateReq) (*UpdateResp, error) {
	resp := UpdateResp{}
	err := c.client.Put(ctx, "/chatrooms/"+chatroomId, req, &resp)
	return &resp, err
}
