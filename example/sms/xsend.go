package main

import (
	"github.com/t0of/submail-sdk-go/service"
)



func main() {
	const APP_ID = "47786"
	const SIGNATURE = "6dba41491c023a4960a4c0c132050255"
	s := service.NewSMS(APP_ID, SIGNATURE)
	s.AddTo("13027232773")
	s.AddContent("【大金门】验证码@var(code),当天使用有效。")
	s.
	//s.AddTag("1")
	s.Xsend()
	//fmt.Println(s)
}
