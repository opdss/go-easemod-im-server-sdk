package message

import (
	"context"
	"github.com/opdss/go-easemod-im-server-sdk/request"
)

type Message interface {
	// 单聊消息发送
	SendTxtToUser(ctx context.Context, req *SendToUserMsgReq[TxtMsg]) (*SendMsgResp, error)
	SendImgToUser(ctx context.Context, req *SendToUserMsgReq[ImgMsg]) (*SendMsgResp, error)
	SendAudioToUser(ctx context.Context, req *SendToUserMsgReq[AudioMsg]) (*SendMsgResp, error)
	SendVideoToUser(ctx context.Context, req *SendToUserMsgReq[VideoMsg]) (*SendMsgResp, error)
	SendFileToUser(ctx context.Context, req *SendToUserMsgReq[FileMsg]) (*SendMsgResp, error)
	SendLocToUser(ctx context.Context, req *SendToUserMsgReq[LocMsg]) (*SendMsgResp, error)
	SendCmdToUser(ctx context.Context, req *SendToUserMsgReq[CmdMsg]) (*SendMsgResp, error)
	SendCustomToUser(ctx context.Context, req *SendToUserMsgReq[CustomMsg]) (*SendMsgResp, error)
	SendToUser(ctx context.Context, req any) (*SendMsgResp, error)

	// 聊天室消息发送
	SendTxtToChatroom(ctx context.Context, req *SendToChatroomMsgReq[TxtMsg]) (*SendMsgResp, error)
	SendImgToChatroom(ctx context.Context, req *SendToChatroomMsgReq[ImgMsg]) (*SendMsgResp, error)
	SendAudioToChatroom(ctx context.Context, req *SendToChatroomMsgReq[AudioMsg]) (*SendMsgResp, error)
	SendVideoToChatroom(ctx context.Context, req *SendToChatroomMsgReq[VideoMsg]) (*SendMsgResp, error)
	SendFileToChatroom(ctx context.Context, req *SendToChatroomMsgReq[FileMsg]) (*SendMsgResp, error)
	SendLocToChatroom(ctx context.Context, req *SendToChatroomMsgReq[LocMsg]) (*SendMsgResp, error)
	SendCmdToChatroom(ctx context.Context, req *SendToChatroomMsgReq[CmdMsg]) (*SendMsgResp, error)
	SendCustomToChatroom(ctx context.Context, req *SendToChatroomMsgReq[CustomMsg]) (*SendMsgResp, error)
	SendToChatroom(ctx context.Context, req any) (*SendMsgResp, error)

	// 广播消息发送，慎用
	BroadcastTxt(ctx context.Context, req *BroadcastMsgReq[TxtMsg]) (*SendMsgResp, error)
	BroadcastImg(ctx context.Context, req *BroadcastMsgReq[ImgMsg]) (*SendMsgResp, error)
	BroadcastAudio(ctx context.Context, req *BroadcastMsgReq[AudioMsg]) (*SendMsgResp, error)
	BroadcastVideo(ctx context.Context, req *BroadcastMsgReq[VideoMsg]) (*SendMsgResp, error)
	BroadcastFile(ctx context.Context, req *BroadcastMsgReq[FileMsg]) (*SendMsgResp, error)
	BroadcastLoc(ctx context.Context, req *BroadcastMsgReq[LocMsg]) (*SendMsgResp, error)
	BroadcastCmd(ctx context.Context, req *BroadcastMsgReq[CmdMsg]) (*SendMsgResp, error)
	BroadcastCustom(ctx context.Context, req *BroadcastMsgReq[CustomMsg]) (*SendMsgResp, error)
	Broadcast(ctx context.Context, req any) (*SendMsgResp, error)
}

type message struct {
	client *request.Client
}

func NewMessage(client *request.Client) Message {
	return &message{
		client: client,
	}
}

type MsgLevel string

const (
	High   MsgLevel = "high"   //高
	Low    MsgLevel = "low"    //低
	Normal MsgLevel = "normal" //普通
)

// Type 消息类型
type Type string

const (
	Txt    Type = "txt"    //文本消息
	Img    Type = "img"    //图片消息
	Audio  Type = "audio"  //  语音消息
	Video  Type = "video"  //视频消息
	File   Type = "file"   //文件消息
	Loc    Type = "loc"    //位置消息
	Cmd    Type = "cmd"    //透传消息
	Custom Type = "custom" //自定义消息
)

// TxtMsg 文本消息
type TxtMsg struct {
	Msg string `json:"msg"`
}

type ImgMsgSize struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

// ImgMsg 图片消息
type ImgMsg struct {
	Filename *string     `json:"filename,omitempty"`
	Secret   *string     `json:"secret,omitempty"`
	Url      string      `json:"url"`
	Size     *ImgMsgSize `json:"size,omitempty"`
}

// AudioMsg 语音消息
type AudioMsg struct {
	Url      string  `json:"url"`
	Filename *string `json:"filename,omitempty"`
	Length   *int    `json:"length,omitempty"`
	Secret   *string `json:"secret,omitempty"`
}

// VideoMsg 视频消息
type VideoMsg struct {
	Filename    *string `json:"filename,omitempty"`
	Thumb       *string `json:"thumb,omitempty"`
	Length      *int    `json:"length,omitempty"`
	Secret      *string `json:"secret,omitempty"`
	FileLength  *int64  `json:"file_length,omitempty"`
	ThumbSecret *string `json:"thumb_secret,omitempty"`
	Url         string  `json:"url"`
}

// FileMsg 文件消息
type FileMsg struct {
	Filename *string `json:"filename,omitempty"`
	Secret   *string `json:"secret,omitempty"`
	Url      string  `json:"url"`
}

// LocMsg 位置消息
type LocMsg struct {
	Lat  string `json:"lat"`
	Lng  string `json:"lng"`
	Addr string `json:"addr"`
}

// CmdMsg 透传消息
type CmdMsg struct {
	Action string `json:"action"`
}

// CustomMsg 自定义消息body
type CustomMsg struct {
	CustomEvent *string    `json:"customEvent,omitempty"`
	CustomExts  CustomExts `json:"customExts,omitempty"`
}

type CustomExts map[string]string

type MsgBody interface {
	TxtMsg | ImgMsg | AudioMsg | VideoMsg | FileMsg | LocMsg | CmdMsg | CustomMsg
}

type MsgData map[string]string

type SendMsgResp struct {
	request.CommonResp
	Data MsgData `json:"data"`
}
