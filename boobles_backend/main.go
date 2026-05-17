package main

import (
	"fmt"
	"os"

	"boobles.cloud/backend/startup"
)

func main() {

	fmt.Println("Starting...")

	if !startup.SetupTabels() {
		fmt.Println("Failed to connect to the database... \n Check the logs for more information!")
		os.Exit(1)
	}
}
