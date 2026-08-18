package tenantstructs

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"ecomps.boobles.cloud/backend/database"
	"ecomps.boobles.cloud/backend/logging"
)

type Tenant struct {
	TenantId          uint      `json:"TenantId"`
	TenantName        string    `json:"TenantName"`
	TenantAdminUserId uint      `json:"TenantAdminUser"`
	TenantCreation    time.Time `json:"-"`
	TenantPwId        uint      `json:"-"`
}

// Creates a tenant and assigns the given user to it.
// Fails if the user doesn't exist or already belongs to a tenant.
func (t *Tenant) CreateTenantInDatabase(ctx context.Context, userId int, dh *database.DbHandler) bool {

	tx, err := dh.DbConnection.BeginTx(ctx, nil)
	if err != nil {
		logging.Log(logging.Error, "[Tenant | CreateTenantInDatabase] "+err.Error())
		return false
	}
	defer tx.Rollback()

	var userHasTenant bool
	err = tx.QueryRowContext(ctx,
		"SELECT UserHasTenant FROM Users WHERE UserId = ? FOR UPDATE", userId,
	).Scan(&userHasTenant)

	if err != nil {
		if err == sql.ErrNoRows {
			logging.Log(logging.Error, fmt.Sprintf("[Tenant | CreateTenantInDatabase] User with ID %d not found", userId))
		} else {
			logging.Log(logging.Error, "[Tenant | CreateTenantInDatabase] "+err.Error())
		}
		return false
	}
	if userHasTenant {
		logging.Log(logging.Error, fmt.Sprintf("[Tenant | CreateTenantInDatabase] User with ID %d already has a tenant", userId))
		return false
	}

	t.TenantCreation = time.Now()

	masterKeyID, ok := createMasterKey(ctx, tx, dh, *t)
	if !ok {
		logging.Log(logging.Error, "[Tenant | CreateTenantInDatabase] Failed to create master key!")
		return false
	}
	t.TenantPwId = masterKeyID

	insertQuery := "INSERT INTO Tenant (TenantName, TenantAdminUserId, TenantCreation, TenantPwId) VALUES (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, insertQuery, t.TenantName, t.TenantAdminUserId, t.TenantCreation, t.TenantPwId)
	if err != nil {
		logging.Log(logging.Error, "[Tenant | CreateTenantInDatabase] "+err.Error())
		return false
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		logging.Log(logging.Error, "[Tenant | CreateTenantInDatabase] Failed to get LastInsertId: "+err.Error())
		return false
	}
	t.TenantId = uint(lastId)

	updateQuery := "UPDATE Users SET TenantId = ?, UserHasTenant = TRUE WHERE UserId = ?"
	if _, err := tx.ExecContext(ctx, updateQuery, lastId, userId); err != nil {
		logging.Log(logging.Error, "[Tenant | CreateTenantInDatabase] "+err.Error())
		return false
	}

	if err := tx.Commit(); err != nil {
		logging.Log(logging.Error, "[Tenant | CreateTenantInDatabase] "+err.Error())
		return false
	}

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
	tp, ok := database.QueryOne[TenantPwStruct](ctx, dh, "SelectTenantPwByTenantId", t.TenantPwId)

	if !ok {
		logging.Log(logging.Error, "Got more then one tenant pw...")
		return ""
	}

	return tp.TenantPwVal
}
