package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sandeeppagatur/MyGolangProject/services"
	"log"
	"net/http"
)


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
	log.Fatal(http.ListenAndServe(":8080", router))
}





