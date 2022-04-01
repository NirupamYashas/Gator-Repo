package controllers

import (
	"encoding/json"
	"net/http"
	"server/cors"
	"server/models"
	"server/utilities"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	cors.SetupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var users []models.User
	err := utilities.App.DB.Table("users").Find(&users).Error

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(users)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	cors.SetupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var user models.User
	vars := mux.Vars(r)
	id := vars["id"]
	user.ID = id
	err := utilities.App.DB.Table("users").Unscoped().Delete(user).Error

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	} else {
		json.NewEncoder(w).Encode("User deleted")
		return
	}
}

func SignupUser(w http.ResponseWriter, r *http.Request) {
	cors.SetupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var user models.User
	var reply models.LoginSignupReply
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		reply = models.LoginSignupReply{Message: "Error in decoding", Allow: false}
		json.NewEncoder(w).Encode(reply)
		return
	}

	err = utilities.App.DB.Table("users").Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error

	if err == gorm.ErrRecordNotFound {
		if user.Firstname == "" || user.Lastname == "" {
			reply = models.LoginSignupReply{Message: "Firstname and Lastname are required", Allow: false}
			json.NewEncoder(w).Encode(reply)
			return
		}

		user.ID = uuid.New().String()
		err = utilities.App.DB.Table("users").Save(&user).Error

		if err != nil {
			reply = models.LoginSignupReply{Message: "Error in saving user", Allow: false}
			json.NewEncoder(w).Encode(reply)
			return
		}

		reply = models.LoginSignupReply{Message: "User created successfully", Allow: true, Userdata: user}
		json.NewEncoder(w).Encode(reply)
		return
	}

	reply = models.LoginSignupReply{Message: "User already exists", Allow: false}
	json.NewEncoder(w).Encode(reply)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	cors.SetupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var user models.User
	var reply models.LoginSignupReply
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		reply = models.LoginSignupReply{Message: err.Error(), Allow: false}
		json.NewEncoder(w).Encode(reply)
		return
	}

	err = utilities.App.DB.Table("users").Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error

	if err == gorm.ErrRecordNotFound {
		reply = models.LoginSignupReply{Message: "Invalid Credentials", Allow: false}
		json.NewEncoder(w).Encode(reply)
		return
	} else if err != nil {
		reply = models.LoginSignupReply{Message: err.Error(), Allow: false}
		json.NewEncoder(w).Encode(reply)
		return
	}

	if user.ID != "" {
		reply = models.LoginSignupReply{Message: "Success", Allow: true, Userdata: user}
		json.NewEncoder(w).Encode(reply)
		return
	} else {
		reply = models.LoginSignupReply{Message: "Invalid Credentials", Allow: false}
		json.NewEncoder(w).Encode(reply)
		return
	}
}
