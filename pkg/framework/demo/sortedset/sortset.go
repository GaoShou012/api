package main

import (
	"fmt"
	"framework/plugin/sortedset/redis_sortdset"
	"github.com/go-redis/redis"
	"net/url"
	"runtime"
	"strconv"
	"time"
)

func connect(dns string) (*redis.Client, error) {
	u, err := url.Parse(dns)
	if err != nil {
		return nil, err
	}

	addr := u.Host
	//username := u.User.Username()
	password, _ := u.User.Password()

	params, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil, err
	}

	var poolMax int
	{
		arr, ok := params["PoolMax"]
		if !ok {
			poolMax = runtime.NumCPU()
		} else {
			num, err := strconv.Atoi(arr[0])
			if err != nil {
				return nil, err
			}
			poolMax = num
		}
	}

	var poolMin int
	{
		arr, ok := params["PoolMin"]
		if !ok {
			poolMin = runtime.NumCPU()
		} else {
			num, err := strconv.Atoi(arr[0])
			if err != nil {
				return nil, err
			}
			poolMin = num
		}
	}

	// new redis_sortdset client
	redisClient := redis.NewClient(&redis.Options{
		Network:            "",
		Addr:               addr,
		Dialer:             nil,
		OnConnect:          nil,
		Password:           password,
		DB:                 0,
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        10 * time.Second,
		ReadTimeout:        30 * time.Second,
		WriteTimeout:       30 * time.Second,
		PoolSize:           poolMax,
		MinIdleConns:       poolMin,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
	})

	// ping redis_sortdset server
	if _, err := redisClient.Ping().Result(); err != nil {
		return nil, err
	}

	return redisClient, nil
}
func main() {

	var params struct {
		Page       uint
		PageSize   uint
		MerchantId uint64
		Username   string
		StartAt    string
		EndAt      string
	}
	//t := fmt.Sprintf("%d", 0)
	fmt.Println(params)

	dns := fmt.Sprintf("redis://:@192.168.0.233:17001?Db=0&PoolMax=100&PoolMin=10")
	r, err := connect(dns)
	if err != nil {
		panic(err)
	}
	p := redis_sortdset.New(redis_sortdset.WithRedisClient(r))
	{
		err = p.SetItem("topic1", "wdnmd1", 1)
		err = p.SetItem("topic1", "wdnmd", 0)
		err = p.SetItem("topic1", "wdnmd4", 90)
		err = p.SetItem("topic1", "wdnmd2", 100)
		if err != nil {
			panic(err)
		}
	}

	item, err := p.Find("topic1", 0, 10)
	if err != nil {
		panic(err)
	}
	for k, v := range item {
		fmt.Println(k, v.Val(), v.Key())
	}
	lens := p.Len("topic1")
	fmt.Println(lens)

	{
		exist, err := p.Exists("topic1z")
		fmt.Println(exist)

		exist, err = p.ExistsItem("topic1", "wdnmd")
		if err != nil {
			panic(err)
		}
		fmt.Println(exist)
	}

	{
		i, _ := p.GetOffset("topic1", "wdnmd1")
		//fmt.Println(err)
		n, _ := p.GetOffsetN("topic1", "wdnmd")
		fmt.Println(i, n)
	}

	item, err = p.GetItemFormPositive("topic1")
	for k, v := range item {
		fmt.Println(k, v.Val(), v.Key())
	}
	item, err = p.GetItemFromNegative("topic1")
	for k, v := range item {
		fmt.Println(k, v.Val(), v.Key())
	}

	//docker exec -it CONTAINERID redis-cli
	//t := sortedset.Sortedset.Exists()
	{
		bo, err := p.DelKey("topic1", "wdnmd")
		if err != nil {
			panic(err)
		}
		fmt.Println(bo)
	}
	{
		bo, err := p.DelItem("topic1")
		if err != nil {
			panic(err)
		}
		fmt.Println(bo)
	}
	{

	}
}
