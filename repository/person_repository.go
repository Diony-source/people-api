package repository

import (
	"github.com/Diony-source/peoplehub-api/utils"
)

func InsertPerson(name string, age int) error {
	query := `INSERT INTO people (name, age) VALUES ($1, $2)`
	_, err := utils.DB.Exec(query, name, age)
	return err
}
