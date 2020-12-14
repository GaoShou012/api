package admin_api

import (
	"testing"
)

type User struct {
	Username string
	Password string
}

func AFunc(v interface{}){
	u := User{
		Username: "aaa",
		Password: "bbb",
	}
	var row []interface{}
	row = append(row,u)
	v = row
}

func TestCasbinHandler(t *testing.T) {

}
