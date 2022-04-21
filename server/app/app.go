package app

import (
	"log"
	"net/http"

	"server/utilities"

	"server/controllers/begin"
	"server/controllers/projects"
	"server/controllers/users"
)

func Start() {
	projects.Start()
	users.Start()
	begin.Start()

	log.Fatal(http.ListenAndServe(":8080", utilities.App.R))
}
