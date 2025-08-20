package message

import (
	"context"
	"github.com/opdss/go-easemod-im-server-sdk/request"
)

type BroadcastMsgBody interface {
	BroadcastTxtMsg | BroadcastImgMsg | BroadcastAudioMsg | BroadcastVideoMsg | BroadcastFileMsg | BroadcastLocMsg | BroadcastCmdMsg | BroadcastCustomMsg
}

// TxtMsg 文本消息
type BroadcastTxtMsg struct {
	Type Type `json:"type"`
	TxtMsg
}

// ImgMsg 图片消息
type BroadcastImgMsg struct {
	Type Type `json:"type"`
	ImgMsg
}

// AudioMsg 语音消息
type BroadcastAudioMsg struct {
	Type Type `json:"type"`
	AudioMsg
}

// VideoMsg 视频消息
type BroadcastVideoMsg struct {
	Type Type `json:"type"`
	VideoMsg
}

// FileMsg 文件消息
type BroadcastFileMsg struct {
	Type Type `json:"type"`
	FileMsg
}

// LocMsg 位置消息
type BroadcastLocMsg struct {
	Type Type `json:"type"`
	LocMsg
}

// CmdMsg 透传消息
type BroadcastCmdMsg struct {
	Type   Type   `json:"type"`
	CmdMsg CmdMsg `json:"cmdMsg"`
}

// CustomMsg 自定义消息body
type BroadcastCustomMsg struct {
	Type Type `json:"type"`
	CustomMsg
}

type BroadcastMsgReq[T BroadcastMsgBody] struct {
	TargetType string            `json:"target_type"`    //广播消息接收方。固定值为 users，表示 app 下的所有用户。
	From       *string           `json:"from,omitempty"` //消息发送方的用户 ID。若不传入该字段，服务器默认设置为 admin。
	Appkey     string            `json:"appkey"`
	Msg        T                 `json:"msg"`
	Ext        map[string]string `json:"ext,omitempty"` //消息支持扩展字段，可添加自定义信息。不能对该参数传入 null。同时，推送通知也支持自定义扩展字段，详见 APNs 自定义显示 和 Android 推送字段说明。
}

type BroadcastMsgResp[T MsgBody] struct {
	request.CommonResp
	Data struct {
		Id int64 `json:"id"`
	} `json:"body"`
}

func (m *message) BroadcastTxt(ctx context.Context, req *BroadcastMsgReq[BroadcastTxtMsg]) (*SendMsgResp, error) {
	return m.Broadcast(ctx, req)
}

func (m *message) BroadcastImg(ctx context.Context, req *BroadcastMsgReq[BroadcastImgMsg]) (*SendMsgResp, error) {
	return m.Broadcast(ctx, req)
}

func (m *message) BroadcastAudio(ctx context.Context, req *BroadcastMsgReq[BroadcastAudioMsg]) (*SendMsgResp, error) {
	return m.Broadcast(ctx, req)
}

func (m *message) BroadcastVideo(ctx context.Context, req *BroadcastMsgReq[BroadcastVideoMsg]) (*SendMsgResp, error) {
	return m.Broadcast(ctx, req)
}

func (m *message) BroadcastFile(ctx context.Context, req *BroadcastMsgReq[BroadcastFileMsg]) (*SendMsgResp, error) {
	return m.Broadcast(ctx, req)
}

func (m *message) BroadcastLoc(ctx context.Context, req *BroadcastMsgReq[BroadcastLocMsg]) (*SendMsgResp, error) {
	return m.Broadcast(ctx, req)
}

func (m *message) BroadcastCmd(ctx context.Context, req *BroadcastMsgReq[BroadcastCmdMsg]) (*SendMsgResp, error) {
	return m.Broadcast(ctx, req)
}

func (m *message) BroadcastCustom(ctx context.Context, req *BroadcastMsgReq[BroadcastCustomMsg]) (*SendMsgResp, error) {
	return m.Broadcast(ctx, req)
}

func (m *message) Broadcast(ctx context.Context, req any) (*SendMsgResp, error) {
	resp := SendMsgResp{}
	err := m.client.Post(ctx, "/messages/broadcast", req, &resp)
	return &resp, err
}
