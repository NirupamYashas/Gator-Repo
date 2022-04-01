package app

import (
	"log"
	"net/http"

	"server/models"
	"server/utilities"

	"server/controllers"
)

// type App struct {
// 	DB *gorm.DB
// 	R  *mux.Router
// }

// func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "*")
// 	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// 	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
// }

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

// func (a *App) getUsers(w http.ResponseWriter, r *http.Request) {
// 	cors.SetupCorsResponse(&w, r)
// 	if (*r).Method == "OPTIONS" {
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")

// 	var users []models.User
// 	err := a.DB.Table("users").Find(&users).Error

// 	if err != nil {
// 		json.NewEncoder(w).Encode(err.Error())
// 		return
// 	}

// 	err = json.NewEncoder(w).Encode(users)

// 	if err != nil {
// 		json.NewEncoder(w).Encode(err.Error())
// 		return
// 	}
// }

// func (a *App) deleteUser(w http.ResponseWriter, r *http.Request) {
// 	cors.SetupCorsResponse(&w, r)
// 	if (*r).Method == "OPTIONS" {
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")

// 	var user models.User
// 	vars := mux.Vars(r)
// 	id := vars["id"]
// 	user.ID = id
// 	err := a.DB.Table("users").Unscoped().Delete(user).Error

// 	if err != nil {
// 		json.NewEncoder(w).Encode(err.Error())
// 		return
// 	} else {
// 		json.NewEncoder(w).Encode("User deleted")
// 		return
// 	}
// }

// func (a *App) signupUser(w http.ResponseWriter, r *http.Request) {
// 	cors.SetupCorsResponse(&w, r)
// 	if (*r).Method == "OPTIONS" {
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")

// 	var user models.User
// 	var reply models.LoginSignupReply
// 	err := json.NewDecoder(r.Body).Decode(&user)

// 	if err != nil {
// 		reply = models.LoginSignupReply{Message: "Error in decoding", Allow: false}
// 		json.NewEncoder(w).Encode(reply)
// 		return
// 	}

// 	err = a.DB.Table("users").Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error

// 	if err == gorm.ErrRecordNotFound {
// 		if user.Firstname == "" || user.Lastname == "" {
// 			reply = models.LoginSignupReply{Message: "Firstname and Lastname are required", Allow: false}
// 			json.NewEncoder(w).Encode(reply)
// 			return
// 		}

// 		user.ID = uuid.New().String()
// 		err = a.DB.Table("users").Save(&user).Error

// 		if err != nil {
// 			reply = models.LoginSignupReply{Message: "Error in saving user", Allow: false}
// 			json.NewEncoder(w).Encode(reply)
// 			return
// 		}

// 		reply = models.LoginSignupReply{Message: "User created successfully", Allow: true, Userdata: user}
// 		json.NewEncoder(w).Encode(reply)
// 		return
// 	}

// 	reply = models.LoginSignupReply{Message: "User already exists", Allow: false}
// 	json.NewEncoder(w).Encode(reply)
// }

// func (a *App) loginUser(w http.ResponseWriter, r *http.Request) {
// 	cors.SetupCorsResponse(&w, r)
// 	if (*r).Method == "OPTIONS" {
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")

// 	var user models.User
// 	var reply models.LoginSignupReply
// 	err := json.NewDecoder(r.Body).Decode(&user)

// 	if err != nil {
// 		reply = models.LoginSignupReply{Message: err.Error(), Allow: false}
// 		json.NewEncoder(w).Encode(reply)
// 		return
// 	}

// 	err = a.DB.Table("users").Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error

// 	if err == gorm.ErrRecordNotFound {
// 		reply = models.LoginSignupReply{Message: "Invalid Credentials", Allow: false}
// 		json.NewEncoder(w).Encode(reply)
// 		return
// 	} else if err != nil {
// 		reply = models.LoginSignupReply{Message: err.Error(), Allow: false}
// 		json.NewEncoder(w).Encode(reply)
// 		return
// 	}

// 	if user.ID != "" {
// 		reply = models.LoginSignupReply{Message: "Success", Allow: true, Userdata: user}
// 		json.NewEncoder(w).Encode(reply)
// 		return
// 	} else {
// 		reply = models.LoginSignupReply{Message: "Invalid Credentials", Allow: false}
// 		json.NewEncoder(w).Encode(reply)
// 		return
// 	}
// }

