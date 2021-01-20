package yaml

import (
	"api/models"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path"
	"runtime"
)

type RbacMenuData struct {
	RbacMenus []RbacMenuGroupRow `yaml:"rbac_menus"`
}

type RbacMenuGroupRow struct {
	Group models.RbacMenuGroup
	Menus []models.RbacMenu
}

func GetRbacMenuData() (*RbacMenuData, error) {
	var yamlFilePath string
	data := &RbacMenuData{}

	{
		_, filename, _, _ := runtime.Caller(0)
		yamlFilePath = path.Join(path.Dir(filename), "rbac_menu.yaml")
		fmt.Printf("yaml路径:%s\n", yamlFilePath)
	}

	f, err := ioutil.ReadFile(yamlFilePath)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(f, data); err != nil {
		return nil, err
	}
	return data, nil
}
