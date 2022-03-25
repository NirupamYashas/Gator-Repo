package app

import (
		"bytes"
		"encoding/json"
		"net/http"
		"net/http/httptest"
		"testing"
	
		"github.com/google/uuid"
		"gorm.io/driver/sqlite"
		"gorm.io/gorm"
	)
	
	func initApp() App {
		db, _ := gorm.Open(sqlite.Open("testing.db"), &gorm.Config{})
		db.AutoMigrate(&Project{})
		return App{DB: db}
	}
	
	
	