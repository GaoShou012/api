package utils

import (
	"gopkg.in/mgo.v2"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func MongoDbClient(dns string) (*mgo.Session, error) {
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

	//var poolMin int
	//{
	//	arr, ok := params["PoolMin"]
	//	if !ok {
	//		poolMin = runtime.NumCPU()
	//	} else {
	//		num, err := strconv.Atoi(arr[0])
	//		if err != nil {
	//			return nil, err
	//		}
	//		poolMin = num
	//	}
	//}
	//var poolMax int
	//{
	//	arr, ok := params["PoolMax"]
	//	if !ok {
	//		poolMax = runtime.NumCPU()
	//	} else {
	//		num, err := strconv.Atoi(arr[0])
	//		if err != nil {
	//			return nil, err
	//		}
	//		poolMax = num
	//	}
	//}

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

	//isSsl := false
	//
	//url := fmt.Sprintf("mongodb://admin:123456@192.168.0.2:27017,192.168.0.2:27017")
	//dailInfo, err := mgo.ParseURL(url)
	//if err != nil {
	//	return nil, err
	//}
	//
	//if isSsl {
	//
	//}

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
