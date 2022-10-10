package database

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"main.go/config"
	user "main.go/features/book/repository"
	book "main.go/features/user/repository"
)

func InitDB(c *config.AppConfig) *gorm.DB {
	str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DBUser,
		c.DBPwd,
		c.DBHost,
		c.DBPort,
		c.DBName,
	)

	db, err := gorm.Open(mysql.Open(str), &gorm.Config{})
	if err != nil {
		log.Error("db config error :", err.Error())
		return nil
	}
	return db
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&user.Book{})
	db.AutoMigrate(&book.User{})
}
