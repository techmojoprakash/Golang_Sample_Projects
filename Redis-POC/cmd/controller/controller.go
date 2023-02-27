package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"redispoc/cmd/cache"
	"redispoc/cmd/model"
	"time"

	"github.com/gorilla/mux"
)

var RCache *cache.RedisCache

const (
	PORT = ":8080"
)

func InitializeRouter() {
	routerX := mux.NewRouter()

	routerX.HandleFunc("/getstudent/{id}", GetStudent).Methods("GET")
	routerX.HandleFunc("/setstudent", SetStudent).Methods("POST")
	routerX.HandleFunc("/delstudent/{id}", DeleteStudent).Methods("DELETE")

	fmt.Println("Starting Server at ", PORT)
	log.Fatal(http.ListenAndServe(PORT, routerX))
}

func InsertDummyData() {
	s1 := model.Student{
		Id:         "1",
		Name:       "Prakash",
		Gpa:        3.9,
		IsEligible: false,
	}

	RCache.Set(s1.Id, &s1)
	fmt.Println("Basic Data inserted with ID", s1.Id)
	firstVal, err := RCache.Get(s1.Id)
	if err != nil {
		fmt.Println("Unable to retrive data from Redis ", err)
	}
	fmt.Println("Basic Data retrived ", firstVal)
}
func InitializeRedis() {
	RCache = cache.NewRedisCache("localhost:6379", 0, 5*time.Minute)
	println("Redis-Cache created...! ", RCache)
	InsertDummyData()
	println("Dummy Data Inserted...! ", RCache)
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetStudent .... invoked")
	params := mux.Vars(r)
	studentID := params["id"]
	stuData, getErr := RCache.Get(studentID)
	if getErr != nil {
		fmt.Println("Key not found in cache")
	}
	fmt.Println("GetStudent Func Sending => ", stuData)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stuData)

}

func SetStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SetStudent .... invoked")
	w.Header().Set("Content-Type", "application/json")
	var stuObj model.Student
	_ = json.NewDecoder(r.Body).Decode(&stuObj) // must pass pointer
	err := RCache.Set(stuObj.Id, &stuObj)
	if err != nil {
		fmt.Println("Unable to Set : Error getErr")
	}

}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteStudent .... invoked")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	studentID := params["id"]

	err := RCache.Del(studentID)
	if err != nil {
		fmt.Println("Unable to Delete Key : Error DeleteErr", studentID)
	}

}
