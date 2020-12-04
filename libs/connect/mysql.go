package connect

import (
	"api/config"
	"api/libs/logs"
	"api/models"
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB


func InitDB() *gorm.DB {
	driverName := config.GetConfig().Mysql.MysqlDriverName
	host := config.GetConfig().Mysql.MysqlHost
	port := config.GetConfig().Mysql.MysqlPort
	database := config.GetConfig().Mysql.MysqlDatabase
	username := config.GetConfig().Mysql.MysqlUsername
	password := config.GetConfig().Mysql.MysqlPassword
	charset := config.GetConfig().Mysql.MysqlCharset
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		logs.Fatal(err,"数据库连接出错")
	}

	db.AutoMigrate(&models.Admins{})
	db.AutoMigrate(&models.Menu{})
	db.AutoMigrate(&models.MenuAction{})
	db.AutoMigrate(&models.MenuActionResource{})
	db.AutoMigrate(&models.Role{})
	db.AutoMigrate(&models.RoleMenu{})
	db.AutoMigrate(&models.UserRole{})
	db.AutoMigrate(&models.User{})


	DB = db

	return db
}

func GetDB() *gorm.DB {
	return DB
}
