package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//User struct
type User struct {
	gorm.Model
	ID   string `json:"id"`
	Name string `json:"name"`
}

var users []User
var db *gorm.DB
var err error

//InitialMigration initialize working for db
func InitialMigration() {
	db, err := gorm.Open("mysql", "root:oussema@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	db.AutoMigrate(&users)
}

//GetUsers function
func GetUsers(w http.ResponseWriter, s *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("start connection")
	db, err := gorm.Open("mysql", "root:oussema@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	fmt.Println("connected successfully")

	db.Find(&users)

	json.NewEncoder(w).Encode(users)
}

//GetUser gets a user by id
func GetUser(w http.ResponseWriter, s *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("start connection")
	db, err := gorm.Open("mysql", "root:oussema@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	fmt.Println("connected successfully")
	params := mux.Vars(s)
	//execute the query
	var user User
	db.Where("id = ?", params["id"]).Find(&user)

	json.NewEncoder(w).Encode(user)
}

//CreateUser create a new user
func CreateUser(w http.ResponseWriter, s *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("start connection")
	db, err := gorm.Open("mysql", "root:oussema@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	fmt.Println("connected successfully")
	params := mux.Vars(s)
	//new user:
	var user User
	user.ID = strconv.Itoa(rand.Intn(10000))
	user.Name = params["name"]
	//execute the query
	db.Create(&user)

	json.NewEncoder(w).Encode("new user created successfully")
}

//DeleteUser gets a user by id
func DeleteUser(w http.ResponseWriter, s *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("start connection")
	db, err := gorm.Open("mysql", "root:oussema@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	fmt.Println("connected successfully")
	params := mux.Vars(s)
	//execute the query
	var user User
	db.Where("id = ?", params["id"]).Find(&user)
	db.Delete(&user)
	json.NewEncoder(w).Encode("User is deleted successfully")
}

//UpdateUser gets a user by id
func UpdateUser(w http.ResponseWriter, s *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("start connection")
	db, err := gorm.Open("mysql", "root:oussema@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	fmt.Println("connected successfully")
	params := mux.Vars(s)
	//execute the query
	var user User
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println("error converting param")
	}
	db.Where("id = ?", id).Find(&user)
	user.Name = params["name"]
	user.ID = params["id"]
	db.Save(&user)
	//new comment added
	json.NewEncoder(w).Encode("User is updated successfully")
}
