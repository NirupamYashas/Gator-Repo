package app

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"server/controllers/projects"
	"server/models"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initApp() models.App {
	db, _ := gorm.Open(sqlite.Open("testing.db"), &gorm.Config{})
	db.AutoMigrate(&models.Project{})
	return models.App{DB: db}
}

func TestGetProjects(t *testing.T) {
	app := initApp()
	proj := models.Project{
		ID:         uuid.New().String(),
		Name:       "test_project",
		Department: "test_department",
		Email:      "test@email.com",
		Link:       "test_link.com",
	}
	app.DB.Table("projects").Save(proj)

	req, _ := http.NewRequest("GET", "/api/projects", nil)
	r := httptest.NewRecorder()
	handler := http.HandlerFunc(projects.GetProjects)

	handler.ServeHTTP(r, req)

	checkStatusCode(r.Code, http.StatusOK, t)
	checkContentType(r, t)
	checkBody(r.Body, proj, t)
}

func TestAddProject(t *testing.T) {
	app := initApp()
	var rqBody = toReader(`{
		"id": "test_id", 
		"name": "test_project", 
		"department": "test_department", 
		"email": "test@email.com", 
		"link": "test_link.com"
	}`)
	req, _ := http.NewRequest("POST", "/api/projects", rqBody)
	r := httptest.NewRecorder()
	handler := http.HandlerFunc(projects.AddProject)

	handler.ServeHTTP(r, req)

	checkStatusCode(r.Code, http.StatusCreated, t)
	checkContentType(r, t)
	checkProperties(firstProject(app), t)
}

func toReader(content string) io.Reader {
	return bytes.NewBuffer([]byte(content))
}

func firstProject(app models.App) models.Project {
	var all []models.Project
	app.DB.Find(&all)
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

func checkBody(body *bytes.Buffer, st models.Project, t *testing.T) {
	var projects []models.Project
	_ = json.Unmarshal(body.Bytes(), &projects)
	if len(projects) != 1 {
		t.Errorf("Wrong lenght: got %v want 1", len(projects))
	}
	if projects[0] != st {
		t.Errorf("Wrong body: got %v want %v", projects[0], st)
	}
}

func checkProperties(p models.Project, t *testing.T) {
	if p.ID != "test_id" {
		t.Errorf("ID should match: got %v want %v", p.ID, "test_id")
	}
	if p.Name != "test_name" {
		t.Errorf("Name should match: got %v want %v", p.Name, "test_name")
	}
	if p.Department != "test_department" {
		t.Errorf("Department should match: got %v want %v", p.Department, "test_department")
	}
	if p.Email != "test@email.com" {
		t.Errorf("Department should match: got %v want %v", p.Email, "test@email.com")
	}
	if p.Link != "test_link.com" {
		t.Errorf("Department should match: got %v want %v", p.Link, "test_link.com")
	}
}
