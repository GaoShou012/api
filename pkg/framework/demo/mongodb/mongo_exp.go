package main

import (
	"fmt"
	"framework/utils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func initMgo() *mgo.Session {
	mongoDbClient, err := utils.MongoDbClient("mongodb://admin:123456@192.168.0.20:27017")
	if err != nil {
		panic(err)
	}

	mongoDbClient.SetMode(mgo.Primary, true)
	mongoDbClient.SetPoolLimit(4)
	return mongoDbClient
}

func insert() {
	mgo := initMgo()
	collection := mgo.DB("test").C("test1")
	type test struct {
		UserName string
		Phone    []int64
		Subject  string
		Sort     int64
	}
	//var c []interface{}
	//for i := 0; i < 20; i++ {
	//	c = append(c, test{
	//		UserName: fmt.Sprintf("jax%d", i),
	//		Phone:    []int64{1, 2, 34},
	//		Subject:  fmt.Sprintf("subject%d", i),
	//		Sort:     int64(i),
	//	})
	//}
	//err := collection.Insert(c)

	for i := 0; i < 20; i++ {
		c := test{
			UserName: fmt.Sprintf("jax%d", i),
			Phone:    []int64{1, 2, 34},
			Subject:  fmt.Sprintf("subject%d", i),
			Sort:     int64(i),
		}
		err := collection.Insert(c)
		if err != nil {
			panic(err)
		}
	}

	//c := test{
	//	UserName: "jax",
	//	Phone:    []int64{1, 2, 3, 4},
	//	Subject:  "test",
	//}
	//err := collection.Insert(c)
	//
	//if err != nil {
	//	panic(err)
	//}
}

func insertAll() {
	mgo := initMgo()
	collection := mgo.DB("test").C("test2")
	var doc []interface{}
	for i := 0; i < 100; i++ {
		doc = append(doc, bson.M{
			"username": fmt.Sprintf("jax%d", i),
			"role_id":  i,
			"sort":     i,
			"menus":    []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
		})
	}
	err := collection.Insert(doc...)
	fmt.Println(err)
}

func findOne() {
	mgo := initMgo()
	collection := mgo.DB("test").C("test")
	type test struct {
		UserName *string
		Phone    *[]int64
		Subject  *string
	}
	iTest := &test{}
	//select 条件写法  只需要username字段->bson.M{"username":1}  除了username外都要->bson.M{"username":0} 所有->nil
	err := collection.Find(bson.M{"username": "jax"}).Select(nil).One(iTest)
	if err != nil {
		panic(err)
	}
	fmt.Println(iTest)
}

//分页查询
func findWithPage() {
	mgo := initMgo()
	collection := mgo.DB("test").C("test")
	type test struct {
		UserName *string
		Phone    *[]int64
		Subject  *string
	}
	iTest := &[]test{}
	page := 2
	pageSize := 10
	collection.Find(nil).Select(bson.M{"username": 1}).Skip((page - 1) * pageSize).Limit(pageSize).All(iTest)
	fmt.Println(iTest)
}

//更新
func update() {
	mgo := initMgo()
	collection := mgo.DB("test").C("test")
	err := collection.Update(bson.M{"username": "jax1"}, bson.M{"$set": bson.M{"subject": "update"}})
	if err != nil {
		panic(err)
	}
}

//复杂查询时
func pipe() {
	mgo := initMgo()
	collection := mgo.DB("test").C("test1")
	type test struct {
		UserName *string
		Phone    *[]int64
		Subject  *string
		Sort     int64
	}
	result := []test{}
	page := 1
	pageSize := 10
	pipM := []bson.M{
		//{"$match": bson.M{"username": "jax"}},//条件
		{"$skip": (page - 1) * pageSize}, //offset
		{"$limit": pageSize},             //limit
		{"$sort": bson.M{"sort": -1}},    //排序 -1反序 1正序
	}
	collection.Pipe(pipM).All(&result)
	fmt.Println(result)
}
func main() {
	insert()
	//findOne()
	//findWithPage()
	//update()
	//pipe()
	//insertAll()
}
