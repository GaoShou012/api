package yaml

import (
	"fmt"
	"testing"
)

func TestGetRbacMenuData(t *testing.T) {
	data, err := GetRbacMenuData()
	if err != nil {
		panic(err)
	}
	fmt.Println("menu data",data)
}
