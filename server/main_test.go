package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/google/uuid"
// 	"gorm.io/driver/sqlite"
// 	"gorm.io/gorm"
// )

// func initApp() App {
// 	db, _ := gorm.Open(sqlite.Open("testing.db"), &gorm.Config{})
// 	db.AutoMigrate(&Project{})
// 	return App{db: db}
// }

// func TestGetProjects(t *testing.T) {
// 	app := initApp()
// 	proj := Project{ID: uuid.New().String(), Name: "Game Project", Department: "CISE", Email: "game@gmail.com", Link: "github.com/game"}
// 	app.db.Save(proj)

// 	req, _ := http.NewRequest("GET", "/api/projects", nil)
// 	r := httptest.NewRecorder()
// 	handler := http.HandlerFunc(app.getProjects)

// 	handler.ServeHTTP(r, req)

// 	checkStatusCode(r.Code, http.StatusOK, t)
// 	checkContentType(r, t)
// 	checkBody(r.Body, proj, t)
// }

// func checkStatusCode(code int, want int, t *testing.T) {
// 	if code != want {
// 		t.Errorf("Wrong status code: got %v want %v", code, want)
// 	}
// }

// func checkContentType(r *httptest.ResponseRecorder, t *testing.T) {
// 	ct := r.Header().Get("Content-Type")
// 	if ct != "application/json" {
// 		t.Errorf("Wrong Content Type: got %v want application/json", ct)
// 	}
// }

// func checkBody(body *bytes.Buffer, st Project, t *testing.T) {
// 	var projects []Project
// 	_ = json.Unmarshal(body.Bytes(), &projects)
// 	if len(projects) != 1 {
// 		t.Errorf("Wrong lenght: got %v want 1", len(projects))
// 	}
// 	if projects[0] != st {
// 		t.Errorf("Wrong body: got %v want %v", projects[0], st)
// 	}
// }
