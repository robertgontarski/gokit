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

	file, err := os.OpenFile(fmt.Sprintf("var/log/%s.log", Env["APP_ENV"]), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	return file
}

func createLog(file *os.File, logType string) *log.Logger {
	return log.New(file, logType+": ", log.Ldate|log.Ltime|log.Lshortfile)
}
