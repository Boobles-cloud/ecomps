package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"boobles.cloud/backend/internal/auth"
	authHandlers "boobles.cloud/backend/internal/auth/handlers"
	"boobles.cloud/backend/internal/middleware"
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

	// ============ REST-API config stuff ============

	// Configuring our global middleware
	globalMiddlewareConfig := middleware.CreateNewMiddlewareStack(
		middleware.PanicRecoverMiddleware,
		// TODO: add logging middleware here
	)

	muxRouter := http.NewServeMux()

	// ============ Auth stuff ============
	muxRouter.HandleFunc("GET /authwall/login", authHandlers.HandleLogin)
	muxRouter.HandleFunc("POST /authwall/register", authHandlers.HandleRegistration)
	muxRouter.HandleFunc("GET /authwall/logout", authHandlers.HandleLogout)

	// ============ Tenant stuff ============
	// Tenant middleware config
	tenantMiddleware := middleware.CreateNewMiddlewareStack(
		middleware.PanicRecoverMiddleware,
		middleware.AuthMiddleware,
		middleware.PermissionMiddleware,
	)

	tenantRouter := http.NewServeMux()
	tenantRouter.Handle("/", tenantMiddleware(tenantRouter))

	// POST Requests
	tenantRouter.HandleFunc("POST /tenant/create", tenantHandlers.HandleTenantCreation)
	tenantRouter.HandleFunc("POST /tenant/change", tenantHandlers.HandleTenantChange)
	tenantRouter.HandleFunc("POST /tenant/delete", tenantHandlers.HandleTenantDeltion)

	// GET Requests
	tenantRouter.HandleFunc("GET /tenant/{tenant-id}", tenantHandlers.HandleGetTenantByTenantId)
	tenantRouter.HandleFunc("GET /tenant/by/user={user-id}", tenantHandlers.HandleGetTenantByUserId)

	// Creating our http server
	httpServer := http.Server{
		Addr:    ":8080",
		Handler: globalMiddlewareConfig(muxRouter),
	}

	if err := httpServer.ListenAndServe(); err != nil {
		logging.Log(logging.Error, err.Error())
		fmt.Println(logging.ErrorColor, "Failed to start: ", err, logging.ResetColor)
		os.Exit(1)
	}
}
