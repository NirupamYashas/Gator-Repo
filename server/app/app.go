package app

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
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
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

func (a *App) Start() {
	a.DB.AutoMigrate(&User{})
	a.DB.AutoMigrate(&Project{})

	// var user User = User{ID: uuid.New().String(), Firstname: "Jagan", Lastname: "Dwarampudi", Email: "scientistjagan2000@gmail.com", Password: "password"}
	// a.DB.Table("users").Create(&user)

	// var user User
	// a.DB.Table("users").Where("email = ?", "nirupamyashas@gmail.com").First(&user)
	// a.DB.Table("users").Delete(&user)

	// var project Project = Project{ID: uuid.New().String(), Name: "Game Project", Department: "CISE", Email: "game@gmail.com", Link: "github.com/game"}
	// a.db.Table("projects").Create(&project)

	// project = Project{ID: uuid.New().String(), Name: "ML Project", Department: "CISE", Email: "ml@gmail.com", Link: "github.com/ml"}
	// a.db.Table("projects").Create(&project)

	a.R.HandleFunc("/api/users/signup", a.signupUser).Methods("POST", "OPTIONS")
	a.R.HandleFunc("/api/users/login", a.loginUser).Methods("POST", "OPTIONS")

	a.R.HandleFunc("/api/projects", a.getProjects).Methods("GET", "OPTIONS")
	a.R.HandleFunc("/api/projects", a.addProject).Methods("POST", "OPTIONS")
	a.R.HandleFunc("/api/projects/{id}", a.updateProject).Methods("PUT", "OPTIONS")
	a.R.HandleFunc("/api/projects/{id}", a.deleteProject).Methods("DELETE", "OPTIONS")
	a.R.HandleFunc("/api/projects/departments/{department}", a.getProjectsByDepartment).Methods("GET", "OPTIONS")
	a.R.HandleFunc("/api/projects/search/{search_phrase}", a.getProjectsBySearch).Methods("GET", "OPTIONS")

	log.Fatal(http.ListenAndServe(":8080", a.R))
}

func (a *App) signupUser(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = a.DB.Table("users").Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error

	if err == gorm.ErrRecordNotFound {
		if user.Firstname == "" || user.Lastname == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		user.ID = uuid.New().String()
		err = a.DB.Table("users").Save(&user).Error

		if err != gorm.ErrRecordNotFound {
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

func (a *App) loginUser(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = a.DB.Table("users").Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error

	if err == gorm.ErrRecordNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	if user.ID != "" {
		w.WriteHeader((http.StatusOK))
		return
	} else {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func (a *App) getProjects(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var projects []Project
	err := a.DB.Table("projects").Find(&projects).Error

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
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var project Project
	err := json.NewDecoder(r.Body).Decode(&project)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	project.ID = uuid.New().String()
	err = a.DB.Table("projects").Save(&project).Error

	if err != gorm.ErrRecordNotFound {
		json.NewEncoder(w).Encode(err.Error())
		return
	} else {
		w.WriteHeader(http.StatusCreated)
		return
	}
}

func (a *App) updateProject(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var project Project
	err := json.NewDecoder(r.Body).Decode(&project)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	project.ID = mux.Vars(r)["id"]
	err = a.DB.Table("projects").First(&project).Error

	if err == gorm.ErrRecordNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	} else {
		err = a.DB.Table("projects").Save(&project).Error

		if err != nil {
			json.NewEncoder(w).Encode(err.Error())
			return
		} else {
			w.WriteHeader(http.StatusOK)
			return
		}
	}
}

func (a *App) deleteProject(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err := a.DB.Table("projects").Unscoped().Delete(&Project{ID: mux.Vars(r)["id"]}).Error

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	} else {
		w.WriteHeader(http.StatusOK)
		return
	}
}

func (a *App) getProjectsByDepartment(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var projects []Project
	err := a.DB.Table("projects").Find(&projects, "department = ?", mux.Vars(r)["department"]).Error

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
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	res := strings.Split(mux.Vars(r)["search_phrase"], " ")
	tx := a.DB.Table("projects")

	for _, element := range res {
		search_term := "%" + element + "%"
		tx = tx.Where("name LIKE ? OR email LIKE ?", search_term, search_term)
	}

	var projects []Project
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
