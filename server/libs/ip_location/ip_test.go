package libs_ip_location

import (
	"api/config"
	"fmt"
	"testing"
)

func TestGetLocation(t *testing.T) {
	config.LocalLoadGaoShou()
	Init(config.GetConfig().IpLocation.Path)
	location, err := GetLocation("113.93.107.116")
	if err != nil {
		panic(err)
	}
	fmt.Println(location)
}
