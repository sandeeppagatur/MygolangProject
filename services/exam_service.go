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

func CreateExam(w http.ResponseWriter, r *http.Request) {
	var newQuestion models.Exam
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &newQuestion)
	repository.CreateNewExam(&newQuestion)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newQuestion)
}



func GetAllExams(w http.ResponseWriter, r *http.Request) {
	questions:=repository.FindAllExams()
	json.NewEncoder(w).Encode(questions)
}

