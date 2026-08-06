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
}

// TODO: Refactor this!!
// Creates a tenant in the database
func (t *Tenant) CreateTenantInDatabase(userId int, dh *database.DbHandler) bool {

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	t.TenantCreation = time.Now()

	// Creates our transaction stuff
	tx, err := dh.DbConnection.BeginTx(ctx, nil)

	if err != nil {
		logging.Log(logging.Error, "[Tenant | CreateTenantInDatabase] "+err.Error())
		return false
	}

	// Rollback on error
	defer tx.Rollback()

	// Create our master key and insert it
	id, ok := createMasterKey(*t, dh)

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
	user, ok := database.QueryOne[userstructs.UserStruct](ctx, dh, "SelectUserById", []any{userId})

	// Check that its only one user and its ok
	if !ok {
		return ok
	}

	// Set the tenant id
	user.TenantId = uint(lastId)

	// Update our User
	if _, err := tx.ExecContext(ctx, "UPDATE Users VALUES TenantId = ? WHERE UserId = ?", []any{user.TenantId, user.UserId}); err != nil {
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
func (t *Tenant) IsUserAdmin(userId uint) bool {

	if t.TenantAdminUserId == userId {
		return true
	}
	return false
}

// Gets the tenant pw
func (t *Tenant) GetPw(dh *database.DbHandler, ctx context.Context) string {
	tp, ok := database.QueryOne[TenantPwStruct](ctx, dh, "SelectTenantPwByTenantId", []any{t.TenantPwId})

	if !ok {
		logging.Log(logging.Error, "Got more then one tenant pw...")
		return ""
	}

	return tp.TenantPwVal
}
