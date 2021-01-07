package event

import "encoding/json"

type Event interface {
	Type() string
}

func Encode(event Event) ([]byte, error) {
	m := make(map[string]interface{})
	m["Type"] = event.Type()
	m["Data"] = event
	return json.Marshal(m)
}
