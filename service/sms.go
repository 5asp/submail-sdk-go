package service

import (
	"github.com/t0of/submail-sdk-go/submail/client"
)

/**
*	Service 服务层
 */
type smsErrResponse struct {
	Status string `json:"status"`
	Code   string `json:"code"`
	Msg    string `json:"msg"`
}
type smsSuccessResponse struct {
	Status string `json:"status"`
	SendId string `json:"send_id"`
	Fee    string `json:"fee"`
}
type SMS struct {
	To          string
	Content     string
	Tag         string
	SignVersion string
	SignType    string
	Vars        string
	Project     string
	Signature	string
	Multi       string
	Client      *client.Client
}

func (this *SMS) AddTo(to string) *SMS {
	this.To = to
	return this
}

func (this *SMS) AddContent(content string) *SMS {
	this.Content = content
	return this

}

func (this *SMS) AddTag(tag string) *SMS {
	this.Tag = tag
	return this

}

func (this *SMS) AddSignVersion(sign_version string) *SMS {
	this.SignVersion = sign_version
	return this

}

func (this *SMS) AddSignType(sign_type string) *SMS {
	this.SignType = sign_type
	return this

}

func (this *SMS) AddVars(vars string) *SMS {
	this.Vars = vars
	return this

}

func (this *SMS) AddProject(project string) *SMS {
	this.Project = project
	return this

}

func (this *SMS) AddMulti(multi string) *SMS {
	this.Multi = multi
	return this
}

func (this *SMS) AddSignature(signature string) *SMS {
	this.Signature = signature
	return this
}

func (this *SMS) Send() {
	var data = make(map[string]string)
	if this.To != "" {
		data["to"] = this.To
	}
	if this.Content != "" {
		data["content"] = this.Content
	}
	if this.SignType != "" {
		data["sign_type"] = this.SignType
	} else {
		data["sign_type"] = "normal"
	}
	if this.SignVersion != "" {
		data["sign_version"] = this.SignVersion
	}
	if this.Tag != "" {
		data["tag"] = this.Tag
	}
	this.Client.Do(data)
}
func (this *SMS) Xsend() {

}
func (this *SMS) Multisend() {

}
func (this *SMS) Multixsend() {

}
func NewSMS(APP_ID, SIGNATURE string) *SMS {
	return &SMS{Client: client.NewClient(APP_ID, SIGNATURE)}
}
