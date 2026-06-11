package tenantstructs

import (
	"context"
	"time"

	"boobles.cloud/backend/database"
	userstructs "boobles.cloud/backend/internal/user/user_structs"
	"boobles.cloud/backend/logging"
)

type Tenant struct {
	TenantId          uint      `json:"TenantId"`
	TenantName        string    `json:"TenantName"`
	TenantAdminUserId uint      `json:"TenantAdminUser"`
	TenantCreation    time.Time `json:"-"`
	TenantPwId        uint      `json:"-"`
	// TODO: add more stuff here and also in the tables.sql
}

// TODO: Refactor this!!
// Creates a tenant in the database
func (t *Tenant) CreateTenantInDatabase() bool {

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	t.TenantCreation = time.Now()

	// Creates a connection
	db, ok := database.CreateDBConn()

	if !ok {
		return ok
	}

	// Always close connection!
	defer db.Close()

	// Creates our transaction stuff
	tx, err := db.BeginTx(ctx, nil)

	if err != nil {
		logging.Log(logging.Error, "[Tenant | CreateTenantInDatabase] "+err.Error())
		return false
	}

	// Rollback on error
	defer tx.Rollback()

	// Create our master key and insert it
	id, ok := createMasterKey(*t)

	if !ok {
		logging.Log(logging.Error, "[Tenant | CreateTenantInDatabase] Failed to create master key!")
		return false
	}

	// Sets the id from the given master key column
	t.TenantPwId = id

	result, err := tx.ExecContext(ctx, "INSERT INTO Tenant() VALUES(DEFAULT, ?, ?, ?, ?)", []any{t.TenantName, t.TenantAdminUserId, t.TenantCreation, t.TenantPwId})

	lastId, _ := result.LastInsertId()
	// Set the last Id here
	t.TenantId = uint(lastId)

	if err != nil {
		logging.Log(logging.Error, "[Tenant | CreateTenantInDatabase] "+err.Error())
		return false
	}

	// Get the wanted User
	user, ok := database.QueryDatabase[userstructs.UserStruct]("SelectUserById", []any{t.TenantId})

	// Check that its only one user and its ok
	if !ok || len(user) > 1 {
		return ok
	}

	// Set the tenant id
	user[0].TenantId = uint(lastId)

	// Update our User
	if _, err := tx.ExecContext(ctx, "UPDATE Users VALUES TenantId = ? WHERE UserId = ?", []any{user[0].TenantId, user[0].UserId}); err != nil {
		logging.Log(logging.Error, "[Tenant | CreateTenantInDatabase] "+err.Error())
		return false
	}

	// Commits our transaction
	if err := tx.Commit(); err != nil {
		logging.Log(logging.Error, "[Tenant | CreateTenantInDatabase] "+err.Error())
		return false
	}

	// Everything is awesomeeeee!!!!
	return true
}

// Checks if a user is the admin of the tenant
func (t *Tenant) IsUserAdmin(u *userstructs.UserStruct) bool {

	if t.TenantAdminUserId == u.UserId {
		return true
	}
	return false
}

// Gets all Users for the given tenant
func (t *Tenant) GetAllUsersForTenant() []userstructs.UserStruct {

	users, _ := database.QueryDatabase[userstructs.UserStruct]("SelectAllUsersByTenant", []any{t.TenantId})

	return users
}
