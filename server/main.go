package main

import (
	"encoding/json"
	"log"
	"net/http"

	// "strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type App struct {
	db *gorm.DB
	r  *mux.Router
}

type Project struct {
	ID         string `gorm:"primaryKey" json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
	Email      string `json:"email"`
	Link       string `json:"link"`
}

var app App
var projects []Project
var project Project

func (a *App) start() {
	a.db.AutoMigrate(&Project{})

	project = Project{ID: uuid.New().String(), Name: "Game Project", Department: "CISE", Email: "game@gmail.com", Link: "github.com/game"}
	a.db.Create(&project)

	project = Project{ID: uuid.New().String(), Name: "ML Project", Department: "CISE", Email: "ml@gmail.com", Link: "github.com/ml"}
	a.db.Create(&project)

	a.r.HandleFunc("/api/projects", a.getProjects).Methods("GET")
	a.r.HandleFunc("/api/projects", a.addProject).Methods("POST")
	a.r.HandleFunc("/api/projects/{id}", a.updateProject).Methods("PUT")
	a.r.HandleFunc("/api/projects/{id}", a.deleteProject).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", a.r))
}

func (a *App) getProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := a.db.Find(&projects).Error

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(projects)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}
}

func (a *App) addProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&project)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	project.ID = uuid.New().String()
	err = a.db.Save(&project).Error

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func (a *App) updateProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&project)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	project.ID = mux.Vars(r)["id"]
	err = a.db.First(&project).Error

	if err == gorm.ErrRecordNotFound {
		w.WriteHeader(http.StatusNotFound)
	} else {
		err = a.db.Save(&project).Error

		if err != nil {
			json.NewEncoder(w).Encode(err.Error())
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}

func (a *App) deleteProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := a.db.Unscoped().Delete(&Project{ID: mux.Vars(r)["id"]}).Error

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func main() {
	db, err := gorm.Open(sqlite.Open("gator-repo.db"), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	app = App{
		db: db,
		r:  mux.NewRouter(),
	}

	app.start()
}
