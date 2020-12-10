package main

import (
	"github.com/t0of/submail-sdk-go/service"
)

type Multis []*Multi

type Multi struct {
	To   string
	Vars map[string]string
}

func main() {
	const APP_ID = "47786"
	const SIGNATURE = "6dba41491c023a4960a4c0c132050255"
	s := service.NewSMS(APP_ID, SIGNATURE)
	s.AddTo("13027232773")
	s.AddProject("DrP9S3")
	var ms Multis
	var  Multi
	ms = {
		Multi{
			To:"1e323131"，
			Vars["ac"]:"sss",
			Vars["sss"]:"sss",
		},
		Multi{
			To:"1e323131"，
			Vars["ac"]:"sss",
			Vars["sss"]:"sss",
		}
	}
	//as := Multi{{To: "13027232773",Mul}}
	var multi =make(map[string]string)
	multi["code"]	=	"1000"
	multi["name"]	=	"2223"
	s.AddMulti("13027232773",multi)
	s.AddMulti("13175217275",multi)
	//s.AddVars("msg","10000")
	//s.AddTag("1")
	s.Multisend()
	//fmt.Println(s)
}
