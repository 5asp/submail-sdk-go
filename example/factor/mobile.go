package main

import (
	"github.com/t0of/submail-sdk-go/service"
)



func main() {
	APP_ID := "10010"
	SIGNATURE := "a8b681cde917ca064d63b29381ec58f8"
	s := service.NewFactorMobile(APP_ID, SIGNATURE)
	s.AddName("夏学进")
	s.AddMobile("13027232773")
	s.Send()
	//fmt.Println(s)
}
