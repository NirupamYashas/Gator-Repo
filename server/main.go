package main

import (
	"server/app"
	"server/models"
	"server/utilities"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("gator-repo.db"), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	utilities.App = models.App{
		DB: db,
		R:  mux.NewRouter(),
	}

	app.Start()
}
