package casbin

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"log"
)

func TestCasbin() {
	e, _ := casbin.NewEnforcer("model.conf", "rbac.csv")

	fmt.Printf("RBAC test start\n") // output for debug

	// superAdmin
	res, _ := e.Enforce("quyuan", "project", "read")
	if res {
		log.Println("superAdmin can read project")
	} else {
		log.Fatal("ERROR: superAdmin can not read project")
	}
}
