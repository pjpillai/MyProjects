package main

import (
	"encoding/json"
	"log"
	"net/http"
	"restservice/model"

	"github.com/gorilla/mux"
)

var requestCount int
var people []model.Person

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", UpdatePerson).Methods("PUT")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}

// GetPeople ...
func GetPeople(w http.ResponseWriter, r *http.Request) {
	requestCount++
	log.Println("Got request - GetPeople", requestCount)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

// GetPerson ...
func GetPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&model.Person{})
}

// CreatePerson ...
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var person model.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(person)

}

// UpdatePerson ...
func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var person model.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	var isPersonFound = false
	for index, item := range people {
		if item.ID == params["id"] {
			//people = append(people[:index], people[index+1:]...)
			people[index] = person
			isPersonFound = true
			break
		}
	}

	if isPersonFound == false {
		w.WriteHeader(http.StatusBadRequest)
		var er = model.ErrorResponse{StatusCode: 404, Message: "Invalid Data"}
		json.NewEncoder(w).Encode(er)
	} else {
		json.NewEncoder(w).Encode(person)
	}
}

// DeletePerson ...
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var isPersonFound = false
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			isPersonFound = true
			break
		}
	}

	if isPersonFound == false {
		w.WriteHeader(http.StatusBadRequest)
		var er = model.ErrorResponse{StatusCode: 404, Message: "Invalid Data"}
		json.NewEncoder(w).Encode(er)
	} else {
		var or = model.OkResponse{StatusCode: 200, Message: "Person Deleted"}
		json.NewEncoder(w).Encode(or)
	}
}

// init - This is called by go
func init() {
	people = append(people, model.Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &model.Address{City: "Irvine", State: "CA"}})
	people = append(people, model.Person{ID: "2", Firstname: "Jane", Lastname: "Doe", Address: &model.Address{City: "Newyork", State: "NY"}})
	people = append(people, model.Person{ID: "3", Firstname: "Jason", Lastname: "Doe", Address: &model.Address{City: "Seatle", State: "WA"}})
}
