package chatroom

import (
	"context"
	"github.com/opdss/go-easemod-im-server-sdk/request"
)

type Chatroom interface {
	// GetInfo 获取聊天室信息
	// https://doc.easemob.com/document/server-side/chatroom_manage.html#%E6%9F%A5%E8%AF%A2%E8%81%8A%E5%A4%A9%E5%AE%A4%E8%AF%A6%E6%83%85
	GetInfo(ctx context.Context, chatroomId string) (*InfoResp, error)
	// Create 创建聊天室
	// https://doc.easemod.com/document/server-side/chatroom_manage.html#%E5%88%9B%E5%BB%BA%E8%81%8A%E5%A4%A9%E5%AE%A4
	Create(ctx context.Context, req *CreateReq) (*CreateResp, error)
	// Update 修改聊天室信息
	// https://doc.easemob.com/document/server-side/chatroom_manage.html#%E4%BF%AE%E6%94%B9%E8%81%8A%E5%A4%A9%E5%AE%A4%E4%BF%A1%E6%81%AF
	Update(ctx context.Context, chatroomId string, req *UpdateReq) (*UpdateResp, error)
	// SetAnnouncement 设置聊天室公告
	// https://doc.easemob.com/document/server-side/chatroom_attribute.html#%E4%BF%AE%E6%94%B9%E8%81%8A%E5%A4%A9%E5%AE%A4%E5%85%AC%E5%91%8A
	SetAnnouncement(ctx context.Context, chatRoomId string, announcement string) (*SetAnnouncementResp, error)
	// GetAnnouncement 获取聊天室公告
	// https://doc.easemod.com/document/server-side/chatroom_manage.html#%E8%8E%B7%E5%8F%96%E8%81%8A%E5%A4%A9%E5%AE%A4%E5%86%99%E5%85%A5
	GetAnnouncement(ctx context.Context, chatRoomId string) (*GetAnnouncementResp, error)
	// SetMetadata 设置聊天室元数据
	// https://doc.easemob.com/document/server-side/chatroom_attribute.html#%E8%AE%BE%E7%BD%AE%E8%81%8A%E5%A4%A9%E5%AE%A4%E8%87%AA%E5%AE%9A%E4%B9%89%E5%B1%9E%E6%80%A7
	SetMetadata(ctx context.Context, chatRoomId string, username string, req *SetMetaDataReq) (*SetMetaDataResp, error)
	// GetMetadata 获取聊天室元数据
	// https://doc.easemob.com/document/server-side/chatroom_attribute.html#%E8%8E%B7%E5%8F%96%E8%81%8A%E5%A4%A9%E5%AE%A4%E8%87%AA%E5%AE%9A%E4%B9%89%E5%B1%9E%E6%80%A7
	GetMetadata(ctx context.Context, chatRoomId string, keys []string) (*GetMetaDataResp, error)
	// DelMetadata 删除聊天室元数据
	// https://doc.easemob.com/document/server-side/chatroom_attribute.html#%E5%88%A0%E9%99%A4%E8%81%8A%E5%A4%A9%E5%AE%A4%E8%87%AA%E5%AE%9A%E4%B9%89%E5%B1%9E%E6%80%A7
	DelMetadata(ctx context.Context, chatRoomId string, username string, keys []string) (*DelMetaDataResp, error)
	// ForceSetMetadata 强制设置聊天室元数据
	// https://doc.easemob.com/document/server-side/chatroom_attribute.html#%E5%BC%BA%E5%88%B6%E8%AE%BE%E7%BD%AE%E8%81%8A%E5%A4%A9%E5%AE%A4%E8%87%AA%E5%AE%9A%E4%B9%89%E5%B1%9E%E6%80%A7
	ForceSetMetadata(ctx context.Context, chatRoomId string, username string, req *SetMetaDataReq) (*SetMetaDataResp, error)
	// ForceDelMetadata 强制删除聊天室元数据
	// https://doc.easemob.com/document/server-side/chatroom_attribute.html#%E5%BC%BA%E5%88%B6%E5%88%A0%E9%99%A4%E8%81%8A%E5%A4%A9%E5%AE%A4%E8%87%AA%E5%AE%9A%E4%B9%89%E5%B1%9E%E6%80%A7
	ForceDelMetadata(ctx context.Context, chatRoomId string, username string, keys []string) (*DelMetaDataResp, error)

	// GetMembers 获取房间成员
	// https://doc.easemob.com/document/server-side/chatroom_member_obtain.html
	GetMembers(ctx context.Context, chatroomId string, page, pageSize int) (*MemberResp, error)
	// RemoveMember 移除房间成员
	// https://doc.easemob.com/document/server-side/chatroom_member_add_delete.html#%E7%A7%BB%E9%99%A4%E5%8D%95%E4%B8%AA%E8%81%8A%E5%A4%A9%E5%AE%A4%E6%88%90%E5%91%98
	RemoveMember(ctx context.Context, chatroomId string, username string) (*RemoveMemberResp, error)
}

type chatroom struct {
	client *request.Client
}

func NewChatroom(client *request.Client) Chatroom {
	return &chatroom{
		client: client,
	}
}

type CommonResp struct {
	Uri       string `json:"uri"`
	Timestamp int64  `json:"timestamp"`
	Action    string `json:"action"`
}
