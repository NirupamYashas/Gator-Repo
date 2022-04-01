package controllers

import (
	"encoding/json"
	"net/http"
	"server/cors"
	"server/models"
	"server/utilities"
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
