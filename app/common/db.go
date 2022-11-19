package common

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 初始化数据库连接
func init() {
	// copy from https://gorm.io/zh_CN/docs/connecting_to_the_database.html
	// 数据库服务器、端口、用户名和密码，以及选项参数，组成DSN
	// DSN Data Source Name，数据源名称
	dsn := "root:demo@tcp(127.0.0.1:3316)/cms?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	// 数据库连接成功
	DB = db
}
