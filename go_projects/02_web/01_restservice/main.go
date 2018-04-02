package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pjpillai/MyProjects/go_projects/02_web/01_restservice/route"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/people", route.GetPeople).Methods("GET")
	router.HandleFunc("/people/", route.GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", route.GetPerson).Methods("GET")
	router.HandleFunc("/people/", route.CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", route.UpdatePerson).Methods("PUT")
	router.HandleFunc("/people/{id}", route.DeletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}
