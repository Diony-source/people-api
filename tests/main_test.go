package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/Diony-source/peoplehub-api/handlers"
	"github.com/Diony-source/peoplehub-api/utils"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	utils.InitDB()
}

func TestPostPersonHandler(t *testing.T) {
	body := []byte(`{"name":"TestUser","age":40}`)
	req := httptest.NewRequest(http.MethodPost, "/people", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handlers.PostPersonHandler(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", rr.Code)
	}
}

func TestGetPeopleHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/people", nil)
	rr := httptest.NewRecorder()

	handlers.GetPeopleHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}
}

func TestSearchPeopleHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/people/search?name=Test", nil)
	rr := httptest.NewRecorder()

	handlers.SearchPeopleHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}
}

func TestStatsHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/people/stats", nil)
	rr := httptest.NewRecorder()

	handlers.StatsHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}
}

func TestPatchAndDeletePerson(t *testing.T) {
	// Create person first
	body := []byte(`{"name":"PatchUser","age":33}`)
	postReq := httptest.NewRequest(http.MethodPost, "/people", bytes.NewBuffer(body))
	postReq.Header.Set("Content-Type", "application/json")
	postRes := httptest.NewRecorder()
	handlers.PostPersonHandler(postRes, postReq)
	if postRes.Code != http.StatusCreated {
		t.Fatalf("Person creation failed")
	}

	// Get ID from last created person
	getReq := httptest.NewRequest(http.MethodGet, "/people", nil)
	getRes := httptest.NewRecorder()
	handlers.GetPeopleHandler(getRes, getReq)

	var people []map[string]interface{}
	json.Unmarshal(getRes.Body.Bytes(), &people)

	if len(people) == 0 {
		t.Fatalf("No people to test PATCH/DELETE")
	}

	lastID := int(people[len(people)-1]["id"].(float64))

	// PATCH
	patchBody := []byte(`{"age":99}`)
	patchReq := httptest.NewRequest(http.MethodPatch, "/people/"+strconv.Itoa(lastID), bytes.NewBuffer(patchBody))
	patchRes := httptest.NewRecorder()
	handlers.PatchPersonHandler(patchRes, patchReq)

	if patchRes.Code != http.StatusOK {
		t.Errorf("PATCH failed, got %d", patchRes.Code)
	}

	// DELETE
	deleteReq := httptest.NewRequest(http.MethodDelete, "/people/"+strconv.Itoa(lastID), nil)
	deleteRes := httptest.NewRecorder()
	handlers.DeletePersonHandler(deleteRes, deleteReq)

	if deleteRes.Code != http.StatusNoContent {
		t.Errorf("DELETE failed, got %d", deleteRes.Code)
	}
}

func TestPostPersonValidation(t *testing.T) {
	body := `{"name":"", "age":25}`
	req, _ := http.NewRequest("POST", "/people", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.PostPersonHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected 400 BadRequest for empty name, got %v", rr.Code)
	}

	body = `{"name":"John", "age":0}`
	req, _ = http.NewRequest("POST", "/people", strings.NewReader(body))
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected 400 BadRequest for invalid age, got %v", rr.Code)
	}
}

func TestGetPersonByID(t *testing.T) {
	// Geçerli ID (örnek: 1)
	req, _ := http.NewRequest("GET", "/people/1", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetPeopleByIDHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK && rr.Code != http.StatusNotFound {
		t.Errorf("expected 200 OK or 404 Not Found, got %v", rr.Code)
	}

	// Geçersiz ID
	req, _ = http.NewRequest("GET", "/people/99999", nil)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("expected 404 Not Found for invalid ID, got %v", rr.Code)
	}
}
