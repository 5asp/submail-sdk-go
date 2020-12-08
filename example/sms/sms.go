package main

import (
	"github.com/t0of/submail-sdk-go/service"
)

const APP_ID = "22452"
const SIGNATURE = "xxxxxxxxx"

func main() {
	s := service.NewSMS(APP_ID,SIGNATURE)
	s.AddTo("13027232773")
	s.AddContent("content")
	s.AddTag("1")
	s.Send()
	//fmt.Println(s)
}