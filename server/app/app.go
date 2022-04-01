package app

import (
	"log"
	"net/http"

	"server/models"
	"server/utilities"

	"server/controllers"
)

func Start() {
	utilities.App.DB.AutoMigrate(&models.User{})
	utilities.App.DB.AutoMigrate(&models.Project{})

	// var user User = User{ID: uuid.New().String(), Firstname: "Jagan", Lastname: "Dwarampudi", Email: "scientistjagan2000@gmail.com", Password: "password"}
	// a.DB.Table("users").Create(&user)

	// var user User
	// a.DB.Table("users").Where("email = ?", "nirupamyashas@gmail.com").First(&user)
	// a.DB.Table("users").Delete(&user)

	// var project Project = Project{ID: uuid.New().String(), Name: "Game Project", Department: "CISE", Email: "game@gmail.com", Link: "github.com/game"}
	// a.db.Table("projects").Create(&project)

	// project = Project{ID: uuid.New().String(), Name: "ML Project", Department: "CISE", Email: "ml@gmail.com", Link: "github.com/ml"}
	// a.db.Table("projects").Create(&project)

	// var project Project
	// a.DB.Table("projects").Where("id = ?", "b69cb878-148c-462e-9c3e-f66d04033570").First(&project)
	// a.DB.Table("projects").Delete(&project)

	utilities.App.R.HandleFunc("/api/users/signup", controllers.SignupUser).Methods("POST", "OPTIONS")
	utilities.App.R.HandleFunc("/api/users/login", controllers.LoginUser).Methods("POST", "OPTIONS")
	utilities.App.R.HandleFunc("/api/users", controllers.GetUsers).Methods("GET", "OPTIONS")
	utilities.App.R.HandleFunc("/api/users/{id}", controllers.DeleteUser).Methods("DELETE", "OPTIONS")

	utilities.App.R.HandleFunc("/api/projects", controllers.GetProjects).Methods("GET", "OPTIONS")
	utilities.App.R.HandleFunc("/api/projects", controllers.AddProject).Methods("POST", "OPTIONS")
	utilities.App.R.HandleFunc("/api/projects/{id}", controllers.UpdateProject).Methods("PUT", "OPTIONS")
	utilities.App.R.HandleFunc("/api/projects/{id}", controllers.DeleteProject).Methods("DELETE", "OPTIONS")
	utilities.App.R.HandleFunc("/api/projects/departments/{department}", controllers.GetProjectsByDepartment).Methods("GET", "OPTIONS")
	utilities.App.R.HandleFunc("/api/projects/search/{search_phrase}", controllers.GetProjectsBySearch).Methods("GET", "OPTIONS")

	log.Fatal(http.ListenAndServe(":8080", utilities.App.R))
}
