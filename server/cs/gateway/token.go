package gateway

import "api/cs"

type Token struct {
	TenantCode string
	UserType   string
	UserId     uint64
}

func (t *Token) Encrypt() (string, error) {
	return cs.CipherOfToken.Encrypt(t)
}
func (t *Token) Decrypt(str string) error {
	return cs.CipherOfToken.Decrypt(str, t)
}
