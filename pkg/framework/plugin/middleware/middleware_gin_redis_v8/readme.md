## Gin框架操作者信息中间件

````

Get 使用实例

opContext := &middleware.OperatorContext{}
op,err := opContext.Get(ctx *gin.Context)

````


````

使用超时验证，operator必须要支持 SetContextId,GetContextId
保存和读取上下文ID

````