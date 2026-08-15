package main

import (
	"context"
	"fmt"
	"os"

	"boobles.cloud/backend/database"
	hostedServices "boobles.cloud/backend/hosted_services"
	"boobles.cloud/backend/logging"
	"boobles.cloud/backend/startup"
)

func main() {

	ctx := context.Background()

	defer ctx.Done()

	databaseConf, ok := database.CreateDbHandler()

	if !ok {
		fmt.Println(logging.ErrorColor, "Failed to connect to the database... \nCheck the logs for more information!", logging.ResetColor)
		os.Exit(1)
	}

	logging.Log(logging.Information, "Startig application")

	if !startup.SetupTabels(databaseConf) {
		fmt.Println(logging.ErrorColor, "Failed to setup table \nCheck the logs for more information!", logging.ResetColor)
		os.Exit(1)
	}

	if !startup.GenerateFrontendApiToken() {
		fmt.Println(logging.ErrorColor, "Failed to generate api key...", logging.ResetColor)
		os.Exit(1)
	}

	// If first init -> set to false
	if f := os.Getenv("first-init"); f == "true" {
		os.Setenv("first-init", "false")
	}

	// Starts our goroutine for deleting expired JWT
	go hostedServices.DeleteExpiredJWT(ctx, databaseConf)
	go hostedServices.DeleteTenants(ctx, databaseConf)

	// ============ REST-API config stuff ============
	httpServer := startup.ConfigureHTTPServer(databaseConf)

	logging.Log(logging.Information, "Boobles starting and listening on :8080")

	if err := httpServer.ListenAndServe(); err != nil {
		logging.Log(logging.Error, err.Error())
		fmt.Println(logging.ErrorColor, "Failed to start: ", err, logging.ResetColor)
		os.Exit(1)
	}
}
