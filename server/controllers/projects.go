package controllers

import (
	"encoding/json"
	"net/http"
	"server/cors"
	"server/models"
	"server/utilities"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetProjects(w http.ResponseWriter, r *http.Request) {
	cors.SetupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var projects []models.Project
	err := utilities.App.DB.Table("projects").Find(&projects).Error

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(projects)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
}

func AddProject(w http.ResponseWriter, r *http.Request) {
	cors.SetupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var project models.Project
	err := json.NewDecoder(r.Body).Decode(&project)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	project.ID = uuid.New().String()
	err = utilities.App.DB.Table("projects").Save(&project).Error

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	} else {
		w.WriteHeader(http.StatusCreated)
		return
	}
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	cors.SetupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var project models.Project
	err := json.NewDecoder(r.Body).Decode(&project)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	project.ID = mux.Vars(r)["id"]
	err = utilities.App.DB.Table("projects").First(&project).Error

	if err == gorm.ErrRecordNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	} else {
		err = utilities.App.DB.Table("projects").Save(&project).Error

		if err != nil {
			json.NewEncoder(w).Encode(err.Error())
			return
		} else {
			w.WriteHeader(http.StatusOK)
			return
		}
	}
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	cors.SetupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var project models.Project
	vars := mux.Vars(r)
	id := vars["id"]
	project.ID = id
	err := utilities.App.DB.Table("projects").Unscoped().Delete(project).Error

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	} else {
		json.NewEncoder(w).Encode("Project deleted successfully")
		return
	}
}

func GetProjectsByDepartment(w http.ResponseWriter, r *http.Request) {
	cors.SetupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var projects []models.Project
	err := utilities.App.DB.Table("projects").Find(&projects, "department = ?", mux.Vars(r)["department"]).Error

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(projects)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}
}

func GetProjectsBySearch(w http.ResponseWriter, r *http.Request) {
	cors.SetupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	res := strings.Split(mux.Vars(r)["search_phrase"], " ")
	tx := utilities.App.DB.Table("projects")

	for _, element := range res {
		search_term := "%" + element + "%"
		tx = tx.Where("name LIKE ? OR email LIKE ?", search_term, search_term)
	}

	var projects []models.Project
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
