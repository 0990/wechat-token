package main

import (
	"encoding/json"

	"github.com/levigross/grequests"
)

const AccessTokenAPI = "https://api.weixin.qq.com/cgi-bin/token"

type Token struct {
	AccessToken string `json:"access_token"`
	Expire      int    `json:"expires_in"`
}

// 获取AppID的access_token
func (t *Token) Get(appid string, secret string) string {
	var params = map[string]string{
		"appid":      appid,
		"secret":     secret,
		"grant_type": "client_credential",
	}

	ro := &grequests.RequestOptions{
		Params: params,
	}

	res, _ := grequests.Get(AccessTokenAPI, ro)

	if err := json.Unmarshal(res.Bytes(), t); err != nil {
		return ""
	}

	return res.String()
}
