package main

import (
	"github.com/t0of/submail-sdk-go/service"
)



func main() {
	const APP_ID = "47786"
	const SIGNATURE = "6dba41491c023a4960a4c0c132050255"
	s := service.NewSMS(APP_ID, SIGNATURE)
	s.AddTo("13027232773")
	s.AddProject("DrP9S3")
	s.AddVars("code","100")
	s.AddVars("msg","10000")
	//s.AddTag("1")
	s.Xsend()
	//fmt.Println(s)
}
