package yaml

import (
	"fmt"
	"testing"
)

func TestGetRbacRoleData(t *testing.T) {
	data, err := GetRbacRoleData()
	if err != nil {
		panic(err)
	}
	fmt.Println("role data",data)
}