package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var logger *log.Logger
var logFile *os.File

func InitLogger() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	exPath := filepath.Dir(ex)

	logFile, err = os.OpenFile(exPath+"/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	logger = log.New(logFile, "", log.Ldate|log.Ltime)
}

func Info(message string) {
	logger.Println("INFO: " + message)
}
