package signer

import (
	"crypto/sha256"
	"encoding/hex"
)

type Sha256Signer struct {
}

func (str Sha256Signer) Create(s string) string {
	hash := sha256.New()
	hash.Write([]byte(s))
	md := hash.Sum(nil)
	return hex.EncodeToString(md)
}
