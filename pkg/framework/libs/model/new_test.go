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
}
