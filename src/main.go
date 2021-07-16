package main

import (
	"./functions"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", functions.BasicResponse).Methods("GET")
	r.HandleFunc("/books/", functions.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", functions.GetBook).Methods("GET")
	r.HandleFunc("/books/", functions.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", functions.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", functions.DeleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe("localhost:8180", r))
}