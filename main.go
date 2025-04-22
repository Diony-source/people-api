package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/Diony-source/peoplehub-api/utils"
	"github.com/Diony-source/peoplehub-api/handlers"
)

func main() {
	godotenv.Load()
	utils.InitDB()

	http.HandleFunc("/people", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetPeopleHandler(w, r)
		case http.MethodPost:
			handlers.PostPersonHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/people/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodDelete:
			handlers.DeletePersonHandler(w, r)
		case http.MethodPatch:
			handlers.PatchPersonHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/people/search", handlers.SearchPeopleHandler)
	http.HandleFunc("/people/stats", handlers.StatsHandler)

	log.Println("🚀 Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
