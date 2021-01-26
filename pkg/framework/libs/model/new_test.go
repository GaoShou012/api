package lib_model

import (
	"fmt"
	"testing"
)

func TestNewModel(t *testing.T) {
	m1 := &testModel{Username: "333"}
	m2 := NewModel(m1)
	m1.Username = "123"
	//m2.(*testModel).Username = "456"

	fmt.Println(m1)
	fmt.Println(m2)

	//arr := make([]*testModel,0)
	//arrInstance := NewModel(arr)
	//fmt.Println(arrInstance)

	//var tt []testModel
	//tt = append(tt,testModel{Username: "123"})
	//arr := New(tt)
	//fmt.Println("arr:",arr)
	//arr1 := arr.(*[]testModel)
	//fmt.Println(arr1)
	m3 := NewModelSlice(m1)
	m4 := m3.([]*testModel)
	fmt.Println(m3)

	m4 = append(m4,&testModel{Username: "999"})

	fmt.Println("m4",m4)
	fmt.Println(m4[0].Username)
}
