package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Diony-source/peoplehub-api/services"
	"github.com/Diony-source/peoplehub-api/utils"
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
	people, err := service.GetAllPeople()
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, "Error fetching people", err)
		return
	}
	utils.JSON(w, http.StatusOK, people)
}

func PostPersonHandler(w http.ResponseWriter, r *http.Request) {
	var req createPersonRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}
	if err := service.InsertPerson(req.Name, req.Age); err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid input", err)
		return
	}
	utils.JSON(w, http.StatusCreated, map[string]string{"message": "Person created"})
}

func DeletePersonHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/people/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid ID", err)
		return
	}
	if err := service.DeletePerson(id); err != nil {
		utils.Error(w, http.StatusInternalServerError, "Delete failed", err)
		return
	}
	utils.JSON(w, http.StatusNoContent, nil)
}

func PatchPersonHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/people/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid ID", err)
		return
	}
	var req updatePersonRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid body", err)
		return
	}
	if err := service.UpdatePerson(id, req.Name, req.Age); err != nil {
		utils.Error(w, http.StatusBadRequest, "Update failed", err)
		return
	}
	utils.JSON(w, http.StatusOK, map[string]string{"message": "Person updated"})
}

func GetPeopleByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/people/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid ID", err)
		return
	}
	person, err := service.GetPersonByID(id)
	if err != nil {
		utils.Error(w, http.StatusNotFound, "Person not found", err)
		return
	}
	utils.JSON(w, http.StatusOK, person)
}

func SearchPeopleHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		utils.Error(w, http.StatusBadRequest, "Missing name param", nil)
		return
	}
	people, err := service.SearchPeopleByName(name)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, "Search failed", err)
		return
	}
	utils.JSON(w, http.StatusOK, people)
}

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	count, err := service.CountPeople()
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, "Stats failed", err)
		return
	}
	utils.JSON(w, http.StatusOK, map[string]int{"total": count})
}

func GetPeopleByAgeRangeHandler(w http.ResponseWriter, r *http.Request) {
	min, _ := strconv.Atoi(r.URL.Query().Get("min"))
	max, _ := strconv.Atoi(r.URL.Query().Get("max"))

	people, err := service.GetPeopleByAgeRange(min, max)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, "Error fetching people", err)
		return
	}
	utils.JSON(w, http.StatusOK, people)
}

func GetRecentPeopleHandler(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit <= 0 {
		limit = 5
	}
	people, err := service.GetRecentPeople(limit)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, "Error fetching data", err)
		return
	}
	utils.JSON(w, http.StatusOK, people)
}
