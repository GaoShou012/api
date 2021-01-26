package yaml

import (
	"fmt"
	"testing"
)

func TestGetRbacApiData(t *testing.T) {
	data, err := GetRbacApiData()
	if err != nil {
		panic(err)
	}
	fmt.Println("api data",data)
}
