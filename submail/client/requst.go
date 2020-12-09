package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func (c *Client) sendRequest(url,s string) string {
	fmt.Println(url)
	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader(s))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func (c *Client) xsendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Accept", "application/json; charset=utf-8")
	return nil
}

func (c *Client) multiSendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Accept", "application/json; charset=utf-8")
	return nil
}

func (c *Client) multiXsendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Accept", "application/json; charset=utf-8")
	return nil
}
