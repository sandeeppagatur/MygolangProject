package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
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
	//http.Handle("/", &WithCORS{router})

	router.HandleFunc("/", homeLink)
	router.HandleFunc("/newuserform", services.New)
	router.HandleFunc("/event", services.CreateEvent).Methods("POST")
	router.HandleFunc("/events", services.GetAllEvents).Methods("GET")
	router.HandleFunc("/events/{id}", services.GetOneEvent).Methods("GET")
	router.HandleFunc("/events/{id}", services.UpdateEvent).Methods("PATCH")
	router.HandleFunc("/events/{id}", services.DeleteEvent).Methods("DELETE")
	router.HandleFunc("/user", services.CreateUser).Methods("POST")
	router.HandleFunc("/users", services.GetAllUsers).Methods("GET")
	router.HandleFunc("/events/{id}", services.FindUser).Methods("GET")
	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":8080", handler))
}

//type WithCORS struct {
//	r *mux.Router
//
//}
//
//func (s *WithCORS) ServeHTTP(res http.ResponseWriter, req *http.Request) {
//
//		res.Header().Set("Access-Control-Allow-Origin", "*")
//		res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
//		res.Header().Set("Access-Control-Allow-Headers",
//			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
//
//	// Stop here for a Preflighted OPTIONS request.
//	if req.Method == "OPTIONS" {
//		return
//	}
//	// Lets Gorilla work
//	s.r.ServeHTTP(res, req)
//}
