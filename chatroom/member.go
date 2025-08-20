package chatroom

import (
	"context"
	"github.com/opdss/go-easemod-im-server-sdk/request"
	"strconv"
)

type MemberResp struct {
	Action      string `json:"action"`
	Application string `json:"application"`
	Params      struct {
		PageSize []string `json:"pagesize"`
		PageNum  []string `json:"pagenum"`
	} `json:"params"`
	Uri             string              `json:"uri"`
	Entities        []interface{}       `json:"entities"`
	Data            []map[string]string `json:"data"`
	Timestamp       int64               `json:"timestamp"`
	Duration        int                 `json:"duration"`
	Organization    string              `json:"organization"`
	ApplicationName string              `json:"applicationName"`
	Count           int                 `json:"count"`
}

func (c *chatroom) GetMembers(ctx context.Context, chatroomId string, page, pageSize int) (*MemberResp, error) {
	resp := MemberResp{}
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 1000 {
		pageSize = 1000
	}
	err := c.client.Get(ctx, "/chatrooms/"+chatroomId+"/users?pagenum="+strconv.Itoa(page)+"&pagesize="+strconv.Itoa(pageSize), &resp)
	return &resp, err
}

type RemoveMemberResp struct {
	request.CommonResp
	Data struct {
		Result bool   `json:"result"`
		Action string `json:"action"`
		User   string `json:"user"`
		Id     string `json:"id"`
	} `json:"data"`
}

func (c *chatroom) RemoveMember(ctx context.Context, chatroomId string, username string) (*RemoveMemberResp, error) {
	resp := RemoveMemberResp{}
	err := c.client.Delete(ctx, "/chatrooms/"+chatroomId+"/users/"+username, nil, &resp)
	return &resp, err
}
