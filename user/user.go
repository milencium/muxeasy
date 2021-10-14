package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "database connection url link"

type User struct {
	gorm.Model
	FirstName string `json:firstname` 
	LastName string	 `json:lastname` 
	Email string	 `json:email`
}

func InitialMigration(){
	DB, err = gorm.Open(mysql.Open(DNS), &gormConfig)
	if err != nil {
		fmt.Println(err.Error())
		panic("cannot connect to db")
	}
	DB.AutoMigrate(&User{})

}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user []User
	DB.First(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user []User
	json.NewDecoder(r.Body).Decode(&user) //decodiranje, koristi se kod kreiranja usera
	DB.Create(&user)
	json.NewEncoder(w).Encode(user	)
}

func UpdateUser(w http.ResponseWriter, r *http){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user []User
	DB.First(&user, params["id"])
	json.newDecoder(r).Decode(&user)//decodiranje, koristi se kod updejtanje usera
	DB.save(&user)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user []User
	DB.Delete(&user,params["id"])
	json.newEncoder(w).Encode("The user deleted sucessfully")
}