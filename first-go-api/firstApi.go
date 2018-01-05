package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Person struct {
	ID string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname string `json:"lastname,omitempty"`
	Address *Address `json:"address,omitempty"`
}

type Address struct {
	City string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people[]Person

func GetPersonEndpoint(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	for item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
	json.NewEncoder(writer).Encode(&Person{})
}

func main() {
	people = append(people, Person{ID: "1", Firstname: "Alexes", Lastname: "Presswood", Address: &Address{City: "Kansas City", State: "MO"}})
	people = append(people, Person{ID: "2", Firstname: "Test", Lastname: "User"})
}
