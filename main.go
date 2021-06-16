package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mstzn/modanisa_backend/routes"
)

var listenPort = 3000

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/todos", routes.GetAllToDos).Methods("GET")
	router.HandleFunc("/todos", routes.AddNewTodo).Methods("POST")
	// ...

	fmt.Println(fmt.Sprintf("Listening on %d", listenPort))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s%d", ":", listenPort), router))
}

func main() {
	handleRequests()
}
