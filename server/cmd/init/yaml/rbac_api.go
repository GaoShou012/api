package yaml

import (
	"api/models"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path"
	"runtime"
)

type RbacApiData struct {
	RbacApi []models.RbacApi `yaml:"rbac_api"`
}

type RbacApi struct {
	Method string
	Path   string
}

func GetRbacApiData() (*RbacApiData, error) {
	var yamlFilePath string
	data := &RbacApiData{}

	{
		_, filename, _, _ := runtime.Caller(0)
		yamlFilePath = path.Join(path.Dir(filename), "rbac_api.yaml")
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