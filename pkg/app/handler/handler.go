package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/niconosenzo/devopsapi/pkg/app/model"
)

//Create user through POST request
func CreateUser(w http.ResponseWriter, r *http.Request, Users []model.User) []model.User {
	user := model.User{}

	//get the body request
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
	}
	json.Unmarshal(reqBody, &user)

	//check if user already exists
	for _, u := range Users {
		if u.ID == user.ID {
			respondError(w, http.StatusBadRequest, "User ID already exists")
		}
	}

	// update  global Users array
	Users = append(Users, user)
	respondJSON(w, http.StatusOK, user)
	return Users
}

//GetUser by ID
func GetUser(w http.ResponseWriter, r *http.Request, Users []model.User) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, user := range Users {
		if user.ID == id {
			respondJSON(w, http.StatusOK, user)
			return
		}
	}
	respondError(w, http.StatusNotFound, "User not found")
}

//DeleteUser by ID
func DeleteUser(w http.ResponseWriter, r *http.Request, Users []model.User) []model.User {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, user := range Users {
		if user.ID == id {
			Users = append(Users[:index], Users[index+1:]...)
		}
	}
	return Users
}

//GetAllUsers
func GetAllUsers(w http.ResponseWriter, r *http.Request, Users []model.User) {
	//Implemented for testing only
	respondJSON(w, http.StatusOK, Users)
}
