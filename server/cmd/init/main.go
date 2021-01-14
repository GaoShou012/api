package main

import (
	"api/config"
	"api/global"
	"api/initialize"
	libs_log "api/libs/logs"
	"api/models"
	"fmt"
	"os"
)

func init() {
	conf := config.GetConfig()
	conf.Load("./../../config/resource/dev.ini")

	// 初始化 日志
	{
		initialize.InitLogger()
	}

	// 初始化 mysql master
	{
		conf := config.GetConfig().MysqlMaster
		fmt.Println("db master :", conf)
		if err := initialize.InitMysqlMaster(conf); err != nil {
			libs_log.Error(err)
			os.Exit(0)
		}
	}

	// 初始化 mysql slave
	{
		conf := config.GetConfig().MysqlSlave
		if err := initialize.InitMysqlSlave(conf); err != nil {
			libs_log.Error(err)
			os.Exit(0)
		}
	}

	// 初始化 rbac
	{
		initialize.InitRBAC()
	}
}

func main() {
	createSuperManager()
	createMenus()
	createRole()
	createApi()
}

func createSuperManager() {
	model := &models.Admins{}
	exists, err := model.IsExistsByUsername("admin")
	if err != nil {
		panic(err)
	}
	if exists {
		fmt.Println("超级管理员已经存在，创建失败")
		return
	}

	enable := true
	state := uint64(1)
	userType := uint64(models.AdminsUserTypeSuperManager)
	username := "admin"
	nickname := "超级管理员"
	password := "123456"
	passwordEncrypted, err := model.EncryptPassword(password)
	if err != nil {
		panic(err)
	}
	admin := &models.Admins{
		Id:        nil,
		Enable:    &enable,
		State:     &state,
		UserType:  &userType,
		Username:  &username,
		Password:  &passwordEncrypted,
		Nickname:  &nickname,
		CreatedAt: nil,
		UpdatedAt: nil,
	}
	res := global.DBMaster.Table(model.GetTableName()).Create(admin)
	if res.Error != nil {
		panic(res.Error)
	}
	fmt.Println("创建超级管理员成功")
}

func createMenus() {
	{
		sort := uint64(999)
		name := "系统设置"
		icon := ""
		code := "system"
		desc := "菜单，角色，权限，等等..."

		group := &models.RbacMenuGroup{
			Model: models.Model{},
			Sort:  &sort,
			Name:  &name,
			Code:  &code,
			Icon:  &icon,
			Desc:  &desc,
		}
		if err := group.Insert(); err != nil {
			global.Logger.Error(err)
		}
	}

	{
		var groupId uint64
		{
			model := &models.RbacMenuGroup{}
			if err := model.SelectByCode("*", "system"); err != nil {
				panic(err)
			}
			groupId = *model.Id
		}
		{
			sort := uint64(0)
			name := "账号管理"
			code := "account"
			icon := ""
			desc := ""
			menu := &models.RbacMenu{
				Model:   models.Model{},
				GroupId: &groupId,
				Sort:    &sort,
				Name:    &name,
				Code:    &code,
				Icon:    &icon,
				Desc:    &desc,
			}
			if err := menu.Insert(); err != nil {
				global.Logger.Error(err)
			}
		}
		{
			sort := uint64(1)
			name := "角色管理"
			code := "role"
			icon := ""
			desc := ""
			menu := &models.RbacMenu{
				Model:   models.Model{},
				GroupId: &groupId,
				Sort:    &sort,
				Name:    &name,
				Code:    &code,
				Icon:    &icon,
				Desc:    &desc,
			}
			if err := menu.Insert(); err != nil {
				global.Logger.Error(err)
			}
		}
		{
			sort := uint64(2)
			name := "接口管理"
			code := "interface"
			icon := ""
			desc := ""
			menu := &models.RbacMenu{
				Model:   models.Model{},
				GroupId: &groupId,
				Sort:    &sort,
				Name:    &name,
				Code:    &code,
				Icon:    &icon,
				Desc:    &desc,
			}
			if err := menu.Insert(); err != nil {
				global.Logger.Error(err)
			}
		}
	}
}

