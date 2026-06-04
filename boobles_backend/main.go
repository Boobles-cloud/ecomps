package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"boobles.cloud/backend/auth"
	"boobles.cloud/backend/auth/handlers"
	"boobles.cloud/backend/logging"
	"boobles.cloud/backend/startup"
)

func main() {

	ctx := context.Background()

	defer ctx.Done()

	logging.Log(logging.Information, "Startig application")

	if !startup.SetupTabels() {
		fmt.Println(logging.ErrorColor, "Failed to connect to the database... \nCheck the logs for more information!", logging.ResetColor)
		os.Exit(1)
	}

	// Starts our goroutine for deleting expired JWT
	go auth.DeleteExpiredJWT(ctx)

	muxRouter := http.NewServeMux()

	muxRouter.HandleFunc("GET /authwall/login", handlers.HandleLogin)
	muxRouter.HandleFunc("POST /authwall/register", handlers.HandleRegistration)

	if err := http.ListenAndServe(":8080", muxRouter); err != nil {
		logging.Log(logging.Error, err.Error())
		fmt.Println("Failed to start: ", err)
		os.Exit(1)
	}
}
