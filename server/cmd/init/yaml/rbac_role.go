package yaml

import (
	"api/models"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path"
	"runtime"
)

type RbacRoleData struct {
	RbacRole []models.RbacRole `yaml:"rbac_role"`
}

func GetRbacRoleData() (*RbacRoleData, error) {
	var yamlFilePath string
	data := &RbacRoleData{}

	{
		_, filename, _, _ := runtime.Caller(0)
		yamlFilePath = path.Join(path.Dir(filename), "rbac_role.yaml")
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
