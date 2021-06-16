package models

type ToDo struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Done 	bool `json:"done"`
}

