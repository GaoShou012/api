package cipher

type Cipher interface {
	// 设置密码本
	SetCipherKey(key []byte)
	Encrypt(model interface{}) (string, error)
	Decrypt(str string, model interface{}) error

	/*
		@desc
		加密

		@params
		cipherKey 密码本
		model 需要加密的数据模型

		@return
		返回加密后的密文
	*/
	EncryptWithCipherKey(model interface{}, cipherKey []byte) (string, error)
	/*
		@desc
		解密

		@params
		cipherKey 密码本
		token 密文
		model 解密到指定的数据模型
	*/
	DecryptWithCipherKey(str string, model interface{}, cipherKey []byte) error
}
