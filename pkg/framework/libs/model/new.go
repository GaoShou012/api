package lib_model

import "reflect"

type testModel struct {
	Username string
}

func New(v interface{}) interface{} {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return reflect.New(t).Interface()
}

func NewModel(v interface{}) interface{} {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return reflect.New(t).Interface()
}

func NewModelSlice(v interface{}) interface{} {
	//t := reflect.TypeOf(v)
	//if t.Kind() == reflect.Ptr {
	//	t = t.Elem()
	//}
	//
	//sliceT := reflect.SliceOf(t)
	//
	//return reflect.MakeSlice(sliceT,0,0)

	typeOfT := reflect.TypeOf(v)
	sliceOfT := reflect.SliceOf(typeOfT)

	s := reflect.MakeSlice(sliceOfT, 0, 0).Interface()
	//In order to pass a pointer to the slice without knowing the type, you can create the pointer first, then set the slice value:

	//ptr := reflect.New(sliceOfT)
	//ptr.Elem().Set(reflect.MakeSlice(sliceOfT, 0, 0))
	//s := ptr.Interface()
	return s
}