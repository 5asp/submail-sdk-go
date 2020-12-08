package client

import (
	"fmt"
	"github.com/t0of/submail-sdk-go/submail/signer"
	"net/http"
	"net/url"
	"time"
)

const BaseURLv1 = "https://api.mysubmail.com"

type Client struct {
	BaseURL    string
	apiKey     string
	appId      string
	HTTPClient *http.Client
	Data       interface{}
	param      string
}

/**
*	获取配置
**/

func NewClient(appId, apiKey string) *Client {
	return &Client{
		BaseURL: BaseURLv1,
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
	param := sortMap(data)
	paramStr := httpBuildQuery(param)
	var sign signer.Signer
	var signStr string
	switch data["sign_type"] {
	case "md5":
		sign = new(signer.Md5Signer)
		signStr = sign.Create(paramStr)
	case "sha1":
		sign = new(signer.Sha1Signer)
		signStr = sign.Create(paramStr)
	case "normal":
		signStr = param["signature"]
	default:
		fmt.Println("1")
		sign = new(signer.Sha256Signer)
		signStr = sign.Create(paramStr)
	}
	fmt.Println(signStr)
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
