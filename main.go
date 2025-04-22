package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/Diony-source/peoplehub-api/utils"
	"github.com/Diony-source/peoplehub-api/repository"
)

func main() {
	godotenv.Load()
	utils.InitDB()

	err := repository.InsertPerson("Alice", 30)
	if err != nil {
		log.Fatal("Insert failed:", err)
	}

	log.Println("✔ Person inserted.")
}
