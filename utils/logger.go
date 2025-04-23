package utils

import (
	"log"
	"os"
)

var Logger *log.Logger

func InitLogger() {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("❌ Failed to create log file: %v", err)
	}
	Logger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}
