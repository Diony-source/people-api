package service

import (
	"errors"
	"testing"

	"github.com/Diony-source/peoplehub-api/models"
)

type mockRepo struct{}

func (m mockRepo) InsertPerson(name string, age int) error {
	if name == "fail" {
		return errors.New("insert failed")
	}
	return nil
}

func (m mockRepo) UpdatePerson(id int, name *string, age *int) error {
	return nil
}

func (m mockRepo) GetPeopleByID(id int) (models.Person, error) {
	if id == 999 {
		return models.Person{}, errors.New("not found")
	}
	return models.Person{ID: id, Name: "Mock", Age: 30}, nil
}

func (m mockRepo) SearchPeopleByName(name string) ([]models.Person, error) {
	if name == "fail" {
		return nil, errors.New("search error")
	}
	return []models.Person{{ID: 1, Name: "Mock", Age: 25}}, nil
}

func (m mockRepo) CountPeople() (int, error) {
	return 42, nil
}

func (m mockRepo) GetPeopleByAgeRange(minAge, maxAge int) ([]models.Person, error) {
	return []models.Person{{ID: 2, Name: "AgePerson", Age: 30}}, nil
}

func (m mockRepo) GetRecentPeople(limit int) ([]models.Person, error) {
	return []models.Person{{ID: 3, Name: "Recent", Age: 24}}, nil
}

func TestInsertPerson_Valid(t *testing.T) {
	InjectRepository(mockRepo{})
	err := InsertPerson("John", 25)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestInsertPerson_Invalid(t *testing.T) {
	InjectRepository(mockRepo{})
	if err := InsertPerson("", 25); err != ErrInvalidName {
		t.Errorf("expected name error, got %v", err)
	}
	if err := InsertPerson("John", -1); err != ErrInvalidAge {
		t.Errorf("expected age error, got %v", err)
	}
}

func TestGetPersonByID_InvalidID(t *testing.T) {
	InjectRepository(mockRepo{})
	_, err := GetPersonByID(0)
	if err == nil {
		t.Error("expected error for invalid id")
	}
}

func TestGetPersonByID_NotFound(t *testing.T) {
	InjectRepository(mockRepo{})
	_, err := GetPersonByID(999)
	if err == nil {
		t.Error("expected error for not found")
	}
}

func TestUpdatePerson_InvalidName(t *testing.T) {
	InjectRepository(mockRepo{})
	empty := ""
	err := UpdatePerson(1, &empty, nil)
	if err != ErrInvalidName {
		t.Errorf("expected ErrInvalidName, got %v", err)
	}
}

func TestUpdatePerson_InvalidAge(t *testing.T) {
	InjectRepository(mockRepo{})
	age := -10
	err := UpdatePerson(1, nil, &age)
	if err != ErrInvalidAge {
		t.Errorf("expected ErrInvalidAge, got %v", err)
	}
}

func TestSearchPeopleByName(t *testing.T) {
	InjectRepository(mockRepo{})
	people, err := SearchPeopleByName("Mock")
	if err != nil || len(people) == 0 {
		t.Error("expected results for valid name")
	}
}

func TestCountPeople(t *testing.T) {
	InjectRepository(mockRepo{})
	count, err := CountPeople()
	if err != nil || count != 42 {
		t.Errorf("expected 42, got %d (err: %v)", count, err)
	}
}

func TestGetPeopleByAgeRange(t *testing.T) {
	InjectRepository(mockRepo{})
	people, err := GetPeopleByAgeRange(20, 40)
	if err != nil || len(people) == 0 {
		t.Error("expected results for valid range")
	}
}

func TestGetRecentPeople(t *testing.T) {
	InjectRepository(mockRepo{})
	people, err := GetRecentPeople(5)
	if err != nil || len(people) == 0 {
		t.Error("expected recent people")
	}
}
