package logging

import (
	"errors"
	"log"
	"os"
	"path"
	"time"
)

const (
	// Spaces for better view :)
	Error       string = " [Error] "
	Information string = " [Information] "
)

// Creates a new log entry
// needs a type and a message
// pls only use the given constants from this package
func Log(logType, logMsg string) {

	file := getCurrLogFile()

	currLogFile, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer currLogFile.Close()

	if _, err := currLogFile.WriteString("\n" + time.Now().Format("2006/01/02 15:04:05") + logType + logMsg); err != nil {
		log.Fatal(err)
	}
}

func getCurrLogFile() string {
	currDir, _ := os.Getwd()
	currDir += "\\logs"

	if _, err := os.Stat(currDir); errors.Is(err, os.ErrNotExist) {
		os.Mkdir(currDir, os.ModeDir)
	}

	files, err := os.ReadDir(currDir)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fileInfo, err := file.Info()

		if err != nil {
			continue
		}

		if time.Now().Format(time.DateOnly) == fileInfo.ModTime().Format(time.DateOnly) {
			return path.Join(currDir, file.Name())
		} else {
			continue
		}
	}

	date := time.Now().Truncate(24 * time.Hour)

	const formatString = "02-01-2006"

	newLogFile := date.Format(formatString) + ".log"

	nFile, err := os.Create(path.Join(currDir, newLogFile))

	if err != nil {
		log.Fatal(err)
	}

	defer nFile.Close()

	return nFile.Name()
}
