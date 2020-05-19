package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)


func GetUser(w http.ResponseWriter, r *http.Request) {
	db := InitDb()
	defer db.Close()

	var user Users
	vars := mux.Vars(r)
	id := vars["UserID"]

	db.First(&user, id)
	if user.Id != 0 {
		json.NewEncoder(w).Encode(user)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("User not found %v", id)))
	}


}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db := InitDb()
	defer db.Close()
	var users []Users
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	db := InitDb()
	defer db.Close()

	var user Users
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Invalid field given %v", err)))
		return
	}
	db.Create(&user)
	if user.Id == 0 {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte(fmt.Sprintf("User already exists %v", user)))
		return
	}
	json.NewEncoder(w).Encode(user)
}
