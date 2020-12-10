package client

import (
	"fmt"
	"github.com/t0of/submail-sdk-go/submail/signer"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

//const BaseURLv1 = "https://api.mysubmail.com"

type Client struct {
	apiKey     string
	appId      string
	HTTPClient *http.Client
	param      string
}

/**
*	获取配置
**/
func NewClient(appId, apiKey string) *Client {
	return &Client{
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
func (c *Client) Do(Gateway string,data map[string]string) {
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
		data["signature"] = c.apiKey
	default:
		sign = new(signer.Sha256Signer)
		data["signature"] = sign.Create(paramStr)
	}
	data["appid"] = c.appId
	delete(data, "appkey")
	fmt.Println(data)
	fmt.Println(c.sendRequest(Gateway, httpBuildQuery(data)))
}

/**
*	发送请求
 */
func (c *Client) sendRequest(url, s string) string {
	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader(s))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
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
