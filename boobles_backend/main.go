package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"boobles.cloud/backend/internal/auth"
	authHandlers "boobles.cloud/backend/internal/auth/handlers"
	tenantHandlers "boobles.cloud/backend/internal/tenant/handlers"
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

	// ============ Auth stuff ============
	muxRouter.HandleFunc("GET /authwall/login", authHandlers.HandleLogin)
	muxRouter.HandleFunc("POST /authwall/register", authHandlers.HandleRegistration)
	muxRouter.HandleFunc("GET /authwall/logout", authHandlers.HandleLogout)

	// ============ Tenant stuff ============
	// NOTE: if you add more always go through the middleware!!
	// TODO: Add permissions middleware to check if the user has the rights to do this
	muxRouter.Handle("POST /tenant/create", auth.AuthMiddleware(http.HandlerFunc(tenantHandlers.HandleTenantCreation)))
	muxRouter.Handle("POST /tenant/change", auth.AuthMiddleware(http.HandlerFunc(tenantHandlers.HandleTenantChange)))
	muxRouter.Handle("POST /tenant/delete", auth.AuthMiddleware(http.HandlerFunc(tenantHandlers.HandleTenantDeltion)))

	muxRouter.Handle("GET /tenant/{tenant-id}", auth.AuthMiddleware(http.HandlerFunc(tenantHandlers.HandleGetTenantByTenantId)))
	muxRouter.Handle("GET /tenant/by/user={user-id}", auth.AuthMiddleware(http.HandlerFunc(tenantHandlers.HandleGetTenantByUserId)))

	if err := http.ListenAndServe(":8080", muxRouter); err != nil {
		logging.Log(logging.Error, err.Error())
		fmt.Println(logging.ErrorColor, "Failed to start: ", err, logging.ResetColor)
		os.Exit(1)
	}
}
