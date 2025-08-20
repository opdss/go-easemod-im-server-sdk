package user

import (
	"context"
	"github.com/opdss/go-easemod-im-server-sdk/request"
)

// https://doc.easemob.com/document/server-side/account_system.html#批量授权注册用户
type User interface {
	// Register 注册一个用户 https://doc.easemob.com/document/server-side/account_system.html#%E6%8E%88%E6%9D%83%E6%B3%A8%E5%86%8C%E5%8D%95%E4%B8%AA%E7%94%A8%E6%88%B7
	Register(ctx context.Context, req *RegisterReq) (*RegistryResp, error)
	// BatchRegistry 批量注册用户 https://doc.easemob.com/document/server-side/account_system.html#%E6%89%B9%E9%87%8F%E6%8E%88%E6%9D%83%E6%B3%A8%E5%86%8C%E7%94%A8%E6%88%B7
	BatchRegistry(ctx context.Context, req []*RegisterReq) (*RegistryResp, error) // 批量注册用户，单次请求最多可注册 60 个用户 ID。
	// GetToken 获取用户token https://doc.easemob.com/document/server-side/easemob_user_token.html#%E9%80%9A%E8%BF%87%E7%94%A8%E6%88%B7-id-%E8%8E%B7%E5%8F%96%E7%94%A8%E6%88%B7-token
	GetToken(ctx context.Context, req *TokenReq) (*TokenResp, error)
}

type user struct {
	client *request.Client
}

func NewUser(c *request.Client) User {
	return &user{
		client: c,
	}
}
