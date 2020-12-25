package middleware_gin

import (
	"fmt"
	"framework/class/middleware"
	"testing"
)

var pluginTest middleware.OperatorContext

func TestPlugin_Init(t *testing.T) {
	pluginTest = New(
		WithModel(&operatorTest{}),
		WithCipherKey([]byte("123123")),
	)

	opInfo := &operatorTest{
		UserId:    1,
		Username:  "gaoshou",
		ContextId: "",
	}

	str, err := pluginTest.SignedString(opInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("string:", str)

	{
		opInfo := &operatorTest{}
		p := pluginTest.(*plugin)
		if err := p.decrypt(p.cipherKey,str,opInfo); err != nil {
			panic(err)
		}
		fmt.Println(opInfo)
	}
}
