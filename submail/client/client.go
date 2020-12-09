package client

import (
	"fmt"
	"github.com/t0of/submail-sdk-go/submail/signer"
	"net/http"
	"net/url"
	"time"
)

//const BaseURLv1 = "https://api.mysubmail.com"

type Client struct {
	BaseURL    string
	apiKey     string
	appId      string
	HTTPClient *http.Client
	param      string
}

/**
*	获取配置
**/

func NewClient(appId, apiKey, APIUrl string) *Client {
	return &Client{
		BaseURL: APIUrl,
		apiKey:  apiKey,
		appId:   appId,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

/**
*	先加密
*	再拼接字符串
*	再发动请求
 */

func (c *Client) Do(data map[string]string) {
	param := make(map[string]string)
	data["appkey"] = c.apiKey
	param = sortMap(data)
	paramStr := httpBuildQuery(param)
	var sign signer.Signer
	switch data["sign_type"] {
	case "md5":
		sign = new(signer.Md5Signer)
		data["signature"] = sign.Create(paramStr)
	case "sha1":
		sign = new(signer.Sha1Signer)
		data["signature"] = sign.Create(paramStr)
	case "normal":
		data["signature"] = param["signature"]
	default:
		sign = new(signer.Sha256Signer)
		data["signature"] = sign.Create(paramStr)
	}
	data["appid"] = c.appId
	delete(data, "appkey")
	fmt.Println(c.sendRequest(c.BaseURL,httpBuildQuery(data)))
}

func httpBuildQuery(m map[string]string) string {
	var urlS url.URL
	q := urlS.Query()
	for k := range m {
		q.Add(k, m[k])
	}
	return q.Encode()
}

func sortMap(m map[string]string) map[string]string {
	var s []string
	for key := range m {
		s = append(s, key)
	}
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s)-1; j++ {
			if m[s[i]] < m[s[j]] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	return m
}
