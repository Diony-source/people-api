// @title PeopleHub API
//@version 1.0
// @description RESTful API for managing people
// @host localhost:8080
// @BasePath /
package main

import (
	"log"
	"net/http"

	"github.com/Diony-source/peoplehub-api/handlers"
	"github.com/Diony-source/peoplehub-api/utils"
	"github.com/joho/godotenv"

	_ "github.com/Diony-source/peoplehub-api/docs"
	httpSwagger "github.com/swaggo/http-swagger"
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
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	log.Println("🚀 Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
