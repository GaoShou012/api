package main

import (
	"crypto/md5"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func InitMongoDb(dns string) (*mgo.Session, error) {
	u, err := url.Parse(dns)
	if err != nil {
		return nil, err
	}

	addr := u.Host
	username := u.User.Username()
	password, _ := u.User.Password()
	params, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil, err
	}
	var timeout time.Duration
	{
		arr, ok := params["Timeout"]
		if !ok {
			timeout = time.Second * 10
		} else {
			num, err := strconv.Atoi(arr[0])
			if err != nil {
				return nil, err
			}
			timeout = time.Second * time.Duration(num)
		}
	}
	dailInfo := &mgo.DialInfo{
		Addrs:          strings.Split(addr, ","),
		Direct:         false,
		Timeout:        timeout,
		FailFast:       false,
		Database:       "",
		ReplicaSetName: "",
		Source:         "",
		Service:        "",
		ServiceHost:    "",
		Mechanism:      "",
		Username:       username,
		Password:       password,
		PoolLimit:      0,
		DialServer:     nil,
		Dial:           nil,
	}
	session, err := mgo.DialWithInfo(dailInfo)
	if err != nil {
		return nil, err
	}

	session.SetSyncTimeout(1 * time.Minute)
	session.SetSocketTimeout(1 * time.Minute)

	return session, nil
}

type Test struct {
	UserName string
	Phone    int64
	Subject  string
}

func main() {
	mongoDbClient, err := InitMongoDb("mongodb://admin:123456@192.168.0.20:27017")
	if err != nil {
		panic(err)
	}
	defer mongoDbClient.Close()

	//选择数据库  选择数据库 没有是就回自动创建一个
	db := mongoDbClient.DB("test")
	//创建集合 如存在直接返回
	collection := db.C("person")
	count, err := collection.Count()
	if err != nil {
		panic(err)
	}
	fmt.Println("count:", count)

	c1 := &Test{
		UserName: "jax",
		Phone:    112220,
		Subject:  "ccc",
	}
	//插入 insert(c1,c2,c3,.....)
	err = collection.Insert(c1)
	if err != nil {
		panic(err)
	}
	fmt.Println(collection.Count())

	{
		c := Test{}
		//查单挑数据
		err := collection.Find(bson.M{"username": "jax"}).One(&c)
		if err != nil {
			panic(err)
		}
		fmt.Println(c)
	}
	{
		//查多条
		var All []Test
		c := Test{}
		iter := collection.Find(nil).Iter()
		for iter.Next(&c) {
			All = append(All, c)
		}
	}

	{
		//修改集合中的数据 如果有多个只修改第一个
		err := collection.Update(bson.M{"username": "jax"}, bson.M{"$set": bson.M{"subject": "aaa"}})
		if err != nil {
			panic(err)
		}
		//修改所以，mgo字段没有类型规定
		info, err := collection.UpdateAll(bson.M{"subject": "ccc"}, bson.M{"$set": bson.M{"phone": "1230c@@000"}})
		if err != nil {
			panic(err)
		}
		fmt.Println(info)
		//字段不匹配回在集合中新增字段
		info, err = collection.UpdateAll(bson.M{"subject": "ccc"}, bson.M{"$set": bson.M{"pHone": "7879879"}})
		if err != nil {
			panic(err)
		}
		fmt.Println(info)
		//匹配不到
		info, err = collection.UpdateAll(bson.M{"Subject": "ccc"}, bson.M{"$set": bson.M{"pHone": "7879879"}})
		if err != nil {
			panic(err)
		}
		fmt.Println(info)
	}

	fmt.Println(md5.New())
}
