package lib_model

import "reflect"

type testModel struct {
	Username string
}

func NewModel(v interface{}) interface{} {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return reflect.New(t).Interface()
}
