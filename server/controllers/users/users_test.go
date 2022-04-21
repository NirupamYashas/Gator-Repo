package users

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"server/models"
	"server/utilities"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initApp() {
	db, _ := gorm.Open(sqlite.Open("testing.db"), &gorm.Config{})
	db.AutoMigrate(&models.User{})
	utilities.App = models.App{DB: db}
}

// func TestGetProjects(t *testing.T) {
// 	initApp()
// 	proj := models.Project{
// 		ID:         uuid.New().String(),
// 		Name:       "test_project",
// 		Department: "test_department",
// 		Email:      "test@email.com",
// 		Link:       "test_link.com",
// 	}
// 	utilities.App.DB.Table("projects").Save(proj)

// 	req, _ := http.NewRequest("GET", "/api/projects", nil)
// 	r := httptest.NewRecorder()
// 	handler := http.HandlerFunc(GetProjects)

// 	handler.ServeHTTP(r, req)

// 	checkStatusCode(r.Code, http.StatusOK, t)
// 	checkContentType(r, t)
// 	checkBody(r.Body, proj, t)
// 	resetDB(firstProject(utilities.App))
// }

// func TestAddProject(t *testing.T) {
// 	initApp()
// 	var rqBody = toReader(`{
// 		"name": "test_name",
// 		"department": "test_department",
// 		"email": "test@email.com",
// 		"link": "test_link.com"
// 	}`)
// 	req, _ := http.NewRequest("POST", "/api/projects", rqBody)
// 	r := httptest.NewRecorder()
// 	handler := http.HandlerFunc(AddProject)

// 	handler.ServeHTTP(r, req)

// 	checkStatusCode(r.Code, http.StatusCreated, t)
// 	checkContentType(r, t)
// 	checkProperties(firstProject(utilities.App), t)
// 	resetDB(firstProject(utilities.App))
// }

func toReader(content string) io.Reader {
	return bytes.NewBuffer([]byte(content))
}

func firstUser() models.User {
	var all []models.User
	utilities.App.DB.Table("users").Find(&all)
	return all[0]
}

func checkStatusCode(code int, want int, t *testing.T) {
	if code != want {
		t.Errorf("Wrong status code: got %v want %v", code, want)
	}
}

func checkContentType(r *httptest.ResponseRecorder, t *testing.T) {
	ct := r.Header().Get("Content-Type")
	if ct != "application/json" {
		t.Errorf("Wrong Content Type: got %v want application/json", ct)
	}
}

func checkBody(body *bytes.Buffer, st models.User, t *testing.T) {
	var users []models.User
	_ = json.Unmarshal(body.Bytes(), &users)
	if len(users) != 1 {
		t.Errorf("Wrong lenght: got %v want 1", len(users))
	}
	if users[0] != st {
		t.Errorf("Wrong body: got %v want %v", users[0], st)
	}
}

func checkProperties(p models.User, t *testing.T) {
	// if p.Name != "test_name" {
	// 	t.Errorf("Name should match: got %v want %v", p.Name, "test_name")
	// }
	// if p.Department != "test_department" {
	// 	t.Errorf("Department should match: got %v want %v", p.Department, "test_department")
	// }
	// if p.Email != "test@email.com" {
	// 	t.Errorf("Department should match: got %v want %v", p.Email, "test@email.com")
	// }
	// if p.Link != "test_link.com" {
	// 	t.Errorf("Department should match: got %v want %v", p.Link, "test_link.com")
	// }
}

func resetDB(p models.User) {
	utilities.App.DB.Table("users").Delete(&p)
}
