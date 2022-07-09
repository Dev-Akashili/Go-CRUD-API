package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// Structure of Person
type Person struct {
	ID         string    `json:"id"`
	FullName   *FullName `json:"fullname"`
	Dob        string    `json:"dob"`
	Occupation string    `json:"occupation"`
}

//Create Person variable as a slice of Person
var Persons []Person

// Structure of FullNamw
type FullName struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Get all Persons
func getPersons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Persons)
}

// Get a specific Person
func getPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//Loop through Persons to find the id
	for _, item := range Persons {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

// Create a new Person
func createPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Person Person
	_ = json.NewDecoder(r.Body).Decode(&Person)
	Person.ID = strconv.Itoa(rand.Intn(1000)) //Generate a random ID for new Person
	Persons = append(Persons, Person)
	json.NewEncoder(w).Encode(Person)
}

// Update a Persons
func updatePerson(w http.ResponseWriter, r *http.Request) {
	// Combine methos from delete and update to achieve result
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Persons {
		if item.ID == params["id"] {
			Persons = append(Persons[:index], Persons[index+1:]...)

			// Update function merged
			var Person Person
			_ = json.NewDecoder(r.Body).Decode(&Person)
			Person.ID = params["id"]
			Persons = append(Persons, Person)
			json.NewEncoder(w).Encode(Person)
		}
	}
	json.NewEncoder(w).Encode(Persons)
}

// Delete a Person
func deletePersons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Persons {
		if item.ID == params["id"] {
			Persons = append(Persons[:index], Persons[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Persons)
}

func main() {
	//Initiate the Mux Router from imported packade "github.com/gorilla/mux"
	router := mux.NewRouter()

	//Create Mock Data to test API with Postman locally or integrate with DB

	// Create Route Handelers and endpoints
	router.HandleFunc("/api/Persons", getPersons).Methods("GET")
	router.HandleFunc("/api/Persons/{id}", getPerson).Methods("GET")
	router.HandleFunc("/api/Persons", createPerson).Methods("POST")
	router.HandleFunc("/api/Persons/{id}", updatePerson).Methods("PUT")
	router.HandleFunc("/api/Persons/{id}", deletePersons).Methods("DELETE")

	// Initiate server
	log.Fatal(http.ListenAndServe(":3000", router))
}
