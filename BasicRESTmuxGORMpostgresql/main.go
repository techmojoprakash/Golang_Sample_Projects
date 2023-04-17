package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// -------------------
type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

var DB *gorm.DB
var err error

const DNS = "postgres://postgres:root@localhost:5432/godb"

// username : postgres
// pw : root
// db name : godb

func InitialMigration() {
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	} else {
		fmt.Println("Connected to DB successfully...!")
	}

	DB.AutoMigrate(&User{})
	/*
		This Auto Migration feature will automatically migrate your schema.
		It will automatically create the table based on your model.
		We don’t need to create the table manually.
	*/

}

// GET
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	db_err := DB.Find(&users).Error
	// checking if we got any error from db query
	if db_err == gorm.ErrRecordNotFound {
		// now sending error to client
		http.Error(w, db_err.Error(), 404)
		return
	}
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}

// GET
func GetUserbyId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	db_err := DB.First(&user, params["id"]).Error
	// checking if we got any error from db query
	if db_err == gorm.ErrRecordNotFound {
		// now sending error to client
		http.Error(w, db_err.Error(), 404)
		return
	}
	err_server := json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err_server.Error(), 500)
		return
	}
}

// POST
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	cli_err := json.NewDecoder(r.Body).Decode(&user)
	if cli_err != nil {
		http.Error(w, cli_err.Error(), 500)
		return
	}
	// need to handle properly below
	result := DB.Create(&user)
	if result != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// err := json.NewEncoder(w).Encode(user)
	// if err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }
	// w.WriteHeader(http.StatusCreated) // not working
	// resp := make(map[string]string)
	// resp["message"] = "Status Created"
	// _ := json.NewEncoder(w).Encode(resp).Error()
	// // if err != nil {
	// // 	http.Error(w, err1.Error(), 500)
	// // 	return
	// // }
	// // w.Write(resp)
	// return

}

// PUT  //doubt
func UpdateUserbyId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)         // getting id
	var user User                 // to save user from db we are creating user struct
	DB.First(&user, params["id"]) // fetching value by id
	cli_err := json.NewDecoder(r.Body).Decode(&user)
	if cli_err != nil {
		http.Error(w, cli_err.Error(), 500)
		return
	}
	DB.Save(&user)
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// DELETE
// need to write func by using flag inorder to keep old data

func DeleteUserbyId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User

	// // get code from this file
	// exist_err := DB.First(&user, params["id"]).Error
	// // checking if we got any error from db query
	// if exist_err == gorm.ErrRecordNotFound {
	// 	// now sending error to client
	// 	http.Error(w, exist_err.Error(), 404)
	// 	return
	// }

	db_err := DB.Delete(&user, params["id"]).Error
	// checking if we got any error from db query
	if db_err.Error != nil {
		// now sending error to client
		http.Error(w, db_err.Error(), 404)
		return
	}
	err := json.NewEncoder(w).Encode("User Deleted")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// --------------

func initializeRouter() {
	var r = mux.NewRouter() // creating new router from mux pkg

	//routes
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", GetUserbyId).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUserbyId).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUserbyId).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))

}

func main() {
	fmt.Println("App started at Port no : 8080 ....")
	InitialMigration()
	initializeRouter()

}
