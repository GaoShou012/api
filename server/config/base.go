package config

type Base struct {
	GinMode                   string
	GinPort                   int
	OperatorContextCipherKey  string
	OperatorContextExpiration uint64
}
