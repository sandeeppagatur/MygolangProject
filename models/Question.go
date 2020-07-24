package models

import "github.com/jinzhu/gorm"

type QuestionType struct{
	MultiSelect string `json:"MultiSelect"`
	SingleSelect string `json:"SingleSelect"`
	Plain string `json:"Plain"`
}

type ExamType struct{
	Quarterly string `json:"Quarterly"`
	Halfyearly string `json:"Halfyearly"`
	Annual string `json:"Annual"`
	Daily  string `json:"Daily"`
	Weekly string `json:"Weekly"`
	Monthly string `json:"Monthly"`

}

type Question struct {
	gorm.Model
	Title       string `json:"Title"`
	Priority   	int   `json:"priority"`
	Level 		string `json:"Level"`
	QueType 	string `json:"QueType"`
	QOptions 	[]QOption `json:"QOptions"` 
	Answers  	[]Answer `json:"Answers"`
	CreatedBy 	int    `json:"CreatedBy"`
	Exams       []*Exam  `gorm:"many2many:exam_questions;"`
}
type QOption struct{
	gorm.Model
	Name string
}
type Answer struct{
	gorm.Model
	Name string
}

type Exam struct{
	gorm.Model
	Type         string `json:"Type"`
	Questions []*Question `json:"questions" gorm:"many2many:exam_questions;"`
	CreatedBy int    `json:"CreatedBy"`
	AssignedTo int   `json:"AssignedTo"`

}
