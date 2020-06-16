package repository

import (
	"github.com/sandeeppagatur/MyGolangProject/models"
)



func CreateNewUser(event *models.User)  {
	db:=GetDBObject()
	if db != nil {
		db.Create(event)
	}

}
func FindAllUsers() *[]models.User {
	var err error
	users := []models.User{}
	db:=GetDBObject()
	if db != nil {
		err = db.Debug().Model(&models.User{}).Limit(100).Find(&users).Error
		if err != nil {
			return &[]models.User{}
		}
	}

	return &users
}
