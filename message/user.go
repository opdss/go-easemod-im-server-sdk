package message

import (
	"context"
)

type Ext struct {
	EmIgnoreNotification *bool `json:"em_ignore_notification,omitempty"` //是否发送静默消息： - true：是； - （默认）false：否。
}

type UserMsgReq struct {
	From            *string  `json:"from,omitempty"` //消息发送方的用户 ID。若不传入该字段，服务器默认设置为 admin。
	To              []string `json:"to"`
	Type            Type     `json:"type"`
	SyncDevice      *bool    `json:"sync_device,omitempty"`       //消息发送成功后，是否将消息同步到发送方的所有在线设备。- true：是； - （默认）false：否。
	RoamIgnoreUsers []string `json:"roam_ignore_users,omitempty"` // 设置哪些用户拉漫游消息时拉不到该消息。
	RouteType       *string  `json:"route_type,omitempty"`        // 若传入该参数，其值为 ROUTE_ONLINE，表示接收方只有在线时才能收到消息，若接收方离线则无法收到消息。若不传入该参数，无论接收方在线还是离线都能收到消息。
	Ext             *Ext     `json:"ext,omitempty"`               //消息支持扩展字段，可添加自定义信息。不能对该参数传入 null。同时，推送通知也支持自定义扩展字段，详见 APNs 自定义显示 和 Android 推送字段说明。
}

type SendToUserMsgReq[T MsgBody] struct {
	UserMsgReq
	Body T `json:"body"`
}

func (m *message) SendTxtToUser(ctx context.Context, req *SendToUserMsgReq[TxtMsg]) (*SendMsgResp, error) {
	return m.SendToUser(ctx, req)
}

func (m *message) SendImgToUser(ctx context.Context, req *SendToUserMsgReq[ImgMsg]) (*SendMsgResp, error) {
	return m.SendToUser(ctx, req)
}

func (m *message) SendAudioToUser(ctx context.Context, req *SendToUserMsgReq[AudioMsg]) (*SendMsgResp, error) {
	return m.SendToUser(ctx, req)
}

func (m *message) SendVideoToUser(ctx context.Context, req *SendToUserMsgReq[VideoMsg]) (*SendMsgResp, error) {
	return m.SendToUser(ctx, req)
}

func (m *message) SendFileToUser(ctx context.Context, req *SendToUserMsgReq[FileMsg]) (*SendMsgResp, error) {
	return m.SendToUser(ctx, req)
}

func (m *message) SendLocToUser(ctx context.Context, req *SendToUserMsgReq[LocMsg]) (*SendMsgResp, error) {
	return m.SendToUser(ctx, req)
}

func (m *message) SendCmdToUser(ctx context.Context, req *SendToUserMsgReq[CmdMsg]) (*SendMsgResp, error) {
	return m.SendToUser(ctx, req)
}

func (m *message) SendCustomToUser(ctx context.Context, req *SendToUserMsgReq[CustomMsg]) (*SendMsgResp, error) {
	return m.SendToUser(ctx, req)
}

func (m *message) SendToUser(ctx context.Context, req any) (*SendMsgResp, error) {
	resp := SendMsgResp{}
	err := m.client.Post(ctx, "/messages/users", req, &resp)
	return &resp, err
}
