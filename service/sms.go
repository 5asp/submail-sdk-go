package service

import (
	"encoding/json"
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
	Project     string
	Signature   string
	Multis      []Multi
	Vars        []map[string]string
	Client      *client.Client
}
type Multi struct {
	To   string
	Vars map[string]string
}

const APIGateway = "https://api.mysubmail.com/"

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

func (this *SMS) AddVars(varsKey, varsValue string) *SMS {
	vars := make(map[string]string)
	vars[varsKey] = varsValue
	this.Vars = append(this.Vars, vars)
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

func (this *SMS) AddProject(project string) *SMS {
	this.Project = project
	return this
}

func (this *SMS) AddMulti(to string, data map[string]string) *SMS {
	this.Multis = append(this.Multis, Multi{To: to,Vars: data})
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
	this.Client.Do(APIGateway+"message/send", data)
}

func (this *SMS) Xsend() {
	var data = make(map[string]string)
	if this.To != "" {
		data["to"] = this.To
	}
	if this.Project != "" {
		data["project"] = this.Project
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
	if this.Vars != nil {
		vars := make(map[string]string)
		for _, p := range this.Vars {
			for k, v := range p {
				vars[k] = v
			}
		}
		varStr, _ := json.Marshal(vars)
		data["vars"] = string(varStr)
	}
	this.Client.Do(APIGateway+"message/xsend", data)
}

func (this *SMS) Multisend() {
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
	if this.Multis != nil {
		//vars := make(map[string]string)
		//for _,p := range this.Vars{
		//	for k ,v :=range p {
		//		vars[k] = v
		//	}
		//}
		//varStr, _ := json.Marshal(vars)
		//data["vars"] = string(varStr)
		//fmt.Println(this.Multis)
	}
	this.Client.Do(APIGateway+"message/multisend", data)
}
func (this *SMS) Multixsend() {

}
func NewSMS(APP_ID, SIGNATURE string) *SMS {
	return &SMS{Client: client.NewClient(APP_ID, SIGNATURE)}
}
