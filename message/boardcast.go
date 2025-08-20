package message

import (
	"context"
	"github.com/opdss/go-easemod-im-server-sdk/request"
)
 
type BroadcastMsg[T MsgBody] struct {
	Type Type `json:"type"`
	Msg  T    `json:"msg"`
}

type BroadcastMsgReq[T MsgBody] struct {
	TargetType string            `json:"target_type"`    //广播消息接收方。固定值为 users，表示 app 下的所有用户。
	From       *string           `json:"from,omitempty"` //消息发送方的用户 ID。若不传入该字段，服务器默认设置为 admin。
	Appkey     string            `json:"appkey"`
	Msg        BroadcastMsg[T]   `json:"msg"`
	Ext        map[string]string `json:"ext,omitempty"` //消息支持扩展字段，可添加自定义信息。不能对该参数传入 null。同时，推送通知也支持自定义扩展字段，详见 APNs 自定义显示 和 Android 推送字段说明。
}

type BroadcastMsgResp[T MsgBody] struct {
	request.CommonResp
	Data struct {
		Id int64 `json:"id"`
	} `json:"body"`
}

func (m *message) BroadcastTxt(ctx context.Context, req *BroadcastMsgReq[TxtMsg]) (*SendMsgResp, error) {
	return m.Broadcast(ctx, req)
}

func (m *message) BroadcastImg(ctx context.Context, req *BroadcastMsgReq[ImgMsg]) (*SendMsgResp, error) {
	return m.Broadcast(ctx, req)
}

func (m *message) BroadcastAudio(ctx context.Context, req *BroadcastMsgReq[AudioMsg]) (*SendMsgResp, error) {
	return m.Broadcast(ctx, req)
}

func (m *message) BroadcastVideo(ctx context.Context, req *BroadcastMsgReq[VideoMsg]) (*SendMsgResp, error) {
	return m.Broadcast(ctx, req)
}

func (m *message) BroadcastFile(ctx context.Context, req *BroadcastMsgReq[FileMsg]) (*SendMsgResp, error) {
	return m.Broadcast(ctx, req)
}

func (m *message) BroadcastLoc(ctx context.Context, req *BroadcastMsgReq[LocMsg]) (*SendMsgResp, error) {
	return m.Broadcast(ctx, req)
}

func (m *message) BroadcastCmd(ctx context.Context, req *BroadcastMsgReq[CmdMsg]) (*SendMsgResp, error) {
	return m.Broadcast(ctx, req)
}

func (m *message) BroadcastCustom(ctx context.Context, req *BroadcastMsgReq[CustomMsg]) (*SendMsgResp, error) {
	return m.Broadcast(ctx, req)
}

func (m *message) Broadcast(ctx context.Context, req any) (*SendMsgResp, error) {
	resp := SendMsgResp{}
	err := m.client.Post(ctx, "/messages/chatrooms", req, &resp)
	return &resp, err
}
