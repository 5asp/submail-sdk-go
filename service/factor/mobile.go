package factor

import (
	"fmt"
	"github.com/t0of/submail-sdk-go/submail/client"
)

/**
*	Service 服务层
 */
type factorErrResponse struct {
	Status string `json:"status"`
	Code   string `json:"code"`
	Msg    string `json:"msg"`
}
type factorSuccessResponse struct {
	Status string `json:"status"`
	SendId string `json:"send_id"`
	Fee    string `json:"fee"`
}

type Mobile struct {
	Name string
	Mobile	string
	Timestamp string

	Client      *client.Client
}

func (this *Mobile) AddName(name string) *Mobile {
	this.Name = name
	return this
}

func (this *Mobile) AddMobile(mobile string) *Mobile {
	this.Mobile = mobile
	return this
}

func (this *Mobile) AddTimestamp(timestamp string) *Mobile {
	this.Timestamp = timestamp
	return this
}


func NewFactorMobile(APP_ID, SIGNATURE string) *Mobile {
	return &Mobile{Client: client.NewClient(APP_ID, SIGNATURE)}
}

func (this *Mobile)Send()  {
	var data = make(map[string]string)

	if this.Mobile != "" {
		data["mobile"] = this.Mobile
	}
	if this.Name != "" {
		data["name"] = this.Name
	}
	this.Client.Do(data)
	fmt.Println("111")
}