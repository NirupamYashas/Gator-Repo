package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Project struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
	Email      string `json:"email"`
	Link       string `json:"link"`
}

var projects []Project

func getProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

func addProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var project Project
	_ = json.NewDecoder(r.Body).Decode(&project)
	project.ID = strconv.Itoa(len(projects) + 1)
	projects = append(projects, project)
	json.NewEncoder(w).Encode(project)
}

func main() {
	r := mux.NewRouter()

	projects = append(projects, Project{ID: "1", Name: "Game Project", Department: "CISE", Email: "game@gmail.com", Link: "github.com/game"})
	projects = append(projects, Project{ID: "2", Name: "ML Project", Department: "CISE", Email: "ml@gmail.com", Link: "github.com/ml"})

	r.HandleFunc("/api/projects", getProjects).Methods("GET")
	r.HandleFunc("/api/projects", addProject).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
