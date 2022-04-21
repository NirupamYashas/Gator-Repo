package projects

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"server/models"
	"server/utilities"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initApp() {
	db, _ := gorm.Open(sqlite.Open("testing.db"), &gorm.Config{})
	db.AutoMigrate(&models.Project{})
	utilities.App = models.App{DB: db}
}

func TestGetProjects(t *testing.T) {
	initApp()
	proj := models.Project{
		ID:         uuid.New().String(),
		Name:       "test_project",
		Department: "test_department",
		Email:      "test@email.com",
		Link:       "test_link.com",
	}
	utilities.App.DB.Table("projects").Save(proj)

	req, _ := http.NewRequest("GET", "/api/projects", nil)
	r := httptest.NewRecorder()
	handler := http.HandlerFunc(GetProjects)

	handler.ServeHTTP(r, req)

	checkStatusCode(r.Code, http.StatusOK, t)
	checkContentType(r, t)
	checkBody(r.Body, proj, t)
	resetDB(firstProject())
}

func TestAddProject(t *testing.T) {
	initApp()
	var rqBody = toReader(`{
		"name": "test_name", 
		"department": "test_department", 
		"email": "test@email.com", 
		"link": "test_link.com"
	}`)
	req, _ := http.NewRequest("POST", "/api/projects", rqBody)
	r := httptest.NewRecorder()
	handler := http.HandlerFunc(AddProject)

	handler.ServeHTTP(r, req)

	checkStatusCode(r.Code, http.StatusCreated, t)
	checkContentType(r, t)
	checkProperties(firstProject(), t)
	resetDB(firstProject())
}

// func TestUpdateProject(t *testing.T) {
// 	initApp()
// 	id := uuid.New().String()
// 	proj := models.Project{
// 		ID:         id,
// 		Name:       "test_project",
// 		Department: "test_department",
// 		Email:      "test@email.com",
// 		Link:       "test_link.com",
// 	}
// 	utilities.App.DB.Table("projects").Save(proj)

// 	var rqBody = toReader(`{
// 		"name": "new_test_name",
// 		"department": "new_test_department",
// 		"email": "new_test@email.com",
// 		"link": "new_test_link.com"
// 	}`)

// 	req, _ := http.NewRequest("PUT", "/api/projects/"+id, rqBody)
// 	r := httptest.NewRecorder()
// 	handler := http.HandlerFunc(UpdateProject)

// 	handler.ServeHTTP(r, req)

// 	checkStatusCode(r.Code, http.StatusOK, t)
// 	checkContentType(r, t)
// 	checkUpdatedProperties(firstProject(), t)
// 	resetDB(firstProject())
// }

func toReader(content string) io.Reader {
	return bytes.NewBuffer([]byte(content))
}

func firstProject() models.Project {
	var all []models.Project
	utilities.App.DB.Table("projects").Find(&all)
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

// func checkUpdatedProperties(p models.Project, t *testing.T) {
// 	if p.Name != "new_test_name" {
// 		t.Errorf("Name should match: got %v want %v", p.Name, "new_test_name")
// 	}
// 	if p.Department != "new_test_department" {
// 		t.Errorf("Department should match: got %v want %v", p.Department, "new_test_department")
// 	}
// 	if p.Email != "new_test@email.com" {
// 		t.Errorf("Department should match: got %v want %v", p.Email, "new_test@email.com")
// 	}
// 	if p.Link != "new_test_link.com" {
// 		t.Errorf("Department should match: got %v want %v", p.Link, "new_test_link.com")
// 	}
// }

func resetDB(p models.Project) {
	utilities.App.DB.Table("projects").Delete(&p)
}
