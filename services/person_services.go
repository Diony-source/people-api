package service

import (
	"errors"
	"strings"

	"github.com/Diony-source/peoplehub-api/models"
	"github.com/Diony-source/peoplehub-api/repository"
)

var (
	repo repository.PersonRepository

	ErrInvalidName = errors.New("name must not be empty")
	ErrInvalidAge  = errors.New("age must be between 1 and 150")
)

func InjectRepository(r repository.PersonRepository) {
	repo = r
}

func InsertPerson(name string, age int) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return ErrInvalidName
	}
	if age <= 0 || age > 150 {
		return ErrInvalidAge
	}
	return repo.InsertPerson(name, age)
}

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
	return repo.UpdatePerson(id, name, age)
}

func GetPersonByID(id int) (models.Person, error) {
	if id <= 0 {
		return models.Person{}, errors.New("invalid id")
	}
	return repo.GetPeopleByID(id)
}

func SearchPeopleByName(name string) ([]models.Person, error) {
	return repository.SearchPeopleByName(name)
}

func CountPeople() (int, error) {
	return repository.CountPeople()
}

func GetPeopleByAgeRange(minAge, maxAge int) ([]models.Person, error) {
	return repository.GetPeopleByAgeRange(minAge, maxAge)
}

func GetRecentPeople(limit int) ([]models.Person, error) {
	return repository.GetRecentPeople(limit)
}

