package main

import (
	"fmt"
	"framework/utils"
	"gopkg.in/mgo.v2"
)

func main() {
	mongoDbClient, err := utils.MongoDbClient("mongodb://admin:123456@192.168.0.20:27017")
	if err != nil {
		panic(err)
	}

	mongoDbClient.SetMode(mgo.Primary,true)
	mongoDbClient.SetPoolLimit(4)
	fmt.Println(mongoDbClient)
	//fmt.Println(mongoDbClient.)
	//fmt.Println(mongoDbClient)
}
