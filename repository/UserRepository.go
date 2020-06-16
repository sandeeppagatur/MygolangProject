package repository

import "github.com/sandeeppagatur/MyGolangProject/models"

func CreateUser()  {
	Database:=GetDBObject()
	Database.Create(&models.User{Code: "L1212", Price: 1000})
}
