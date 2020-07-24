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

func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	var newQuestion models.Question
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &newQuestion)
	repository.CreateNewQuestion(&newQuestion)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newQuestion)
}

func UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	//eventID := mux.Vars(r)["id"]
	var updatedEvent models.Event

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedEvent)


}

func DeleteQuestion(w http.ResponseWriter, r *http.Request) {

}

func FindQuestion(w http.ResponseWriter, r *http.Request) {

}

func GetAllQuestions(w http.ResponseWriter, r *http.Request) {
	questions:=repository.FindAllQuestions()
	json.NewEncoder(w).Encode(questions)
}

