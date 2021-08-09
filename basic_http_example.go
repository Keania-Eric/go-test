package main

import(
	"net/http"
	"log"
	"encoding/json"
	"fmt"
)


type testResponse struct {
	Message string
}

type testRequest struct {
	Name string `json:"name"`
}

func main() {

	port := 8080
	// we get the http handlefun
	http.HandleFunc("/test", responseHandlerFunc)
	log.Printf("Starting server on port %v", 8080)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func responseHandlerFunc(w http.ResponseWriter, r *http.Request) {

	var request testRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	response := testResponse{Message: "Hello "+ request.Name}
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}