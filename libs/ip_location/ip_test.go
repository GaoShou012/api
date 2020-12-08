package ip_location

import (
	"fmt"
	"testing"
)

func TestGetLocation(t *testing.T) {
	Init("../../config/ipipfree.ipdb")
	location, err := GetLocation("113.93.107.116")
	if err != nil {
		panic(err)
	}
	fmt.Println(location)
}
