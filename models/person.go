package models

type Person struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	Age       int    `db:"age"`
	CreatedAt string `db:"created_at"`
}
