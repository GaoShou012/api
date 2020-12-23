package alert_event

type Callback struct {
	IsEnableOnCustomerDoesNotAsk
	IsEnableOnCustomerServerDoesNotAnswer
	OnCustomerDoesNotAsk
	OnCustomerServerDoesNotAnswer
}

/*
	是否开启，访客无提问，超时提醒
	@return
	true 允许开启倒计时
	false 不允许开启倒计时
*/
type IsEnableOnCustomerDoesNotAsk func(evt []byte, options map[string]string) bool

/*
	是否开启，客服无应答，自动应答
	@return
	true 允许开启倒计时
	false 不允许开启倒计时
*/
type IsEnableOnCustomerServerDoesNotAnswer func(evt []byte, options map[string]string) bool

/*
	当客户建立会话之后，长时间没有提问
	session 会话
	customer 访客
	counter 超时次数
*/
type OnCustomerDoesNotAsk func(evt []byte, options map[string]string, counter uint64)

/*
	当客服接待服务后，长时间没有应答
	session 会话
	customer 访客
	counter 超时次数
*/
type OnCustomerServerDoesNotAnswer func(evt []byte, options map[string]string, counter uint64)
