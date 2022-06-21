package models

import (
	"api_test/config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DbConnect() *gorm.DB {
    db, err := gorm.Open(config.Config.SQLDriver, "user:password@tcp(test_db:3306)/test_db?charset=utf8&parseTime=True&loc=Local")
    fmt.Println(config.Config.SQLDriver)
    if err != nil {
        panic(err.Error())
    }
	fmt.Println("db connected: ", &db)
	return db
}

