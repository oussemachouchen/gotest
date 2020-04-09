package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Book struct
type Book struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Author *Author `json:"author"`
	ISBN   string  `json:"isbn"`
}

//Author struct
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var books []Book

//get Books with simple action
func getBooks(w http.ResponseWriter, s *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, s *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(s)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

//update a book using all informations
func updateBook(w http.ResponseWriter, s *http.Request) {

}

//delete book with id
func deleteBook(w http.ResponseWriter, s *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(s)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

//create book
func createBook(w http.ResponseWriter, s *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(s.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func handelRequests() {
	//showUsers()
	//init router
	router := mux.NewRouter().StrictSlash(true)
	//MOc data for books
	books = append(books, Book{ID: "1", ISBN: "jojoj", Name: "zozo", Author: &Author{FirstName: "JOKO", LastName: "fff"}})
	books = append(books, Book{ID: "2", ISBN: "1231", Name: "zzzz", Author: &Author{FirstName: "fffff", LastName: "fff"}})
	books = append(books, Book{ID: "3", ISBN: "875", Name: "ssss", Author: &Author{FirstName: "zzz", LastName: "sss"}})
	books = append(books, Book{ID: "4", ISBN: "63512", Name: "ffff", Author: &Author{FirstName: "JOKO", LastName: "fff"}})
	//router handler
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	//users from mysql db
	router.HandleFunc("/api/users", GetUsers).Methods("GET")
	router.HandleFunc("/api/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/api/users/{id}", DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/users/{name}", CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{id}/{name}", UpdateUser).Methods("PUT")
	//an other commit here
	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	InitialMigration()
	handelRequests()

}
