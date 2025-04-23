package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Diony-source/peoplehub-api/repository"
)

type createPersonRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type updatePersonRequest struct {
	Name *string `json:"name"`
	Age  *int    `json:"age"`
}

// GetPeopleHandler godoc
// @Summary List all people
// @Produce json
// @Success 200 {array} models.Person
// @Router /people [get]
func GetPeopleHandler(w http.ResponseWriter, r *http.Request) {
	people, err := repository.GetAllPeople()
	if err != nil {
		http.Error(w, "Error fetching people", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

// PostPersonHandler godoc
// @Summary Add a new person
// @Accept json
// @Produce plain
// @Param person body createPersonRequest true "Person JSON"
// @Success 201 {string} string "Created"
// @Failure 400 {string} string "Invalid"
// @Router /people [post]
func PostPersonHandler(w http.ResponseWriter, r *http.Request) {
	var req createPersonRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Name == "" || req.Age <= 0 {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if err := repository.InsertPerson(req.Name, req.Age); err != nil {
		http.Error(w, "Failed to insert person", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// DeletePersonHandler godoc
// @Summary Delete person by ID
// @Param id path int true "Person ID"
// @Success 204 {string} string "No Content"
// @Router /people/{id} [delete]
func DeletePersonHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/people/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := repository.DeletePerson(id); err != nil {
		http.Error(w, "Delete failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// PatchPersonHandler godoc
// @Summary Update person fields
// @Accept json
// @Param id path int true "Person ID"
// @Param update body updatePersonRequest true "Updated fields"
// @Success 200 {string} string "OK"
// @Router /people/{id} [patch]
func PatchPersonHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/people/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var req updatePersonRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}
	if err := repository.UpdatePerson(id, req.Name, req.Age); err != nil {
		http.Error(w, "Update failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// SearchPeopleHandler godoc
// @Summary Search people by name
// @Param name query string true "Name keyword"
// @Success 200 {array} models.Person
// @Router /people/search [get]
func SearchPeopleHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Missing name param", http.StatusBadRequest)
		return
	}
	people, err := repository.SearchPeopleByName(name)
	if err != nil {
		http.Error(w, "Search failed", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(people)
}

// StatsHandler godoc
// @Summary Get total number of people
// @Success 200 {object} map[string]int
// @Router /people/stats [get]
func StatsHandler(w http.ResponseWriter, r *http.Request) {
	count, err := repository.CountPeople()
	if err != nil {
		http.Error(w, "Stats failed", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]int{"total": count})
}

// GetPersonByIDHandler godoc
// @Summary Get person by ID
// @Param id path int true "Person ID"
// @Success 200 {object} models.Person
// @Failure 404 {string} string "Not Found"
// @Router /people/{id} [get]
func GetPersonByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/people/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	person, err := repository.GetPeopleByID(id)
	if err != nil {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

func GetPeopleByAgeRangeHandler(w http.ResponseWriter, r *http.Request) {
	min, _ := strconv.Atoi(r.URL.Query().Get("min"))
	max, _ := strconv.Atoi(r.URL.Query().Get("max"))

	people, err := repository.GetPeopleByAgeRange(min, max)
	if err != nil {
		http.Error(w, "Error fetching people", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(people)
}

// GetRecentPeopleHandler godoc
// @Summary Get recent people
// @Tags People
// @Param limit query int false "Limit number of people"
// @Success 200 {array} models.Person
// @Router /people/recent [get]
func GetRecentPeopleHandler(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit <= 0 {
		limit = 5
	}
	people, err := repository.GetRecentPeople(limit)
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(people)
}
