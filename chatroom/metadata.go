package chatroom

import (
	"context"
	"fmt"
	"github.com/opdss/go-easemod-im-server-sdk/request"
)

type AnnouncementData struct {
	Announcement string `json:"announcement"`
}

type GetAnnouncementResp struct {
	request.CommonResp
	Data struct {
		Announcement string `json:"announcement"`
	} `json:"data"`
}

type SetAnnouncementResp struct {
	request.CommonResp
	Data struct {
		Id     string `json:"id"`
		Result bool   `json:"result"`
	} `json:"data"`
}

func (c *chatroom) GetAnnouncement(ctx context.Context, chatRoomId string) (*GetAnnouncementResp, error) {
	resp := GetAnnouncementResp{}
	err := c.client.Get(ctx, fmt.Sprintf("/chatrooms/%s/announcement", chatRoomId), &resp)
	return &resp, err
}

func (c *chatroom) SetAnnouncement(ctx context.Context, chatRoomId string, announcement string) (*SetAnnouncementResp, error) {
	resp := SetAnnouncementResp{}
	err := c.client.Post(ctx, fmt.Sprintf("/chatrooms/%s/announcement", chatRoomId), map[string]string{"announcement": announcement}, &resp)
	return &resp, err
}

const (
	Delete   = "DELETE"
	NoDelete = "NO_DELETE"
)

type MetaData map[string]string
type SetMetaDataReq struct {
	MetaData   MetaData `json:"metaData"`
	AutoDelete string   `json:"autoDelete"` //当前成员退出聊天室时是否自动删除该自定义属性。• （默认）'DELETE'：是；• 'NO_DELETE'：否。
}

type MetaDataOperatorResult struct {
	SuccessKeys []string `json:"successKeys"`
	ErrorKeys   MetaData `json:"errorKeys"`
}

type SetMetaDataResp struct {
	CommonResp
	Data MetaDataOperatorResult `json:"data"`
}

type GetMetaDataResp struct {
	CommonResp
	Data MetaData `json:"data"`
}

type DelMetaDataResp struct {
	CommonResp
	Data MetaDataOperatorResult `json:"data"`
}

func (c *chatroom) SetMetadata(ctx context.Context, chatRoomId string, username string, req *SetMetaDataReq) (*SetMetaDataResp, error) {
	resp := SetMetaDataResp{}
	if req.AutoDelete == "" {
		req.AutoDelete = Delete
	}
	err := c.client.Put(ctx, fmt.Sprintf("/metadata/chatroom/%s/user/%s", chatRoomId, username), req, &resp)
	return &resp, err
}

func (c *chatroom) GetMetadata(ctx context.Context, chatRoomId string, keys []string) (*GetMetaDataResp, error) {
	resp := GetMetaDataResp{}
	err := c.client.Post(ctx, fmt.Sprintf("/metadata/chatroom/%s", chatRoomId), map[string][]string{"keys": keys}, &resp)
	return &resp, err
}

func (c *chatroom) DelMetadata(ctx context.Context, chatRoomId string, username string, keys []string) (*DelMetaDataResp, error) {
	resp := DelMetaDataResp{}
	err := c.client.Delete(ctx, fmt.Sprintf("/metadata/chatroom/%s/user/%s", chatRoomId, username), map[string][]string{"keys": keys}, &resp)
	return &resp, err
}

func (c *chatroom) ForceSetMetadata(ctx context.Context, chatRoomId string, username string, req *SetMetaDataReq) (*SetMetaDataResp, error) {
	resp := SetMetaDataResp{}
	if req.AutoDelete == "" {
		req.AutoDelete = Delete
	}
	err := c.client.Put(ctx, fmt.Sprintf("/metadata/chatroom/%s/user/%s/forced", chatRoomId, username), req, &resp)
	return &resp, err
}

func (c *chatroom) ForceDelMetadata(ctx context.Context, chatRoomId string, username string, keys []string) (*DelMetaDataResp, error) {
	resp := DelMetaDataResp{}
	err := c.client.Delete(ctx, fmt.Sprintf("/metadata/chatroom/%s/user/%s/forced", chatRoomId, username), map[string][]string{"keys": keys}, &resp)
	return &resp, err
}
