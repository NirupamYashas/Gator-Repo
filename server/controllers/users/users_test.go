package users

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"server/models"
	"server/utilities"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initApp() {
	db, _ := gorm.Open(sqlite.Open("testing.db"), &gorm.Config{})
	db.AutoMigrate(&models.User{})
	utilities.App = models.App{DB: db}
}

func TestGetUsers(t *testing.T) {
	initApp()
	user := models.User{
		ID:        uuid.New().String(),
		Firstname: "test_firstname",
		Lastname:  "test_lastname",
		Email:     "test@email.com",
		Password:  "test_password",
		Isadmin:   false,
		Created:   "2000-01-01",
	}
	utilities.App.DB.Table("users").Save(user)

	req, _ := http.NewRequest("GET", "/api/users", nil)
	r := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUsers)

	handler.ServeHTTP(r, req)

	checkStatusCode(r.Code, http.StatusOK, t)
	checkContentType(r, t)
	checkBody(r.Body, user, t)
	resetDB(firstUser())
}

func TestAddUser(t *testing.T) {
	initApp()
	var rqBody = toReader(`{
		"firstname": "test_firstname",
		"lastname": "test_lastname",
		"email": "test@email.com",
		"password": "test_password"
	}`)
	req, _ := http.NewRequest("POST", "/api/users/signup", rqBody)
	r := httptest.NewRecorder()
	handler := http.HandlerFunc(SignupUser)

	handler.ServeHTTP(r, req)

	checkStatusCode(r.Code, http.StatusOK, t)
	checkContentType(r, t)
	checkProperties(firstUser(), t)
	resetDB(firstUser())
}

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

func checkProperties(u models.User, t *testing.T) {
	if u.Firstname != "test_firstname" {
		t.Errorf("Firstname should match: got %v want %v", u.Firstname, "test_firstname")
	}
	if u.Lastname != "test_lastname" {
		t.Errorf("Lastname should match: got %v want %v", u.Lastname, "test_lastname")
	}
	if u.Email != "test@email.com" {
		t.Errorf("Email should match: got %v want %v", u.Email, "test@email.com")
	}
	if u.Password != "test_password" {
		t.Errorf("Password should match: got %v want %v", u.Password, "test_password")
	}
	if u.Isadmin != false {
		t.Errorf("Isadmin should match: got %v want %v", u.Isadmin, false)
	}
	if u.Created != time.Now().String()[:10] {
		t.Errorf("Created should match: got %v want %v", u.Created, time.Now().String()[:10])
	}
}

func resetDB(u models.User) {
	utilities.App.DB.Table("users").Delete(&u)
}
