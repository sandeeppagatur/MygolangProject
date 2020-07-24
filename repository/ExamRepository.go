package repository

import (
	"github.com/sandeeppagatur/MyGolangProject/models"
)



func CreateNewExam(event *models.Exam)  {
	db:=GetDBObject()
	if db != nil {
		db.Create(event)
	}

}
func FindAllExams() *[]models.Exam {
	var err error
	exams := []models.Exam{}
	db:=GetDBObject()
	if db != nil {
		err = db.Debug().Model(&models.Exam{}).Limit(100).Find(&exams).Error
		if err != nil {
			return &[]models.Exam{}
		}
	}

	return &exams
}

