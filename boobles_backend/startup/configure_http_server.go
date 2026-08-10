package startup

import (
	"net/http"

	"boobles.cloud/backend/caching"
	"boobles.cloud/backend/database"
	authHandlers "boobles.cloud/backend/internal/auth/handlers"
	customerstructs "boobles.cloud/backend/internal/customer/customer_structs"
	customerhandlers "boobles.cloud/backend/internal/customer/handlers"
	"boobles.cloud/backend/internal/middleware"
	orderhandlers "boobles.cloud/backend/internal/order/handlers"
	orderstructs "boobles.cloud/backend/internal/order/order_structs"
	producthandlers "boobles.cloud/backend/internal/product/handlers"
	productstructs "boobles.cloud/backend/internal/product/product_structs"
	tenanthandlers "boobles.cloud/backend/internal/tenant/handlers"
	tenantstructs "boobles.cloud/backend/internal/tenant/tenant_structs"
	userHandlers "boobles.cloud/backend/internal/user/handlers"
	userstructs "boobles.cloud/backend/internal/user/user_structs"
)

// Creates and configures the rest api
func ConfigureHTTPServer(dh *database.DbHandler) http.Server {

	// ============ Cache config stuff ============

	// Auth handler stuff
	authHandler := authHandlers.CreateAuthHandler(dh)

	// User cache config
	userCache := caching.CreateNewCacheManager[userstructs.UserStruct]()
	permissionCache := caching.CreateNewCacheManager[userstructs.UserPermission]()
	userHandler := userHandlers.CreateNewUserHander(userCache, permissionCache, dh)

	// Tenant cache config
	tenantCache := caching.CreateNewCacheManager[tenantstructs.Tenant]()
	tenantHandler := tenanthandlers.CreateNewUserHander(tenantCache, dh)

	// Product cache config
	productCache := caching.CreateNewCacheManager[productstructs.Product]()
	productHandler := producthandlers.CreateNewProductHandler(productCache, dh)

	// Order cache config
	orderCache := caching.CreateNewCacheManager[orderstructs.Order]()
	orderHandler := orderhandlers.CreateNewOrderHandler(orderCache, dh)

	// Customer cache config
	customerCache := caching.CreateNewCacheManager[customerstructs.Customer]()
	customerHandler := customerhandlers.CreateNewCustomerHandler(customerCache, dh)

	// ============ Middleware config stuff ============

	// Configuring our global middleware
	globalMiddlewareConfig := middleware.CreateNewMiddlewareStack(
		middleware.LoggingMiddleware,
		middleware.PanicRecoverMiddleware,
	)

	// Tenant middleware config
	tenantMiddleware := middleware.CreateNewMiddlewareStack(
		middleware.AuthMiddleware(dh),
		middleware.PermissionMiddleware(dh),
	)

	// Tenant permission changes -> this can only be done by the admin
	tenantPermissionMiddleware := middleware.CreateNewMiddlewareStack(
		middleware.AuthMiddleware(dh),
		middleware.CheckAdminMiddleware(dh),
	)

	// For changes on the user config itself
	userMiddleware := middleware.CreateNewMiddlewareStack(
		middleware.AuthMiddleware(dh),
	)

	// For the frontend to get user information
	userFrontendMiddleware := middleware.CreateNewMiddlewareStack(
		middleware.FrontendAuthMiddleware,
	)

	// For the product stuff
	productMiddleware := middleware.CreateNewMiddlewareStack(
		middleware.AuthMiddleware(dh),
		middleware.PermissionMiddleware(dh),
	)

	// For the order stuff
	orderMiddleware := middleware.CreateNewMiddlewareStack(
		middleware.AuthMiddleware(dh),
		middleware.PermissionMiddleware(dh),
	)

	// For customer stuff
	customerMiddleware := middleware.CreateNewMiddlewareStack(
		middleware.AuthMiddleware(dh),
		middleware.AuthMiddleware(dh),
	)
	// ============ Router stuff ============

	// ====== Main router ======
	muxMainRouter := http.NewServeMux()

	// ============ Auth stuff ============
	authMux := http.NewServeMux()
	authMux.HandleFunc("GET /authwall/login", authHandler.HandleLogin)
	authMux.HandleFunc("GET /authwall/logout", authHandler.HandleLogout)
	authMux.HandleFunc("POST /authwall/register", authHandler.HandleRegistration)

	// ============ Tenant stuff ============

	tenantRouter := http.NewServeMux()

	// POST Requests
	tenantRouter.HandleFunc("POST /tenant/create", tenantHandler.HandleTenantCreation)
	tenantRouter.HandleFunc("POST /tenant/change", tenantHandler.HandleTenantChange)
	tenantRouter.HandleFunc("POST /tenant/delete", tenantHandler.HandleTenantDeletion)

	// GET Requests
	tenantRouter.HandleFunc("GET /tenant/{tenant-id}", tenantHandler.HandleGetTenantByTenantId)
	tenantRouter.HandleFunc("GET /tenant/by/user={user-id}", tenantHandler.HandleGetTenantByUserId)
	tenantRouter.HandleFunc("GET /tenant/all/users", tenantHandler.HandleGettingAllUsersByUserTenantId)

	// ==== Changing permission on tenant ====
	tenantPermissionRouter := http.NewServeMux()
	tenantPermissionRouter.HandleFunc("POST /user/permission/add", userHandler.HandleAddingNewUserPermission)
	tenantPermissionRouter.HandleFunc("POST /user/permission/remove", userHandler.HandleRemovingUserPermission)

	// ============ User stuff ============

	userRouter := http.NewServeMux()

	// POST Requests
	userRouter.HandleFunc("POST /user/change", userHandler.HandleUserChange)

	userRouter.HandleFunc("DELETE /user/deletion", userHandler.HandleUserChange)

	// ==== User frontend stuff ====
	userFrontendRouter := http.NewServeMux()

	// Normal user querys
	userFrontendRouter.HandleFunc("GET /user/by/auth={authtoken}", userHandler.HandleGettingUserByAuthTokenVal)
	userFrontendRouter.HandleFunc("GET /user/by/id", userHandler.HandleGettingUserById)
	userFrontendRouter.HandleFunc("GET /user/by/tenant-id={tenant-id}&user-name={user-name}", userHandler.HandleGettingUserByTenantIdAndUserName)
	userFrontendRouter.HandleFunc("GET /user/has-tenant", userHandler.HandleHasUserATenant)
	userFrontendRouter.HandleFunc("GET /user/permission/all", userHandler.HandleGettingAllPermissions)

	// User permission stuff
	userFrontendRouter.HandleFunc("GET /user/permission/all/by/user-id={user-id}", userHandler.HandleGettingUserPermissions)
	userFrontendRouter.HandleFunc("GET /user/permission/permission-id={permission-id}", userHandler.HandleGettingPermissionById)

	// ============ Product stuff ============

	productRouter := http.NewServeMux()

	// GET Requests
	productRouter.HandleFunc("GET /product/by/product-id={product-id}", productHandler.HandleGettingProductById)
	productRouter.HandleFunc("GET /product/all", productHandler.HandleGettingAllProductsByTenantId)
	productRouter.HandleFunc("GET /product/picture/by/product-id={product-id}", productHandler.HandleGettingPictureByProductId)

	// POST Requests
	productRouter.HandleFunc("POST /product/create", productHandler.HandleCreatingProduct)
	productRouter.HandleFunc("POST /product/picture/create", productHandler.HandleCreatingProductPicture)
	productRouter.HandleFunc("POST /product/change", productHandler.HandleChangingProduct)

	// DELETE Requests
	productRouter.HandleFunc("DELETE /product/delete/by/product-id={product-id}", productHandler.HandleDeletingProduct)

	// ============ Order stuff ============

	orderRouter := http.NewServeMux()

	// GET Requests
	orderRouter.HandleFunc("GET /order/by/order-id={order-id}", orderHandler.HandleGettingOrderById)
	orderRouter.HandleFunc("GET /order/all", orderHandler.HandleGettingAllOrdersByTenantId)
	orderRouter.HandleFunc("GET /order/status/by/status-id={status-id}&language-id={language-id}", orderHandler.HandleGettingStatusById)
	orderRouter.HandleFunc("GET /order/status/all={language-id}", orderHandler.HandleGettingStatusById)

	// POST Requests
	orderRouter.HandleFunc("POST /order/create", orderHandler.HandleCreatingOrder)
	orderRouter.HandleFunc("POST /order/change", orderHandler.HandleChangingOrder)

	// DELETE Requests
	orderRouter.HandleFunc("DELETE /order/delete/by/order-id={order-id}", orderHandler.HandleOrderDeletion)

	// ============ Customer stuff ============

	customerRouter := http.NewServeMux()

	// GET Requests
	customerRouter.HandleFunc("GET /customer/by/customer-id={customer-id}", customerHandler.HandleGettingCustomerById)
	customerRouter.HandleFunc("GET /customer/all", customerHandler.HandleGettingAllCustomerByTenantId)

	// POST Requests
	customerRouter.HandleFunc("POST /customer/create", customerHandler.HandleCustomerCreation)
	customerRouter.HandleFunc("POST /customer/change", customerHandler.HandleCustomerChange)

	// DELETE Requests
	customerRouter.HandleFunc("DELETE /customer/delete/by/customer-id={customer-id}", customerHandler.HandleCustomerDeletion)

	// ============ Adding all subrouters ============
	muxMainRouter.Handle("/", authMux)
	muxMainRouter.Handle("/", tenantMiddleware(tenantRouter))
	muxMainRouter.Handle("/", tenantPermissionMiddleware(tenantPermissionRouter))
	muxMainRouter.Handle("/", userMiddleware(userRouter))
	muxMainRouter.Handle("/", userFrontendMiddleware(userFrontendRouter))
	muxMainRouter.Handle("/", productMiddleware(productRouter))
	muxMainRouter.Handle("/", orderMiddleware(orderRouter))
	muxMainRouter.Handle("/", customerMiddleware(customerRouter))

	return http.Server{
		Addr:    "8080",
		Handler: globalMiddlewareConfig(muxMainRouter),
	}
}
