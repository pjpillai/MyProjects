package service

import (
	"strconv"

	"github.com/pjpillai/MyProjects/go_projects/02_web/01_restservice/model"
)

var people []model.Person

// GetPeople ...
func GetPeople() *[]model.Person {
	return &people
}

// GetPerson ...
func GetPerson(id string) *model.Person {

	var person = model.Person{}
	for _, item := range people {
		if item.ID == id {
			person = item
		}
	}
	return &person
}

// CreatePerson ...
func CreatePerson(person *model.Person) bool {
	personCount := len(people)
	person.ID = strconv.Itoa(personCount + 1)
	people = append(people, *person)
	return true
}

// UpdatePerson ...
func UpdatePerson(person *model.Person) bool {

	for index, item := range people {
		if item.ID == person.ID {
			people[index] = *person
			return true
		}
	}
	return false
}

// DeletePerson ...
func DeletePerson(person *model.Person) bool {

	for index, item := range people {
		if item.ID == person.ID {
			people = append(people[:index], people[index+1:]...)
			return true
		}
	}
	return false
}

// init ...
func init() {
	people = append(people, model.Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &model.Address{City: "Tustin", State: "CA"}})
	people = append(people, model.Person{ID: "2", Firstname: "Jane", Lastname: "Doe", Address: &model.Address{City: "New York", State: "NY"}})
	people = append(people, model.Person{ID: "3", Firstname: "Jason", Lastname: "Doe", Address: &model.Address{City: "Seattle", State: "WA"}})
}
