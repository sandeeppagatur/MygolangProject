package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sandeeppagatur/MyGolangProject/services"
	"log"
	"net/http"
	"text/template"
)
var Tmpl = template.Must(template.ParseGlob("form/*"))

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/event", services.CreateEvent).Methods("POST")
	router.HandleFunc("/events", services.GetAllEvents).Methods("GET")
	router.HandleFunc("/events/{id}", services.GetOneEvent).Methods("GET")
	router.HandleFunc("/events/{id}", services.UpdateEvent).Methods("PATCH")
	router.HandleFunc("/events/{id}", services.DeleteEvent).Methods("DELETE")
	router.HandleFunc("/user", services.CreateUser).Methods("POST")
	router.HandleFunc("/users", services.GetAllUsers).Methods("GET")
	router.HandleFunc("/events/{id}", services.FindUser).Methods("GET")
	//router.HandleFunc("/events/{id}", services.UpdateEvent).Methods("PATCH")
	//router.HandleFunc("/events/{id}", services.DeleteEvent).Methods("DELETE")
	//router.HandleFunc("/", Index)
	//router.HandleFunc("/show", Show)
	//router.HandleFunc("/new", New)
	//router.HandleFunc("/edit", Edit)
	//router.HandleFunc("/insert", Insert)
	//router.HandleFunc("/update", Update)
	//router.HandleFunc("/delete", Delete)
	//http.ListenAndServe(":8080", nil)
	log.Fatal(http.ListenAndServe(":8080", router))
}
