package admin_api

import (
	"api/global"
	libs_http "api/libs/http"
	"api/models"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
)

/*
	operator数据结构
*/
type Operator struct {
	// 租户ID
	TenantId uint64
	// 租户编码
	TenantCode string
	// 用户ID
	UserId uint64
	// 用户类型
	UserType uint64
	// 用户账号
	Username string
	// 用户昵称
	Nickname string
	// 登陆时间
	LoginTime time.Time
	// 角色ID
	Roles string

	// 上下文ID
	ContextId string
}

func (c *Operator) SetContextId(uuid string) {
	c.ContextId = uuid
}
func (c *Operator) GetContextId() string {
	return c.ContextId
}

func (c *Operator) GetTenantId() uint64 {
	return 0
}
func (c *Operator) GetId() uint64 {
	return c.UserId
}
func (c *Operator) GetUsername() string {
	return c.Username
}

func (c *Operator) GetAuthorityId() string {
	return c.Username
}

/*
	获取操作者信息
	@method GET
*/
func (c *Operator) Info(ctx *gin.Context) {
	libs_http.RspData(ctx, 0, "", GetOperator(ctx))
}

/*
	获取操作者菜单
	@method GET
*/
type MenuTree struct {
	Group *models.RbacMenuGroup
	Menus []*models.RbacMenu
}

func (c *Operator) Menu(ctx *gin.Context) {
	operator := GetOperator(ctx)

	roleIdMulti := make([]uint64, 0)
	roleIdMultiStr := strings.Split(operator.Roles, ",")
	for _, row := range roleIdMultiStr {
		id, err := strconv.Atoi(row)
		if err != nil {
			libs_http.RspState(ctx, 1, err)
			return
		}
		roleIdMulti = append(roleIdMulti, uint64(id))
	}

	groups := make([]*models.RbacMenuGroup, 0)
	err := global.RBAC.SelectMenuGroupWithFieldsByRoleIdMulti(operator, roleIdMulti, "`id`,`name`,`desc`", &groups)
	if err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}

	menuTree := make([]*MenuTree, 0)
	menuTreeIndex := make(map[uint64]int)

	{
		for _, group := range groups {
			m := &MenuTree{
				Group: group,
				Menus: make([]*models.RbacMenu, 0),
			}
			menuTree = append(menuTree, m)
			menuTreeIndex[*group.Id] = len(menuTree) - 1
		}
	}

	{
		menus := make([]*models.RbacMenu, 0)
		err := global.RBAC.SelectMenuWithFieldsByRoleIdMulti(operator, roleIdMulti, "`group_id`,`name`,`code`,`icon`,`desc`", &menus)
		if err != nil {
			libs_http.RspState(ctx, 1, err)
			return
		}
		for _, menu := range menus {
			groupId := *menu.GroupId
			menu.GroupId = nil
			index := menuTreeIndex[groupId]
			menuTree[index].Menus = append(menuTree[index].Menus, menu)
		}
	}

	libs_http.RspData(ctx, 0, "", menuTree)
}
