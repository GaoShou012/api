package cipher_jwt

import (
	"framework/class/cipher"
)

type Option func(o *Options)
type Options struct{}

func New(opts ...Option) cipher.Cipher {
	options := &Options{}

	for _, o := range opts {
		o(options)
	}

	p := &plugin{
		opts: options,
	}
	if err := p.Init(); err != nil {
		panic(err)
	}

	return p
}
