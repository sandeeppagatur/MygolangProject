package repository

import (
	"github.com/jinzhu/gorm"
	"log"
)
var Database *gorm.DB
func GetDBObject() *gorm.DB {
	if Database != nil {
		Database, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/Grocery?charset=utf8&parseTime=True")
		if err !=nil {
			log.Fatal("db connection failed")
		}
		defer Database.Close()
	}

	return Database
}
