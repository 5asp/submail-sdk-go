package signer

type Signer interface {
	Create(string) string
}
