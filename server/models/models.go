package models

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type App struct {
	DB *gorm.DB
	R  *mux.Router
}

type Project struct {
	ID         string `gorm:"primaryKey" json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
	Email      string `json:"email"`
	Link       string `json:"link"`
}

type User struct {
	ID        string `gorm:"primaryKey" json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `gorm:"unique" json:"email"`
	Password  string `json:"password"`
	Isadmin   bool   `json:"isadmin"`
}

type LoginSignupReply struct {
	Userdata User   `json:"userdata"`
	Message  string `json:"message"`
	Allow    bool   `json:"allow"`
}
