package signer

import (
	"crypto/md5"
	"encoding/hex"
)

type Md5Signer struct {
}

func (str Md5Signer) Create(s string) string {
	md5 := md5.New()
	md5.Write([]byte(s))
	md5Data := md5.Sum([]byte(""))
	return hex.EncodeToString(md5Data)
}
