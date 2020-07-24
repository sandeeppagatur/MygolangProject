package repository

import (
	"github.com/sandeeppagatur/MyGolangProject/models"
)



func CreateNewQuestion(event *models.Question)  {
	db:=GetDBObject()
	if db != nil {
		db.Create(event)
	}

}
func FindAllQuestions() *[]models.Question {
	var err error
	questions := []models.Question{}
	db:=GetDBObject()
	if db != nil {
		err = db.Debug().Model(&models.Question{}).Limit(100).Find(&questions).Error
		if err != nil {
			return &[]models.Question{}
		}
	}

	return &questions
}

