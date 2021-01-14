package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Data struct {
	RbacMenuGroup []RbacMenuGroup `rbac_menu_group`
}

type RbacMenuGroup struct {
	Name     string
	Icon     string
	Sort     int
	RbacMenu []RbacMenu `rbac_menu`
}
type RbacMenu struct {
	Code string
	Sort int
	Name string
	Icon string
	Desc string
}

type RbacApi struct {
	RbacApi []Api `yaml:"rbac_api"`
}

type Api struct {
	Method string
	Path   string
}

func main() {
	yamlFile, err := ioutil.ReadFile("menu_group.yaml")
	yamlFileApi, err := ioutil.ReadFile("api.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	data := &Data{}
	api := &RbacApi{}
	err = yaml.Unmarshal(yamlFile, data)
	err = yaml.Unmarshal(yamlFileApi, api)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	str, err := json.Marshal(data)
	apiStr, err := json.Marshal(api)
	if err != nil {
		log.Fatalf("json marshal: %v", err)
	}
	fmt.Println(string(str))
	fmt.Println(string(apiStr))
}
