package gokit

import (
	"fmt"
	"log"
	"os"
)

var (
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
)

func initLogger() {
	file := openLogFile()
	InfoLogger = createLog(file, "INFO")
	WarningLogger = createLog(file, "WARNING")
	ErrorLogger = createLog(file, "ERROR")
}

func openLogFile() *os.File {
	path, isPath := Env["APP_LOG_PATH"]
	fileName, isFileName := Env["APP_LOG_FILE"]
	if !isPath {
		path = "./"
	}
	if !isFileName {
		fileName = "file"
	}

	file, err := os.OpenFile(fmt.Sprintf("%s%s.log", path, fileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	return file
}

func createLog(file *os.File, logType string) *log.Logger {
	return log.New(file, logType+": ", log.Ldate|log.Ltime|log.Lshortfile)
}
