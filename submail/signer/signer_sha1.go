package signer

import (
	"crypto/sha1"
	"encoding/hex"
)

type Sha1Signer struct {
}

func (str Sha1Signer) Create(string2 string) string {
	o := sha1.New()
	o.Write([]byte(string2))
	return hex.EncodeToString(o.Sum(nil))
}
