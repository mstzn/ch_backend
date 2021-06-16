package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mstzn/modanisa_backend/database"
	"github.com/mstzn/modanisa_backend/errors"
	"github.com/mstzn/modanisa_backend/models"
	uuid "github.com/satori/go.uuid"
)

func addHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func GetAllToDos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("All todos requested.")
	addHeaders(&w)

	records := (database.GetDatabase()).GetAll()

	json.NewEncoder(w).Encode(records)
}

func AddNewTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add new todo requested.")

	addHeaders(&w)

	reqBody, _ := ioutil.ReadAll(r.Body)
	var todo models.ToDo
	err := json.Unmarshal(reqBody, &todo)
	if err != nil {
		json.NewEncoder(w).Encode(errors.GetInvalidRequest("Can not unmarshal request body"))
		return
	}

	if todo.Title == "" {
		json.NewEncoder(w).Encode(errors.GetInvalidRequest("Title must provided!"))
		return
	}

	if todo.Id == "" {
		todo.Id = uuid.Must(uuid.NewV4(), err).String()
		todo.Done = false
	}

	(database.GetDatabase()).Insert(todo)

	err2 := json.NewEncoder(w).Encode(todo)
	if err2 != nil {
		log.Println("Can not marshal response body")
		return
	}
}
