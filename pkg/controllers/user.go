package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kapitanmx/ATM/pkg/models"
	"github.com/kapitanmx/ATM/pkg/utils"
)

var NewUser models.User

func GetUser(w http.ResponseWriter, r *http.Request) {
	newUsers := models.GetAllUsers()
	res, _ := json.Marshal(newUsers)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["userId"]
	userDetails, _ := models.GetUserById(ID)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)
	b := CreateUser.CreateUser()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["userId"]
	user := models.DeleteUser(ID)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser = &models.User{}
	utils.ParseBody(r, updateUser)
	vars := mux.Vars(r)
	ID := vars["userId"]
	userDetails, db := models.GetUserById(ID)
	if updateUser.Name != "" {
		userDetails.Name = updateUser.Name
	}
	if updateUser.LastName != "" {
		userDetails.LastName = updateUser.LastName
	}
	if updateUser.Email != "" {
		userDetails.Email = updateUser.Email
	}
	if updateUser.Pin != "" {
		userDetails.Pin = updateUser.Pin
	}
	db.Save(&userDetails)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
