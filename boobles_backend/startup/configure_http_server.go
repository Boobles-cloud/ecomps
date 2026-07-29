package startup

import (
	"net/http"

	authHandlers "boobles.cloud/backend/internal/auth/handlers"
	"boobles.cloud/backend/internal/middleware"
	tenantHandlers "boobles.cloud/backend/internal/tenant/handlers"
	userHandlers "boobles.cloud/backend/internal/user/handlers"
)

// Creates and configures the rest api
func ConfigureHTTPServer() http.Server {

	// ============ Middleware config stuff ============

	// Configuring our global middleware
	globalMiddlewareConfig := middleware.CreateNewMiddlewareStack(
		middleware.LoggingMiddleware,
		middleware.PanicRecoverMiddleware,
	)

	// Tenant middleware config
	tenantMiddleware := middleware.CreateNewMiddlewareStack(
		middleware.AuthMiddleware,
		middleware.PermissionMiddleware,
	)

	// Tenant permission changes -> this can only be done by the admin
	tenantPermissionMiddleware := middleware.CreateNewMiddlewareStack(
		middleware.AuthMiddleware,
		middleware.CheckAdminMiddleware,
	)

	// For changes on the user config itself
	userMiddleware := middleware.CreateNewMiddlewareStack(
		middleware.AuthMiddleware,
	)

	// For the frontend to get user information
	userFrontendMiddleware := middleware.CreateNewMiddlewareStack(
		middleware.FrontendAuthMiddleware,
	)

	// ============ Router stuff ============

	// ====== Main router ======
	muxMainRouter := http.NewServeMux()

	// ============ Auth stuff ============
	authMux := http.NewServeMux()
	authMux.HandleFunc("GET /authwall/login", authHandlers.HandleLogin)
	authMux.HandleFunc("GET /authwall/logout", authHandlers.HandleLogout)
	authMux.HandleFunc("POST /authwall/register", authHandlers.HandleRegistration)

	// ============ Tenant stuff ============

	tenantRouter := http.NewServeMux()

	// POST Requests
	tenantRouter.HandleFunc("POST /tenant/create", tenantHandlers.HandleTenantCreation)
	tenantRouter.HandleFunc("POST /tenant/change", tenantHandlers.HandleTenantChange)
	tenantRouter.HandleFunc("POST /tenant/delete", tenantHandlers.HandleTenantDeletion)

	// GET Requests
	tenantRouter.HandleFunc("GET /tenant/{tenant-id}", tenantHandlers.HandleGetTenantByTenantId)
	tenantRouter.HandleFunc("GET /tenant/by/user={user-id}", tenantHandlers.HandleGetTenantByUserId)

	// ==== Changing permission on tenant ====
	tenantPermissionRouter := http.NewServeMux()
	tenantPermissionRouter.HandleFunc("POST /user/permission/add", userHandlers.HandleAddingNewUserPermission)
	tenantPermissionRouter.HandleFunc("POST /user/permission/remove", userHandlers.HandleRemovingUserPermission)

	// ============ User stuff ============

	userRouter := http.NewServeMux()
	userRouter.HandleFunc("/user/change", userHandlers.HandleUserChange)
	// TODO: add more stuff here

	// ==== User frontend stuff ====
	userFrontendRouter := http.NewServeMux()

	// Normal user querys
	userFrontendRouter.HandleFunc("GET /user/by/auth={authtoken}", userHandlers.HandleGettingUserByAuthTokenVal)
	userFrontendRouter.HandleFunc("GET /user/by/id={user-id}", userHandlers.HandleGettingUserById)
	userFrontendRouter.HandleFunc("GET /user/by/tenant-id={tenant-id}&user-name={user-name}", userHandlers.HandleGettingUserByTenantIdAndUserName)
	userFrontendRouter.HandleFunc("GET /user/has-tenant", userHandlers.HandleHasUserATenant)
	userFrontendRouter.HandleFunc("GET /user/permission/all", userHandlers.HandleGettingAllPermissions)

	// User permission stuff
	userFrontendRouter.HandleFunc("GET /user/permission/all/by/user-id={user-id}", userHandlers.HandleGettingUserPermissions)
	userFrontendRouter.HandleFunc("GET /user/permission/permission-id={permission-id}", userHandlers.HandleGettingPermissionById)

	// ============ Adding all subrouters ============
	muxMainRouter.Handle("/", authMux)
	muxMainRouter.Handle("/", tenantMiddleware(tenantRouter))
	muxMainRouter.Handle("/", tenantPermissionMiddleware(tenantPermissionRouter))
	muxMainRouter.Handle("/", userMiddleware(userRouter))
	muxMainRouter.Handle("/", userFrontendMiddleware(userFrontendRouter))

	return http.Server{
		Addr:    "8080",
		Handler: globalMiddlewareConfig(muxMainRouter),
	}
}
