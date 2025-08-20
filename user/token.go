package user

import (
	"context"
)

type TokenReq struct {
	Username       string `json:"username"`
	GrantType      string `json:"grant_type" default:"inherit"` //授权方式。设置为 inherit，表示通过用户 ID 获取用户 Token，需设置 username 参数。
	AutoCreateUser bool   `json:"autoCreateUser"`               //用不不存在时是否自动创建
	Ttl            int    `json:"ttl,omitempty"`                //用户 Token 有效期，单位为秒。设置为 0 表示 Token 有效期为永久。若不传该参数，有效期默认为 60 天。此外，也可通过 环信控制台的 用户管理 页面设置。该参数值以最新设置为准。
}

type TokenResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	User        Entity `json:"user"`
}

func (u *user) GetToken(ctx context.Context, req *TokenReq) (*TokenResp, error) {
	resp := &TokenResp{}
	if req.Username == "" {
		req.GrantType = "inherit"
	}
	err := u.client.Post(ctx, "/token", req, &resp)
	return resp, err
}