func createRole() {
	roleId := uint64(1)

	{
		name := "超级管理员"
		desc := "系统默认创建的超级管理员角色,角色ID必须为1"
		icon := ""
		role := &models.RbacRole{
			Model: models.Model{},
			Name:  &name,
			Desc:  &desc,
			Icon:  &icon,
		}
		if err := role.Insert(); err != nil {
			global.Logger.Error(err)
			os.Exit(1)
		}
	}

	{
		menuGroup := &models.RbacMenuGroup{}
		if err := menuGroup.SelectByCode("*", "system"); err != nil {
			global.Logger.Error(err)
			os.Exit(0)
		}

		model := &models.RbacRoleAssocMenuGroup{
			Model:       models.Model{},
			RoleId:      &roleId,
			MenuGroupId: menuGroup.Id,
		}
		if err := model.Insert(); err != nil {
			global.Logger.Error(err)
		}
	}

	{
		menu := &models.RbacMenu{}
		if err := menu.SelectByCode("*", "account"); err != nil {
			global.Logger.Error(err)
			os.Exit(0)
		}

		model := &models.RbacRoleAssocMenu{
			Model:  models.Model{},
			RoleId: &roleId,
			MenuId: menu.Id,
		}
		if err := model.Insert(); err != nil {
			global.Logger.Error(err)
		}
	}

	{
		menu := &models.RbacMenu{}
		if err := menu.SelectByCode("*", "role"); err != nil {
			global.Logger.Error(err)
			os.Exit(0)
		}

		model := &models.RbacRoleAssocMenu{
			Model:  models.Model{},
			RoleId: &roleId,
			MenuId: menu.Id,
		}
		if err := model.Insert(); err != nil {
			global.Logger.Error(err)
		}
	}

	{
		menu := &models.RbacMenu{}
		if err := menu.SelectByCode("*", "interface"); err != nil {
			global.Logger.Error(err)
			os.Exit(0)
		}

		model := &models.RbacRoleAssocMenu{
			Model:  models.Model{},
			RoleId: &roleId,
			MenuId: menu.Id,
		}
		if err := model.Insert(); err != nil {
			global.Logger.Error(err)
		}
	}
}

func createApi() {
	prefix := "/admin/v1"
	{
		method := "POST"
		path := fmt.Sprintf("%s/login",prefix)
		api := &models.RbacApi{
			Model:  models.Model{},
			Method: &method,
			Path:   &path,
		}
		if err := api.Insert(); err != nil {
			global.Logger.Error(err)
		}
	}
	{
		method := "GET"
		path := fmt.Sprintf("%s/logout",prefix)
		api := &models.RbacApi{
			Model:  models.Model{},
			Method: &method,
			Path:   &path,
		}
		if err := api.Insert(); err != nil {
			global.Logger.Error(err)
		}
	}
	{
		method := "GET"
		path := fmt.Sprintf("%s/auth_code",prefix)
		api := &models.RbacApi{
			Model:  models.Model{},
			Method: &method,
			Path:   &path,
		}
		if err := api.Insert(); err != nil {
			global.Logger.Error(err)
		}
	}
	{
		method := "POST"
		path := fmt.Sprintf("%s/login",prefix)
		api := &models.RbacApi{
			Model:  models.Model{},
			Method: &method,
			Path:   &path,
		}
		if err := api.Insert(); err != nil {
			global.Logger.Error(err)
		}
	}
	{
		method := "POST"
		path := fmt.Sprintf("%s/rbac/role/create",prefix)
		api := &models.RbacApi{
			Model:  models.Model{},
			Method: &method,
			Path:   &path,
		}
		if err := api.Insert(); err != nil {
			global.Logger.Error(err)
		}
	}
	{
		method := "POST"
		path := fmt.Sprintf("%s/rbac/role/update",prefix)
		api := &models.RbacApi{
			Model:  models.Model{},
			Method: &method,
			Path:   &path,
		}
		if err := api.Insert(); err != nil {
			global.Logger.Error(err)
		}
	}
	{
		method := "GET"
		path := fmt.Sprintf("%s/rbac/role/select",prefix)
		api := &models.RbacApi{
			Model:  models.Model{},
			Method: &method,
			Path:   &path,
		}
		if err := api.Insert(); err != nil {
			global.Logger.Error(err)
		}
	}

	{
		method := "POST"
		path := fmt.Sprintf("%s/rbac/role/create",prefix)
		api := &models.RbacApi{
			Model:  models.Model{},
			Method: &method,
			Path:   &path,
		}
		if err := api.Insert(); err != nil {
			global.Logger.Error(err)
		}
	}
}
