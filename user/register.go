package user

import (
	"context"
	"errors"
	"github.com/opdss/go-easemod-im-server-sdk/request"
)

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Entity struct {
	Username  string `json:"username"`
	UUID      string `json:"uuid"`
	Created   int64  `json:"created"`   // 册用户的 Unix 时间戳，单位为毫秒
	Modified  int64  `json:"modified"`  // 最近一次修改用户信息的 Unix 时间戳，单位为毫秒
	Activated bool   `json:"activated"` // 		用户是否为活跃状态：		// - true：用户为活跃状态。		// - false：用户为封禁状态。如要使用已被封禁
	Type      string `json:"type"`
}

type RegisterData struct {
	Username               string `json:"username"`
	RegisterUserFailReason string `json:"registerUserFailReason"`
}

type RegistryResp struct {
	request.CommonResp
	Entities []Entity       `json:"entities"`
	Data     []RegisterData `json:"data"`
}

// https://doc.easemob.com/document/server-side/account_system.html#%E6%8E%88%E6%9D%83%E6%B3%A8%E5%86%8C%E5%8D%95%E4%B8%AA%E7%94%A8%E6%88%B7
func (u *user) Register(ctx context.Context, req *RegisterReq) (*RegistryResp, error) {
	resp := &RegistryResp{}
	err := u.client.Post(ctx, "/users", req, &resp)
	return resp, err
}

// https://doc.easemob.com/document/server-side/account_system.html#%E6%89%B9%E9%87%8F%E6%8E%88%E6%9D%83%E6%B3%A8%E5%86%8C%E7%94%A8%E6%88%B7
func (u *user) BatchRegistry(ctx context.Context, users []*RegisterReq) (*RegistryResp, error) {
	// 用户信息必传
	if len(users) == 0 {
		return nil, errors.New("users is empty")
	}
	resp := &RegistryResp{}
	err := u.client.Post(ctx, "/users", users, &resp)
	return resp, err
}
