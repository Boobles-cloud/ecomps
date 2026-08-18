package startup

import (
	"net/http"

	"ecomps.boobles.cloud/backend/database"
	authHandlers "ecomps.boobles.cloud/backend/internal/auth/handlers"
	customerstructs "ecomps.boobles.cloud/backend/internal/customer/customer_structs"
	customerhandlers "ecomps.boobles.cloud/backend/internal/customer/handlers"
	"ecomps.boobles.cloud/backend/internal/middleware"
	orderhandlers "ecomps.boobles.cloud/backend/internal/order/handlers"
	orderstructs "ecomps.boobles.cloud/backend/internal/order/order_structs"
	producthandlers "ecomps.boobles.cloud/backend/internal/product/handlers"
	productstructs "ecomps.boobles.cloud/backend/internal/product/product_structs"
	tenanthandlers "ecomps.boobles.cloud/backend/internal/tenant/handlers"
	tenantstructs "ecomps.boobles.cloud/backend/internal/tenant/tenant_structs"
	userHandlers "ecomps.boobles.cloud/backend/internal/user/handlers"
	userstructs "ecomps.boobles.cloud/backend/internal/user/user_structs"
	"ecomps.boobles.cloud/backend/utils/caching"
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
		middleware.PermissionMiddleware(dh),
	)
	// ============ Router stuff ============

	// ====== Main router ======
	muxMainRouter := http.NewServeMux()

	// ============ Auth stuff ============
	muxMainRouter.HandleFunc("GET /authwall/login", authHandler.HandleLogin)
	muxMainRouter.HandleFunc("GET /authwall/logout", authHandler.HandleLogout)
	muxMainRouter.HandleFunc("POST /authwall/register", authHandler.HandleRegistration)

	// ============ Tenant stuff ============

	// POST Requests
	muxMainRouter.Handle("POST /tenant/create", tenantMiddleware(http.HandlerFunc(tenantHandler.HandleTenantCreation)))

	// GET Requests
	muxMainRouter.Handle("GET /tenant/{tenant_id}", tenantMiddleware(http.HandlerFunc(tenantHandler.HandleGetTenantByTenantId)))
	muxMainRouter.Handle("GET /tenant/by/{user_id}", tenantMiddleware(http.HandlerFunc(tenantHandler.HandleGetTenantByUserId)))
	muxMainRouter.Handle("GET /tenant/all/users", tenantMiddleware(http.HandlerFunc(tenantHandler.HandleGettingAllUsersByUserTenantId)))

	// ==== Changing permission on tenant ====
	muxMainRouter.Handle("POST /tenant/change", tenantPermissionMiddleware(http.HandlerFunc(tenantHandler.HandleTenantChange)))
	muxMainRouter.Handle("POST /tenant/delete", tenantPermissionMiddleware(http.HandlerFunc(tenantHandler.HandleTenantDeletion)))
	muxMainRouter.Handle("POST /user/permissions/add", tenantPermissionMiddleware(http.HandlerFunc(userHandler.HandleAddingNewUserPermission)))
	muxMainRouter.Handle("POST /user/permissions/remove", tenantPermissionMiddleware(http.HandlerFunc(userHandler.HandleRemovingUserPermission)))

	// ============ User stuff ============

	// POST Requests
	muxMainRouter.Handle("POST /user/change", userMiddleware(http.HandlerFunc(userHandler.HandleUserChange)))

	muxMainRouter.Handle("DELETE /user/deletion", userMiddleware(http.HandlerFunc(userHandler.HandleUserDeletion)))

	// ==== User frontend stuff ====

	// Normal user querys
	muxMainRouter.Handle("GET /user/frontend/by/token/{authtoken}", userFrontendMiddleware(http.HandlerFunc(userHandler.HandleGettingUserByAuthTokenVal)))
	muxMainRouter.Handle("GET /user/frontend/by/id/{user_id}", userFrontendMiddleware(http.HandlerFunc(userHandler.HandleGettingUserById)))
	muxMainRouter.Handle("GET /user/frontend/by/{tenant_id}/{user_name}", userFrontendMiddleware(http.HandlerFunc(userHandler.HandleGettingUserByTenantIdAndUserName)))
	muxMainRouter.Handle("GET /user/frontend/has-tenant/{user_id}", userFrontendMiddleware(http.HandlerFunc(userHandler.HandleHasUserATenant)))
	muxMainRouter.Handle("GET /user/frontend/permission/all", userFrontendMiddleware(http.HandlerFunc(userHandler.HandleGettingAllPermissions)))

	// User permission stuff
	muxMainRouter.Handle("GET /user/frontend/permission/all/by/{user_id}", userFrontendMiddleware(http.HandlerFunc(userHandler.HandleGettingUserPermissions)))
	muxMainRouter.Handle("GET /user/frontend/permission/{permission_id}", userFrontendMiddleware(http.HandlerFunc(userHandler.HandleGettingPermissionById)))

	// ============ Product stuff ============

	// GET Requests
	muxMainRouter.Handle("GET /product/by/{product_id}", productMiddleware(http.HandlerFunc(productHandler.HandleGettingProductById)))
	muxMainRouter.Handle("GET /product/all", productMiddleware(http.HandlerFunc(productHandler.HandleGettingAllProductsByTenantId)))
	muxMainRouter.Handle("GET /product/picture/by/{product_id}", productMiddleware(http.HandlerFunc(productHandler.HandleGettingPictureByProductId)))

	// POST Requests
	muxMainRouter.Handle("POST /product/create", productMiddleware(http.HandlerFunc(productHandler.HandleCreatingProduct)))
	muxMainRouter.Handle("POST /product/picture/create", productMiddleware(http.HandlerFunc(productHandler.HandleCreatingProductPicture)))
	muxMainRouter.Handle("POST /product/change", productMiddleware(http.HandlerFunc(productHandler.HandleChangingProduct)))

	// DELETE Requests
	muxMainRouter.Handle("DELETE /product/delete/by/{product_id}", productMiddleware(http.HandlerFunc(productHandler.HandleDeletingProduct)))

	// ============ Order stuff ============

	// GET Requests
	muxMainRouter.Handle("GET /order/by/{order_id}", orderMiddleware(http.HandlerFunc(orderHandler.HandleGettingOrderById)))
	muxMainRouter.Handle("GET /order/all", orderMiddleware(http.HandlerFunc(orderHandler.HandleGettingAllOrdersByTenantId)))
	muxMainRouter.Handle("GET /order/status/by/{status_id}/{language_id}", orderMiddleware(http.HandlerFunc(orderHandler.HandleGettingStatusById)))
	muxMainRouter.Handle("GET /order/status/{language_id}", orderMiddleware(http.HandlerFunc(orderHandler.HandleGettingStatusById)))

	// POST Requests
	muxMainRouter.Handle("POST /order/create", orderMiddleware(http.HandlerFunc(orderHandler.HandleCreatingOrder)))
	muxMainRouter.Handle("POST /order/change", orderMiddleware(http.HandlerFunc(orderHandler.HandleChangingOrder)))

	// DELETE Requests
	muxMainRouter.Handle("DELETE /order/delete/by/{order_id}", orderMiddleware(http.HandlerFunc(orderHandler.HandleOrderDeletion)))

	// ============ Customer stuff ============

	// GET Requests
	muxMainRouter.Handle("GET /customer/by/{customer_id}", customerMiddleware(http.HandlerFunc(customerHandler.HandleGettingCustomerById)))
	muxMainRouter.Handle("GET /customer/all", customerMiddleware(http.HandlerFunc(customerHandler.HandleGettingAllCustomerByTenantId)))

	// POST Requests
	muxMainRouter.Handle("POST /customer/create", customerMiddleware(http.HandlerFunc(customerHandler.HandleCustomerCreation)))
	muxMainRouter.Handle("POST /customer/change", customerMiddleware(http.HandlerFunc(customerHandler.HandleCustomerChange)))

	// DELETE Requests
	muxMainRouter.Handle("DELETE /customer/delete/by/{customer_id}", customerMiddleware(http.HandlerFunc(customerHandler.HandleCustomerDeletion)))

	return http.Server{
		Addr:    ":8080",
		Handler: globalMiddlewareConfig(muxMainRouter),
	}
}
