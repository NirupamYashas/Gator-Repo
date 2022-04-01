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
