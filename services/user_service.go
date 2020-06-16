package services

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm"
	"github.com/sandeeppagatur/MyGolangProject/models"
	"github.com/sandeeppagatur/MyGolangProject/repository"
	"io/ioutil"
	"net/http"
)

type allUsers []models.Event



func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newUser)
	//events = append(events, newEvent)
	repository.CreateNewUser(&newUser)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newUser)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	//eventID := mux.Vars(r)["id"]
	var updatedEvent models.Event

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedEvent)

	//for i, singleEvent := range events {
	//	if singleEvent.ID == eventID {
	//		singleEvent.Title = updatedEvent.Title
	//		singleEvent.Description = updatedEvent.Description
	//		events = append(events[:i], singleEvent)
	//		json.NewEncoder(w).Encode(singleEvent)
	//	}
	//}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	//ventID := mux.Vars(r)["id"]

	//for i, singleEvent := range events {
	//	if singleEvent.ID == eventID {
	//		events = append(events[:i], events[i+1:]...)
	//		fmt.Fprintf(w, "The event with ID %v has been deleted successfully", eventID)
	//	}
	//}
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	//eventID := mux.Vars(r)["id"]

	//for _, singleEvent := range users {
	//	if singleEvent.ID == eventID {
	//		json.NewEncoder(w).Encode(singleEvent)
	//	}
	//}
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users:=repository.FindAllUsers()
	json.NewEncoder(w).Encode(users)
}
