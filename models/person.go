package models

type Person struct {
	ID        int    `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	Age       int    `db:"age" json:"age"`
	CreatedAt string `db:"created_at" json:"created_at"`
}
