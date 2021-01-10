package cipher_jwt

import (
	"encoding/json"
	"fmt"
	"framework/class/cipher"
	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
)

var _ cipher.Cipher = &plugin{}

type plugin struct {
	opts      *Options
	cipherKey []byte
}

func (p *plugin) Init() error {
	return nil
}

func (p *plugin) SetCipherKey(key []byte) {
	p.cipherKey = key
}

func (p *plugin) Encrypt(model interface{}) (string, error) {
	return p.EncryptWithCipherKey(model, p.cipherKey)
}

func (p *plugin) Decrypt(str string, model interface{}) error {
	return p.DecryptWithCipherKey(str, model, p.cipherKey)
}

func (p *plugin) EncryptWithCipherKey(model interface{}, cipherKey []byte) (string, error) {
	j, err := json.Marshal(model)
	if err != nil {
		return "", err
	}
	m := jwt.MapClaims{}
	if err := json.Unmarshal(j, &m); err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, m)
	str, err := token.SignedString(cipherKey)
	if err != nil {
		return "", err
	}
	return str, nil
}

func (p *plugin) DecryptWithCipherKey(str string, model interface{}, cipherKey []byte) error {
	// parse the string to be token
	token, err := jwt.Parse(str, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v\n", token.Header["alg"])
		}
		return cipherKey, nil
	})
	if err != nil {
		return err
	}

	// convert the claims to model
	{
		decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			DecodeHook:       mapstructure.ComposeDecodeHookFunc(toTimeHookFunc()),
			ErrorUnused:      false,
			ZeroFields:       false,
			WeaklyTypedInput: false,
			Squash:           false,
			Metadata:         nil,
			Result:           model,
			TagName:          "",
		})
		if err != nil {
			return err
		}

		input := token.Claims.(jwt.Claims)

		if err := decoder.Decode(input); err != nil {
			return err
		}
	}

	return nil
}
