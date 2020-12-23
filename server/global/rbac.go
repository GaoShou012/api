package global

import "framework/class/rbac"

var RBAC struct{
	rbac.ApiAdapter
	rbac.MenuAdapter
}
