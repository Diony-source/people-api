package repository

import (
	"strconv"

	"github.com/Diony-source/peoplehub-api/models"
	"github.com/Diony-source/peoplehub-api/utils"
)

func InsertPerson(name string, age int) error {
	_, err := utils.DB.Exec("INSERT INTO people (name, age) VALUES ($1, $2)", name, age)
	return err
}

func GetAllPeople() ([]models.Person, error) {
	var people []models.Person
	err := utils.DB.Select(&people, "SELECT * FROM people ORDER BY id")
	return people, err
}

func DeletePerson(id int) error {
	_, err := utils.DB.Exec("DELETE FROM people WHERE id = $1", id)
	return err
}

func UpdatePerson(id int, name *string, age *int) error {
	query := "UPDATE people SET "
	args := []interface{}{}
	argID := 1

	if name != nil {
		query += "name = $" + strconv.Itoa(argID)
		args = append(args, *name)
		argID++
	}
	if age != nil {
		if len(args) > 0 {
			query += ", "
		}
		query += "age = $" + strconv.Itoa(argID)
		args = append(args, *age)
		argID++
	}
	query += " WHERE id = $" + strconv.Itoa(argID)
	args = append(args, id)

	_, err := utils.DB.Exec(query, args...)
	return err
}

func SearchPeopleByName(name string) ([]models.Person, error) {
	var people []models.Person
	err := utils.DB.Select(&people, "SELECT * FROM people WHERE name ILIKE $1", "%"+name+"%")
	return people, err
}

func CountPeople() (int, error) {
	var count int
	err := utils.DB.Get(&count, "SELECT COUNT(*) FROM people")
	return count, err
}

func GetPeopleByID(id int) (models.Person, error) {
	var person models.Person
	err := utils.DB.Get(&person, "SELECT * FROM people WHERE id = $1", id)
	return person, err
}

func GetPeopleByAgeRange(minAge, maxAge int) ([]models.Person, error) {
	var people []models.Person
	err := utils.DB.Get(&people, "SELECT * FROM people WHERE age BETWEEN $1 AND $2", minAge, maxAge)
	return people, err
}

func GetRecentPeople(limit int) ([]models.Person, error) {
	var people []models.Person
	err := utils.DB.Get(&people, "SELECT * FROM people ORDER BY created_at DESC LIMIT $1", limit)
	return people, err
}
