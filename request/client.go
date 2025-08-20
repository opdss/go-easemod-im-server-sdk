package request

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/sync/singleflight"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type Config struct {
	Endpoint     string `json:"endpoint" ` // 接口请求的地址, 这里注意后缀不可以带有 /
	OrgName      string // 环信企业名
	AppName      string // 环信应用名
	ClientId     string
	ClientSecret string
}

type CommonResp struct {
	Path            string `json:"path"`
	Uri             string `json:"uri"`
	Timestamp       int64  `json:"timestamp"` //毫秒
	Organization    string `json:"organization"`
	Application     string `json:"application"`
	Action          string `json:"action"`
	Duration        int    `json:"duration"` //执行时间，ms
	ApplicationName string `json:"applicationName"`
}
type ErrorResp struct {
	Code             int    `json:"-"`
	Error_           string `json:"error"`
	Exception        string `json:"exception"`
	Timestamp        int64  `json:"timestamp"`
	Duration         int    `json:"duration"`
	ErrorDescription string `json:"error_description"`
}

func (e *ErrorResp) Error() string {
	return fmt.Sprintf("%s[%d]: %s", e.Error_, e.Code, e.ErrorDescription)
}

type RequestErrorResp struct {
	Code int    `json:"code"`
	Body string `json:"body"`
	Err  error  `json:"err"`
}

func (e *RequestErrorResp) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return fmt.Sprintf("%d: %s", e.Code, e.Body)
}

type TokenResp struct {
	AccessToken string    `json:"access_token"` //有效的 Token 字符串。
	Application string    `json:"token_type"`   //当前 App 的 UUID 值。
	ExpiresIn   int64     `json:"expires_in"`   //Token 有效时间，单位为秒，在有效期内不需要重复获取。
	ExpiresAt   time.Time `json:"-"`            //Token 有效时间，单位为秒，在有效期内不需要重复获取。
}

type Client struct {
	config       Config
	token        *TokenResp
	singleFlight *singleflight.Group
}

func NewClient(conf Config) *Client {
	return &Client{
		config:       conf,
		singleFlight: &singleflight.Group{},
	}
}

func (c *Client) Config() Config {
	return c.config
}

func (c *Client) Get(ctx context.Context, apiPath string, resp any) error {
	return c.requestWithToken(ctx, http.MethodGet, apiPath, nil, resp)
}

func (c *Client) Post(ctx context.Context, apiPath string, req any, resp any) error {
	return c.requestWithToken(ctx, http.MethodPost, apiPath, req, resp)
}

func (c *Client) Put(ctx context.Context, apiPath string, req any, resp any) error {
	return c.requestWithToken(ctx, http.MethodPut, apiPath, req, resp)
}

func (c *Client) Delete(ctx context.Context, apiPath string, req any, resp any) error {
	return c.requestWithToken(ctx, http.MethodDelete, apiPath, req, resp)
}

func (c *Client) requestWithToken(ctx context.Context, method string, apiPath string, req any, resp any) error {
	tk, err := c.GetToken(ctx)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if req != nil {
		enc := json.NewEncoder(buf)
		err = enc.Encode(req)
		if err != nil {
			return err
		}
	}
	fmt.Printf(buf.String())
	httpReq, err := http.NewRequestWithContext(ctx, method, c.getApiUrl(apiPath), buf)
	if err != nil {
		return err
	}
	httpReq.Header.Set("Authorization", "Bearer "+tk)
	return c.request(httpReq, resp)
}

func (c *Client) request(req *http.Request, resp any) error {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	//req.Header.Set("Authorization", "Bearer "+tk)
	httpResp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(httpResp.Body)
	defer func() {
		_ = httpResp.Body.Close()
	}()
	log.Println(string(body))
	if err != nil {
		return &RequestErrorResp{Code: httpResp.StatusCode, Err: err, Body: string(body)}
	}
	if httpResp.StatusCode != 200 {
		errResp := &ErrorResp{}
		err = json.Unmarshal(body, errResp)
		if err != nil {
			return &RequestErrorResp{Code: httpResp.StatusCode, Err: err, Body: string(body)}
		}
		errResp.Code = httpResp.StatusCode
		return errResp
	}
	return json.Unmarshal(body, &resp)
}

func (c *Client) getApiUrl(path string) string {
	return fmt.Sprintf("%s/%s/%s%s", c.config.Endpoint, c.config.OrgName, c.config.AppName, path)
}

func (c *Client) GetToken(ctx context.Context) (string, error) {
	_resp, err, _ := c.singleFlight.Do("GetToken", func() (any, error) {
		if c.token != nil && c.token.AccessToken != "" && c.token.ExpiresAt.After(time.Now()) {
			return c.token.AccessToken, nil
		}
		resp, err := c.getToken(ctx)
		if err != nil {
			return "", err
		}
		c.token = resp
		c.token.ExpiresAt = time.Now().Add(time.Second*time.Duration(resp.ExpiresIn) - 3)
		return resp.AccessToken, nil
	})
	if err != nil {
		return "", err
	}
	return _resp.(string), nil
}

func (c *Client) getToken(ctx context.Context) (*TokenResp, error) {
	buf := strings.NewReader("{\"grant_type\":\"client_credentials\",\"client_id\":\"" + c.config.ClientId + "\",\"client_secret\":\"" + c.config.ClientSecret + "\",\"ttl\":86400}")
	req, err := http.NewRequestWithContext(ctx, "POST", c.getApiUrl("/token"), buf)
	if err != nil {
		return nil, err
	}
	resp := &TokenResp{}
	err = c.request(req, resp)
	return resp, err
}
