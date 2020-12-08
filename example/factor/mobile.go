package main

import (
	"github.com/t0of/submail-sdk-go/service/factor"
)

const APP_ID = "22452"
const SIGNATURE = "xxxxxxxxx"

func main() {
	s := factor.NewFactorMobile(APP_ID,SIGNATURE)
	s.AddName("content")
	s.AddTimestamp("1xxxxx")
	s.Send()
	//fmt.Println(s)
}