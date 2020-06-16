package repository

import (
	"fmt"
	"github.com/sandeeppagatur/MyGolangProject/models"
	"log"
)

func CreateNewEvent(event *models.Event)  {
db:=GetDBObject()
	if db != nil {
		db.Create(event)
	}

}


func UpdateEvent()  {

}


func GetEvents() *[]models.Event  {
	var err error
	events := []models.Event{}
	db:=GetDBObject()
	if db != nil {
		err = db.Debug().Model(&models.Event{}).Limit(100).Find(&events).Error
		fmt.Println("get events:: ",events)

		if err != nil {
			log.Fatal("some error occured while fetching events")
			return &[]models.Event{}
		}
	}
    fmt.Println("get events:: ",events)

	return &events
}

func DeleteEvent()  {

}