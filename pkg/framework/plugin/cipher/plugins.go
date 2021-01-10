package cipher

import (
	"framework/class/cipher"
)

var _ cipher.Cipher = &plugin{}

type plugin struct {
	opts *Options
}

func (p *plugin) SetCipherKey(key []byte) {
	panic("implement me")
}

func (p *plugin) Encrypt(model interface{}) (string, error) {
	panic("implement me")
}

func (p *plugin) Decrypt(token string, model interface{}) error {
	panic("implement me")
}

func (p *plugin) EncryptWithCipherKey(model interface{}, cipherKey []byte) (string, error) {
	panic("implement me")
}

func (p *plugin) DecryptWithCipherKey(token string, model interface{}, cipherKey []byte) error {
	panic("implement me")
}

func (p *plugin) Init() interface{} {
	return nil
}
