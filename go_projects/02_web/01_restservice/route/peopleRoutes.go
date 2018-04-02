package route

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pjpillai/MyProjects/go_projects/02_web/01_restservice/model"
	"github.com/pjpillai/MyProjects/go_projects/02_web/01_restservice/service"
)

var requestCount int

// GetPeople ...
func GetPeople(w http.ResponseWriter, r *http.Request) {
	requestCount++
	log.Println("Got request - GetPeople", requestCount)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(service.GetPeople())
}

// GetPerson ...
func GetPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	person := service.GetPerson(id)

	json.NewEncoder(w).Encode(&person)
}

// CreatePerson ...
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(r)
	var person model.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	//person.ID = params["id"]
	isCreated := service.CreatePerson(&person)
	if isCreated == false {
		w.WriteHeader(http.StatusBadRequest)
		var er = model.ErrorResponse{StatusCode: 404, Message: "Invalid Data"}
		json.NewEncoder(w).Encode(er)
	} else {
		json.NewEncoder(w).Encode(&person)
	}

}

// UpdatePerson ...
func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var person model.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	isUpdated := service.UpdatePerson(&person)
	if isUpdated == false {
		w.WriteHeader(http.StatusBadRequest)
		var er = model.ErrorResponse{StatusCode: 404, Message: "Invalid Data"}
		json.NewEncoder(w).Encode(er)
	} else {
		json.NewEncoder(w).Encode(&person)
	}

}

// DeletePerson ...
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var person model.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	isDeleted := service.DeletePerson(&person)
	if isDeleted == false {
		w.WriteHeader(http.StatusBadRequest)
		var er = model.ErrorResponse{StatusCode: 404, Message: "Invalid Data"}
		json.NewEncoder(w).Encode(er)
	} else {
		json.NewEncoder(w).Encode(&person)
	}

}