// func getProjects(w http.ResponseWriter, r *http.Request) {
// 	cors.SetupCorsResponse(&w, r)
// 	if (*r).Method == "OPTIONS" {
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")

// 	var projects []models.Project
// 	err := utilities.App.DB.Table("projects").Find(&projects).Error

// 	if err != nil {
// 		json.NewEncoder(w).Encode(err.Error())
// 		return
// 	}

// 	err = json.NewEncoder(w).Encode(projects)

// 	if err != nil {
// 		json.NewEncoder(w).Encode(err.Error())
// 		return
// 	}
// }

// func addProject(w http.ResponseWriter, r *http.Request) {
// 	cors.SetupCorsResponse(&w, r)
// 	if (*r).Method == "OPTIONS" {
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")

// 	var project models.Project
// 	err := json.NewDecoder(r.Body).Decode(&project)

// 	if err != nil {
// 		json.NewEncoder(w).Encode(err.Error())
// 		return
// 	}

// 	project.ID = uuid.New().String()
// 	err = utilities.App.DB.Table("projects").Save(&project).Error

// 	if err != nil {
// 		json.NewEncoder(w).Encode(err.Error())
// 		return
// 	} else {
// 		w.WriteHeader(http.StatusCreated)
// 		return
// 	}
// }

// func updateProject(w http.ResponseWriter, r *http.Request) {
// 	cors.SetupCorsResponse(&w, r)
// 	if (*r).Method == "OPTIONS" {
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")

// 	var project models.Project
// 	err := json.NewDecoder(r.Body).Decode(&project)

// 	if err != nil {
// 		json.NewEncoder(w).Encode(err.Error())
// 		return
// 	}

// 	project.ID = mux.Vars(r)["id"]
// 	err = utilities.App.DB.Table("projects").First(&project).Error

// 	if err == gorm.ErrRecordNotFound {
// 		w.WriteHeader(http.StatusNotFound)
// 		return
// 	} else {
// 		err = utilities.App.DB.Table("projects").Save(&project).Error

// 		if err != nil {
// 			json.NewEncoder(w).Encode(err.Error())
// 			return
// 		} else {
// 			w.WriteHeader(http.StatusOK)
// 			return
// 		}
// 	}
// }

// func deleteProject(w http.ResponseWriter, r *http.Request) {
// 	cors.SetupCorsResponse(&w, r)
// 	if (*r).Method == "OPTIONS" {
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")

// 	var project models.Project
// 	vars := mux.Vars(r)
// 	id := vars["id"]
// 	project.ID = id
// 	err := utilities.App.DB.Table("projects").Unscoped().Delete(project).Error

// 	if err != nil {
// 		json.NewEncoder(w).Encode(err.Error())
// 		return
// 	} else {
// 		json.NewEncoder(w).Encode("Project deleted successfully")
// 		return
// 	}
// }

// func getProjectsByDepartment(w http.ResponseWriter, r *http.Request) {
// 	cors.SetupCorsResponse(&w, r)
// 	if (*r).Method == "OPTIONS" {
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")

// 	var projects []models.Project
// 	err := utilities.App.DB.Table("projects").Find(&projects, "department = ?", mux.Vars(r)["department"]).Error

// 	if err != nil {
// 		json.NewEncoder(w).Encode(err.Error())
// 		return
// 	}

// 	err = json.NewEncoder(w).Encode(projects)

// 	if err != nil {
// 		json.NewEncoder(w).Encode(err.Error())
// 	}
// }

// func getProjectsBySearch(w http.ResponseWriter, r *http.Request) {
// 	cors.SetupCorsResponse(&w, r)
// 	if (*r).Method == "OPTIONS" {
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")

// 	res := strings.Split(mux.Vars(r)["search_phrase"], " ")
// 	tx := utilities.App.DB.Table("projects")

// 	for _, element := range res {
// 		search_term := "%" + element + "%"
// 		tx = tx.Where("name LIKE ? OR email LIKE ?", search_term, search_term)
// 	}

// 	var projects []models.Project
// 	err := tx.Find(&projects).Error

// 	if err != nil {
// 		json.NewEncoder(w).Encode(err.Error())
// 		return
// 	}

// 	err = json.NewEncoder(w).Encode(projects)

// 	if err != nil {
// 		json.NewEncoder(w).Encode(err.Error())
// 	}
// }
