package gateway_v1

import "im/class/gateway"

type event struct {
	T gateway.MessageType
	H map[string]string
	B []byte
}

func (e *event) Type() gateway.MessageType {
	return e.T
}
func (e *event) Header() map[string]string {
	return e.H
}
func (e *event) Body() []byte {
	return e.B
}
