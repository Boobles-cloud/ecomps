package main

import (
	"context"
	"fmt"
	"os"
	"strings"

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

	// If its first init -> set to false
	if firstInit() {

		file, err := os.OpenFile(os.Getenv("env_path"), os.O_APPEND, 0600)

		if err != nil {
			logging.Log(logging.Error, "[Main | getting env file]"+err.Error())
		}

		defer file.Close()

		if _, err := file.WriteString("first-init=false"); err != nil {
			logging.Log(logging.Error, "[Main] Failed to write to env file! Please add first-init yourself! "+err.Error())
		}
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

func firstInit() bool {

	content, err := os.ReadFile(os.Getenv("env_path"))

	if err != nil {
		logging.Log(logging.Error, "[Main | firstInit]"+err.Error())
	}

	if strings.Contains(string(content), "first-init") {
		return false
	}

	return true
}
