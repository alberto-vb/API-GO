package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

func newBook() {
	posturl := "http://localhost:8080/api/create_books"

	book := &Book{
		Author:    "alberto",
		Title:     "vegetta",
		Publisher: "anaya2",
	}
	json_data, err := json.Marshal(&book)

	if err != nil {
		log.Fatal(err)
	}

	res, err := http.Post(posturl, "application/json", bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	var response map[string]interface{}
	json.NewDecoder(res.Body).Decode(&response)

}

func main() {
	newBook()
}
