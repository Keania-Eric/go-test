package main

import (
	"encoding/json"
	"log"
	"net/http"
)

//This is just an API To carryout CRUD on Toys

type Owner struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

// Definition of toys
type Toy struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type ToyController struct {
	Repository []Toy
}

// Initialize our ToyController
func NewToyController() *ToyController {
	return &ToyController{
		Repository: []Toy{
			Toy{Name: "Cows", Color: "red", ID: "123456"},
		},
	}
}

//Responsible for retrieving all toys
func (h *ToyController) get(w http.ResponseWriter, r *http.Request) {
	toys := h.Repository
	encoder := json.NewEncoder(w)
	encoder.Encode(toys)
}

func (h *ToyController) add(w http.ResponseWriter, r *http.Request) {
	//initialize the toy
	var toy Toy
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&toy)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	toys := append(h.Repository, toy)
	h.Repository = toys
	encoder := json.NewEncoder(w)
	encoder.Encode(toys)
}

// Delete toys by name
func (h *ToyController) delete(w http.ResponseWriter, r *http.Request) {
	// get id from the url query
	id := r.URL.Query().Get("id")

	//we loop through the controller repository and remove the item with ID
	for index, value := range h.Repository {
		if value.ID == id {
			//TODO remove the item from the array
			h.Repository = append(h.Repository[:index], h.Repository[index+1:]...)
			break
		}
	}

	// return whats left of the h.Repository
	encoder := json.NewEncoder(w)
	encoder.Encode(h.Repository)
}

func (h *ToyController) update(w http.ResponseWriter, r *http.Request) {
	// get id from the url query
	id := r.URL.Query().Get("id")

	//retrieve the request body
	var toy Toy
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&toy)

	if err != nil {
		http.Error(w, "An error occured", http.StatusBadRequest)
	}

	//we loop through the controller repository and remove the item with ID
	for index, value := range h.Repository {
		if value.ID == id {
			//TODO remove the item from the array
			h.Repository = append(h.Repository[:index], h.Repository[index+1:]...)
			h.Repository = append(h.Repository, toy)
			break
		}
	}

	// return whats left of the h.Repository
	encoder := json.NewEncoder(w)
	encoder.Encode(h.Repository)
}

func main() {

	//we declare the server port
	port := 9000

	toyController := NewToyController()

	http.HandleFunc("/toys", toyController.get)
	http.HandleFunc("/toys/add", toyController.add)
	http.HandleFunc("/toys/delete", toyController.delete)
	http.HandleFunc("/toys/update", toyController.update)

	// We tell what port the server is listening
	log.Printf("Server is started and is listening on port %v", port)

	// Start the server and listen on a port
	log.Fatal(http.ListenAndServe(":9000", nil))

}
