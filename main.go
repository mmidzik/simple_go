/*
translate CRUD app: https://medium.com/@etiennerouzeaud/how-to-create-a-basic-restful-api-in-go-c8e032ba3181
goal: use raw sqlite and http routing (no gin, gorm, other helpers etc)
*/


package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Users struct {
	Id        int    `gorm:"AUTO_INCREMENT" form:"id" json:"id" `
	Firstname string `gorm:"not null" form:"firstname" json:"firstname"`
	Lastname  string `gorm:"not null" form:"lastname" json:"lastname"`
}

func InitDb() *gorm.DB {
	// Opening file
	db, err := gorm.Open("sqlite3", "./data.db")
	// Display SQL queries
	db.LogMode(true)

	// Error
	if err != nil {
		panic(err)
	}
	// Creating the table
	if !db.HasTable(&Users{}) {
		db.CreateTable(&Users{}).AddUniqueIndex("idx_user_name", "firstname", "lastname")
	}

	return db
}


func main() {
	InitDb()
	log.Println("Server started on: http://localhost:8080")
	router := mux.NewRouter()
	router.HandleFunc("/user", PostUser).Methods("POST")
	router.HandleFunc("/user/{UserID}", GetUser).Methods("GET")
	router.HandleFunc("/users", GetAllUsers).Methods("GET")
	http.Handle("/", router)
	http.ListenAndServe("localhost:8080", nil)
}


