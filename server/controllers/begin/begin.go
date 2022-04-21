package begin

import (
	"encoding/json"
	"net/http"
	"server/cors"
	"server/utilities"
)

func Start() {
	addRoutes()
}

func addRoutes() {
	utilities.App.R.HandleFunc("/api/test", Test).Methods("GET", "OPTIONS")
}

func Test(w http.ResponseWriter, r *http.Request) {
	cors.SetupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{"message": "Hello, World!"})
}
