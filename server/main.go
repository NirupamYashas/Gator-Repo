package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	// "strconv"

	"github.com/google/uuid"
	"github.com/gorilla/handlers"
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

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Headers:", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
		return
	})
}

func (a *App) start() {
	a.db.AutoMigrate(&Project{})

	// project = Project{ID: uuid.New().String(), Name: "Game Project", Department: "CISE", Email: "game@gmail.com", Link: "github.com/game"}
	// a.db.Create(&project)

	// project = Project{ID: uuid.New().String(), Name: "ML Project", Department: "CISE", Email: "ml@gmail.com", Link: "github.com/ml"}
	// a.db.Create(&project)

	a.r.HandleFunc("/api/projects", a.getProjects).Methods("GET")
	a.r.HandleFunc("/api/projects", a.addProject).Methods("POST")
	a.r.HandleFunc("/api/projects/{id}", a.updateProject).Methods("PUT")
	a.r.HandleFunc("/api/projects/{id}", a.deleteProject).Methods("DELETE")
	a.r.HandleFunc("/api/projects/departments/{department}", a.getProjectsByDepartment).Methods("GET")
	a.r.HandleFunc("/api/projects/search/{search_phrase}", a.getProjectsBySearch).Methods("GET")

	// log.Fatal(http.ListenAndServe(":8080", a.r))
	log.Fatal(http.ListenAndServe(
		":8080",
		handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"*"}),
			handlers.AllowedOrigins([]string{"*"}))(a.r)))
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

	fmt.Println("hi")
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

func (a *App) getProjectsByDepartment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := a.db.Find(&projects, "department = ?", mux.Vars(r)["department"]).Error

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(projects)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}
}

func (a *App) getProjectsBySearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res := strings.Split(mux.Vars(r)["search_phrase"], " ")
	tx := a.db

	for _, element := range res {
		search_term := "%" + element + "%"
		tx = tx.Where("name LIKE ? OR email LIKE ?", search_term, search_term)
	}

	err := tx.Find(&projects).Error

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(projects)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
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
