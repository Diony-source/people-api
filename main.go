package main

import (
	"log"
	"net/http"

	"github.com/Diony-source/peoplehub-api/handlers"
	"github.com/Diony-source/peoplehub-api/utils"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	utils.InitDB()
	utils.InitLogger() // ✅ log.txt sistemi başlat

	http.HandleFunc("/people", func(w http.ResponseWriter, r *http.Request) {
		utils.Logger.Println("🔄 /people endpoint hit")
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
		utils.Logger.Println("📌 /people/{id} endpoint hit")
		switch r.Method {
		case http.MethodDelete:
			handlers.DeletePersonHandler(w, r)
		case http.MethodPatch:
			handlers.PatchPersonHandler(w, r)
		case http.MethodGet:
			handlers.GetPeopleByIDHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/people/search", handlers.SearchPeopleHandler)
	http.HandleFunc("/people/stats", handlers.StatsHandler)
	http.HandleFunc("/people/age", handlers.GetPeopleByAgeRangeHandler)
	http.HandleFunc("/people/recent", handlers.GetRecentPeopleHandler)

	utils.Logger.Println("🚀 Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
