package orm

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func InitDB() {
	dsn := os.Getenv("MYSQL_DNS")
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("System: failed to connect database")
	}

	Db.AutoMigrate(&User{})
}