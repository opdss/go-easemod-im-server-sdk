package message

import (
	"context"
)

type ChatroomMsgReq struct {
	From             *string   `json:"from,omitempty"` //消息发送方的用户 ID。若不传入该字段，服务器默认设置为 admin。
	To               []string  `json:"to"`
	ChatroomMsgLevel *MsgLevel `json:"chatroom_msg_level,omitempty"` // 消息级别，默认值为 NORMAL
	Type             Type      `json:"type"`                         // 消息类型
	RoamIgnoreUsers  []string  `json:"roam_ignore_users,omitempty"`  // 设置哪些用户拉漫游消息时拉不到该消息。
	Ext              *Ext      `json:"ext,omitempty"`                //消息支持扩展字段，可添加自定义信息。不能对该参数传入 null。同时，推送通知也支持自定义扩展字段，详见 APNs 自定义显示 和 Android 推送字段说明。
}

type SendToChatroomMsgReq[T MsgBody] struct {
	ChatroomMsgReq
	Body T `json:"body"`
}

func (m *message) SendTxtToChatroom(ctx context.Context, req *SendToChatroomMsgReq[TxtMsg]) (*SendMsgResp, error) {
	return m.SendToChatroom(ctx, req)
}

func (m *message) SendImgToChatroom(ctx context.Context, req *SendToChatroomMsgReq[ImgMsg]) (*SendMsgResp, error) {
	return m.SendToChatroom(ctx, req)
}

func (m *message) SendAudioToChatroom(ctx context.Context, req *SendToChatroomMsgReq[AudioMsg]) (*SendMsgResp, error) {
	return m.SendToChatroom(ctx, req)
}

func (m *message) SendVideoToChatroom(ctx context.Context, req *SendToChatroomMsgReq[VideoMsg]) (*SendMsgResp, error) {
	return m.SendToChatroom(ctx, req)
}

func (m *message) SendFileToChatroom(ctx context.Context, req *SendToChatroomMsgReq[FileMsg]) (*SendMsgResp, error) {
	return m.SendToChatroom(ctx, req)
}

func (m *message) SendLocToChatroom(ctx context.Context, req *SendToChatroomMsgReq[LocMsg]) (*SendMsgResp, error) {
	return m.SendToChatroom(ctx, req)
}

func (m *message) SendCmdToChatroom(ctx context.Context, req *SendToChatroomMsgReq[CmdMsg]) (*SendMsgResp, error) {
	return m.SendToChatroom(ctx, req)
}

func (m *message) SendCustomToChatroom(ctx context.Context, req *SendToChatroomMsgReq[CustomMsg]) (*SendMsgResp, error) {
	return m.SendToChatroom(ctx, req)
}

func (m *message) SendToChatroom(ctx context.Context, req any) (*SendMsgResp, error) {
	resp := SendMsgResp{}
	err := m.client.Post(ctx, "/messages/chatrooms", req, &resp)
	return &resp, err
}
