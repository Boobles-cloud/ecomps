package main

import (
	"context"
	"fmt"
	"os"

	"boobles.cloud/backend/internal/auth"
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

	// ============ REST-API config stuff ============
	httpServer := startup.ConfigureHTTPServer()

	if err := httpServer.ListenAndServe(); err != nil {
		logging.Log(logging.Error, err.Error())
		fmt.Println(logging.ErrorColor, "Failed to start: ", err, logging.ResetColor)
		os.Exit(1)
	}
}
