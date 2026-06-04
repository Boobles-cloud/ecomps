package tenantstructs

import (
	"time"

	"boobles.cloud/backend/database"
	userstructs "boobles.cloud/backend/internal/user/user_structs"
)

type Tenant struct {
	TenantId          uint      `json:"TenantId"`
	TenantName        string    `json:"TenantName"`
	TenantPw          string    `json:"-"`
	TenantAdminUserId uint      `json:"TenantAdminUser"`
	TenantCreation    time.Time `json:"-"`
	// TODO: add more stuff here and also in the tables.sql
}

// Creates a tenant in the database
func (t *Tenant) CreateTenantInDatabase() bool {
	t.TenantCreation = time.Now()

	result := database.ExecuteSQLStatement("InsertTenant", database.Insert, []any{})

	database.QueryDatabase[userstructs.UserStruct]("", []any{t.TenantAdminUserId})

	// TODO: Create master PW and then insert it into the database
	// TODO: Also update the admin user -> so he has the tenant id
	return result.Ok
}

// Checks if a user is the admin of the tenant
func (t *Tenant) IsUserAdmin(u *userstructs.UserStruct) bool

// Gets all Users for the given tenant
func (t *Tenant) GetAllUsersForTenant() []userstructs.UserStruct
