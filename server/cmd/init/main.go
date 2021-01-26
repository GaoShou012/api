package main

import (
	"api/cmd/init/yaml"
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
		//initialize.InitRBAC()
	}
}

func main() {
	createMenus()
	createRole()
	createApi()

	// 角色关联所有菜单
	roleAssocMenu()
	// 角色关联所有API
	roleAssocApi()

	// 创建管理员账号
	// 管理员，关联超级管理员角色
	createAdmins()

	fmt.Println("执行完成")
}

func createAdmins() {
	role := models.RbacRole{}
	if err := role.SelectByName("*", "超级管理员"); err != nil {
		global.Logger.Error(err)
		os.Exit(0)
	}
	roleId := *role.Id

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
	username := "admin"
	nickname := "管理员"
	password := "123456"
	passwordEncrypted, err := model.EncryptPassword(password)
	roles := fmt.Sprintf("%d", roleId)
	if err != nil {
		panic(err)
	}
	admin := &models.Admins{
		Id:        nil,
		Enable:    &enable,
		Username:  &username,
		Password:  &passwordEncrypted,
		Nickname:  &nickname,
		Roles:     &roles,
		CreatedAt: nil,
		UpdatedAt: nil,
	}
	res := global.DBMaster.Table(model.GetTableName()).Create(admin)
	if res.Error != nil {
		panic(res.Error)
	}
	fmt.Println("创建管理员成功")
}

func createRole() {
	data, err := yaml.GetRbacRoleData()
	if err != nil {
		global.Logger.Error(err)
		os.Exit(0)
	}
	for _, row := range data.RbacRole {
		if err := row.Insert(); err != nil {
			global.Logger.Error(err)
		}
	}
}

func createMenus() {
	data, err := yaml.GetRbacMenuData()
	if err != nil {
		global.Logger.Error(err)
		return
	}

	for _, row := range data.RbacMenus {
		group := row.Group
		if err := group.Insert(); err != nil {
			global.Logger.Error(err)
		}
		if err := group.SelectByName("*", *group.Name); err != nil {
			global.Logger.Error(err)
			return
		}

		for _, menu := range row.Menus {
			menu.GroupId = group.Id
			if err := menu.Insert(); err != nil {
				global.Logger.Error(err)
			}

			if err := menu.SelectByCode("*", *menu.Code); err != nil {
				global.Logger.Error(err)
				continue
			}
		}
	}
}

func createApi() {
	prefix := "/admin/api/v1"

	data, err := yaml.GetRbacApiData()
	if err != nil {
		global.Logger.Error(err)
		return
	}

	for _, row := range data.RbacApi {
		path := fmt.Sprintf("%s%s", prefix, *row.Path)
		row.Path = &path
		if err := row.Insert(); err != nil {
			global.Logger.Error(err)
		}
	}
}

func roleAssocApi() {
	role := models.RbacRole{}
	if err := role.SelectByName("*", "超级管理员"); err != nil {
		global.Logger.Error(err)
		os.Exit(0)
	}
	roleId := *role.Id

	apis := make([]models.RbacApi, 0)
	res := global.DBSlave.Table((&models.RbacApi{}).GetTableName()).Find(&apis)
	if res.Error != nil {
		global.Logger.Error(res.Error)
		return
	}

	for _, api := range apis {
		assoc := &models.RbacRoleAssocApi{
			Model:  models.Model{},
			RoleId: &roleId,
			ApiId:  api.Id,
		}
		assoc.Insert()
	}
}
func roleAssocMenu() {
	role := models.RbacRole{}
	if err := role.SelectByName("*", "超级管理员"); err != nil {
		global.Logger.Error(err)
		os.Exit(0)
	}
	roleId := *role.Id

	{
		rows := make([]models.RbacMenuGroup, 0)
		res := global.DBSlave.Table((&models.RbacMenuGroup{}).GetTableName()).Find(&rows)
		if res.Error != nil {
			global.Logger.Error(res.Error)
			return
		}

		for _, row := range rows {
			assoc := &models.RbacRoleAssocMenuGroup{
				Model:       models.Model{},
				RoleId:      &roleId,
				MenuGroupId: row.Id,
			}
			assoc.Insert()
		}
	}

	{
		rows := make([]models.RbacMenu, 0)
		res := global.DBSlave.Table((&models.RbacMenu{}).GetTableName()).Find(&rows)
		if res.Error != nil {
			global.Logger.Error(res.Error)
			return
		}

		for _, row := range rows {
			assoc := &models.RbacRoleAssocMenu{
				Model:  models.Model{},
				RoleId: &roleId,
				MenuId: row.Id,
			}
			assoc.Insert()
		}
	}
}
