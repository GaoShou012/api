package main

import (
	"fmt"
	"framework/utils"
)

func main() {
	conn, err := utils.SqlMysql("root:123456@tcp(127.0.0.1:13306)/api?charset=utf8mb4&loc=Local&parseTime=True", 10, 100)
	if err != nil {
		panic(err)
	}
	fmt.Println("连接成功:", conn)
}
