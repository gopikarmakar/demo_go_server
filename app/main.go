/**
 * @author [Gopi Karmakar]
 * @email [gopi.karmakar@monstar-lab.com]
 * @create date 2018-01-23 07:35:52
 * @modify date 2018-01-23 07:35:52
 * @desc [description]
 */

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/**
 *
 */
type Person struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

/**
 *
 */
type Address struct {
	City    string `json:"city,omitempty"`
	Country string `json:"country,omitempty"`
}

var people []Person

// Healthz Handler for monitoring app health
func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Chilling\n")
}

/**
 * Get all people from JSON DB
 */
func GetPeopleEndPoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

/**
 * Get one person from JSON DB whose ID matches.
 */
func GetPersonEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

/**
 * Create a person in JSON DB.
 */
func CreatePersonEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

/**
 * Deletes a person from JSON DB whose ID matches.
 */
func DeletePersonEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

// function to demonstrate a simple test in main_test.go
func Hello(name string) string {
	return fmt.Sprint("Hello ", name)
}

/**
 * Main
 */
func main() {

	router := mux.NewRouter()
	people = append(people, Person{ID: "1", FirstName: "Isaac", LastName: "Newton", Address: &Address{City: "London", Country: "England"}})
	people = append(people, Person{ID: "2", FirstName: "Albert", LastName: "Einstien", Address: &Address{City: "Berlin", Country: "Germany"}})
	people = append(people, Person{ID: "2", FirstName: "Thomas", LastName: "Edison", Address: &Address{City: "NYC", Country: "U.S.A"}})

	router.HandleFunc("/healthz", Healthz).Methods("GET")
	router.HandleFunc("/", GetPeopleEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndPoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndPoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":80", router))
}
