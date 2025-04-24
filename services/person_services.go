package service

import (
	"errors"
	"strings"

	"github.com/Diony-source/peoplehub-api/models"
	"github.com/Diony-source/peoplehub-api/repository"
)

var (
	ErrInvalidName = errors.New("name must not be empty")
	ErrInvalidAge  = errors.New("age must be between 1 and 150")
)

// InsertPerson validates input and delegates to repository
func InsertPerson(name string, age int) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return ErrInvalidName
	}
	if age <= 0 || age > 150 {
		return ErrInvalidAge
	}
	return repository.InsertPerson(name, age)
}

// UpdatePerson validates optional fields and updates the person
func UpdatePerson(id int, name *string, age *int) error {
	if name != nil {
		trimmed := strings.TrimSpace(*name)
		if trimmed == "" {
			return ErrInvalidName
		}
		*name = trimmed
	}
	if age != nil {
		if *age <= 0 || *age > 150 {
			return ErrInvalidAge
		}
	}
	return repository.UpdatePerson(id, name, age)
}

func GetPersonByID(id int) (models.Person, error) {
	if id <= 0 {
		return models.Person{}, errors.New("invalid id")
	}
	return repository.GetPeopleByID(id)
}
