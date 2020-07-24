package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/sandeeppagatur/MyGolangProject/models"
	"log"
)
var Database *gorm.DB
func GetDBObject() *gorm.DB {
	if Database == nil {
		Database, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/Grocery?charset=utf8&parseTime=True")
		if err !=nil {
			log.Fatal("db connection failed")
		}else {
			Load(Database)
			return Database
		}
	}
	return Database
}




func Load(db *gorm.DB) {
    if db != nil{
		err := db.Debug().AutoMigrate(&models.Event{}).Error
		if err != nil {
			log.Fatalf("cannot migrate table: %v", err)
		}

		err = db.Debug().AutoMigrate(&models.User{}).Error
		if err != nil {
			log.Fatalf("cannot migrate table: %v", err)
		}
		err = db.Debug().AutoMigrate(&models.Exam{}).Error
		if err != nil {
			log.Fatalf("cannot migrate table: %v", err)
		}
		err = db.Debug().AutoMigrate(&models.Question{}).Error
		if err != nil {
			log.Fatalf("cannot migrate table: %v", err)
		}
		err = db.Debug().AutoMigrate(&models.QOption{}).Error
		if err != nil {
			log.Fatalf("cannot migrate table: %v", err)
		}
		err = db.Debug().AutoMigrate(&models.Answer{}).Error
		if err != nil {
			log.Fatalf("cannot migrate table: %v", err)
		}

	}


}
