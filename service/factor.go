package service

import (
	"github.com/t0of/submail-sdk-go/submail/client"
	"strconv"
	"time"
)

type Factor struct {
	Name   string
	IdNo   string
	Mobile string
	SendURL	string
	Client *client.Client
}

const FactorGateway = "https://tpa.mysubmail.com/"

func (this *Factor) AddIdNo(idNo string) *Factor {
	this.IdNo = idNo
	return this
}
func (this *Factor) AddName(name string) *Factor {
	this.Name = name
	return this
}
func (this *Factor) AddMobile(mobile string) *Factor {
	this.Mobile = mobile
	return this
}

func NewFactorDetail3(APP_ID, SIGNATURE string) *Factor {
	return &Factor{Client: client.NewClient(APP_ID, SIGNATURE),SendURL:FactorGateway+"factor/detail3"}
}
func NewFactorSimple3(APP_ID, SIGNATURE string) *Factor {
	return &Factor{Client: client.NewClient(APP_ID, SIGNATURE),SendURL:FactorGateway+"factor/simple3"}
}
func NewFactorIdCard(APP_ID, SIGNATURE string) *Factor {
	return &Factor{Client: client.NewClient(APP_ID, SIGNATURE),SendURL:FactorGateway+"factor/idcard"}
}
func NewFactorMobile(APP_ID, SIGNATURE string) *Factor {
	return &Factor{Client: client.NewClient(APP_ID, SIGNATURE),SendURL:FactorGateway+"factor/mobile"}
}

func (this *Factor) Send() {
	var data = make(map[string]string)
	if this.IdNo != "" {
		data["idNo"] = this.IdNo
	}
	if this.Name != "" {
		data["name"] = this.Name
	}
	if this.Mobile != "" {
		data["mobile"] = this.Mobile
	}
	data["timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	this.Client.Do(this.SendURL,data)
}
