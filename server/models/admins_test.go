package models

import (
	"api/config"
	"api/global"
	"api/initialize"
	"api/models/bak"
	"encoding/json"
	"fmt"
	"testing"
)

func TestAdmins(t *testing.T) {
	config.LocalLoad()
	if err := initialize.InitMysqlMaster(config.GetConfig().MysqlMaster); err != nil {
		panic(err)
	}
	//admin := &[]Admins{}
	//res := global.DBMaster.Find(admin)
	//if res.Error != nil{
	//	panic(res.Error)
	//}
	//fmt.Println(admin)

	//roleMenu := &[]AuthorityRolesMenus{}
	roleMenuId := &[]bak.AuthorityRolesMenus{}
	//根据role_id获取菜单树和相应的api
	res := global.DBMaster.Select("menu_id").Find(roleMenuId, "role_id in (?)", []int64{1, 2})
	if res.Error != nil {
		panic(res.Error)
	}
	menuId := []uint64{}
	for _, i := range *roleMenuId {
		menuId = append(menuId, *i.MenuId)
	}
	menu := &[]bak.Menus{}
	menuGroupId := []uint64{}
	//fmt.Println(menuId)
	{
		res := global.DBMaster.Select("*").Find(menu, "id in (?)", menuId)
		if res.Error != nil {
			panic(res.Error)
		}
		tempMap := map[uint64]byte{}
		for _, i := range *menu {
			l := len(tempMap)
			tempMap[*i.GroupId] = 0
			if len(tempMap) != l {
				menuGroupId = append(menuGroupId, *i.GroupId)
			}
		}
		//fmt.Println(menuGroupId)
	}

	menuGroups := &[]bak.MenusGroups{}
	{
		res := global.DBMaster.Select("*").Find(menuGroups, "id in (?)", menuId)
		if res.Error != nil {
			panic(res.Error)
		}

		for _, i := range *menuGroups {
			fmt.Println(*i.GroupName)
		}
	}

	authorityRolesApi := &[]bak.AuthorityRolesApis{}
	{
		res := global.DBMaster.Find(authorityRolesApi, "role_id in (?)", []int64{1, 2})
		if res.Error != nil {
			panic(res.Error)
		}
		//for _, i := range *authorityRolesApi {
		//	fmt.Println(*i.ApiMethod, *i.MenuId)
		//}
	}

	{
		type menuGroup struct {
			Group string
			Icon  string
		}
		type pathList struct {
			Path   string
			Method string
		}
		type menuList struct {
			Name     string
			Icon     string
			PathList []pathList
		}
		type list struct {
			MenuGroup menuGroup
			MenuList  []menuList
		}
		menuData := []list{}
		for _, mg := range *menuGroups {
			menuListData := list{}
			menuListData.MenuGroup.Group = *mg.GroupName
			temp := menuList{}
			for _, m := range *menu {
				if *mg.Id == *m.GroupId {
					pL := pathList{}
					for _, ara := range *authorityRolesApi {
						if *ara.MenuId == *m.Id {
							pL.Method = *ara.ApiMethod
							pL.Path = *ara.ApiPath
						}
					}
					temp.Name = *m.Name
					temp.Icon = *m.Icon
					temp.PathList = append(temp.PathList, pL)
				}
			}
			menuListData.MenuList = append(menuListData.MenuList, temp)
			menuData = append(menuData, menuListData)

		}

		b, err := json.Marshal(menuData)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))
	}
	//判断角色是否有权限
}

//func TestAdmins(t *testing.T) {
//	config.LocalLoad()
//	if err := initialize.InitMysqlMaster(config.GetConfig().MysqlMaster); err != nil {
//		log.Error(err)
//	}
//
//	userType := uint64(1)
//	username := "gaoshou"
//	password := "123"
//	nickname := "123123"
//
//	admin := &Admins{
//		Id:        nil,
//		Enable:    nil,
//		State:     nil,
//		UserType:  &userType,
//		Username:  &username,
//		Password:  &password,
//		Nickname:  &nickname,
//		CreatedAt: nil,
//		UpdatedAt: nil,
//	}
//
//	// 添加
//	{
//		res := global.DBMaster.Create(admin)
//		if res.Error != nil {
//			panic(res.Error)
//		}
//	}
//
//	// 更新
//	{
//		enable := true
//		//state := uint64(1)
//		admin := &Admins{
//			Id:        nil,
//			Enable:    &enable,
//			State:     nil,
//			UserType:  nil,
//			Username:  nil,
//			Password:  nil,
//			Nickname:  nil,
//			CreatedAt: nil,
//			UpdatedAt: nil,
//		}
//		res := global.DBMaster.Model(admin).Where("id=?", 1).Updates(admin)
//		if res.Error != nil {
//			panic(res.Error)
//		}
//	}
//
//	// 查询
//	{
//		admin := &Admins{}
//		res := global.DBMaster.First(admin, "id=?", 1)
//		if res.Error != nil {
//			panic(res.Error)
//		}
//		fmt.Println(admin)
//	}
//
//	{
//		admin := &Admins{}
//		res := global.DBMaster.Delete(admin, "id=?", 2)
//		if res.Error != nil {
//			panic(res.Error)
//		}
//	}
//}
