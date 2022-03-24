package main

import (
	"server/app"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("gator-repo.db"), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	app := app.App{
		DB: db,
		R:  mux.NewRouter(),
	}

	app.Start()
}
