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
