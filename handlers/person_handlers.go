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

func GetPeopleHandler(w http.ResponseWriter, r *http.Request) {
	people, err := repository.GetAllPeople()
	if err != nil {
		http.Error(w, "Error fetching people", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

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

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	count, err := repository.CountPeople()
	if err != nil {
		http.Error(w, "Stats failed", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]int{"total": count})
}
