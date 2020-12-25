package config

type Base struct {
	GinMode string
	GinPort int

	EnableAuthCode                  bool

	OperatorContextCipherKey  string
	OperatorContextExpiration uint64
}
