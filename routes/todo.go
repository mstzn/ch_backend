package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"
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

func GetAllToDos(c echo.Context) error {
	fmt.Println("All todos requested.")
	addHeaders(&c.Response().Writer)

	records := (database.GetDatabase()).GetAll()

	json.NewEncoder(c.Response()).Encode(records)

	return nil
}

func AddNewTodo(c echo.Context) error {
	fmt.Println("Add new todo requested.")

	addHeaders(&c.Response().Writer)

	reqBody, _ := ioutil.ReadAll(c.Request().Body)
	var todo models.ToDo
	err := json.Unmarshal(reqBody, &todo)
	if err != nil {
		json.NewEncoder(c.Response()).Encode(errors.GetInvalidRequest("Can not unmarshal request body"))
		return err
	}

	if todo.Title == "" {
		json.NewEncoder(c.Response()).Encode(errors.GetInvalidRequest("Title must provided!"))
		return nil
	}

	if todo.Id == "" {
		todo.Id = uuid.Must(uuid.NewV4(), err).String()
		todo.Done = false
	}

	(database.GetDatabase()).Insert(todo)

	err2 := json.NewEncoder(c.Response()).Encode(todo)
	if err2 != nil {
		log.Println("Can not marshal response body")
		return err2
	}

	return nil
}
