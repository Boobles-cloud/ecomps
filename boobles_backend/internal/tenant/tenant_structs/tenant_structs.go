package tenantstructs

import (
	"context"
	"fmt"
	"time"

	"boobles.cloud/backend/database"
	"boobles.cloud/backend/logging"
)

type Tenant struct {
	TenantId          uint      `json:"TenantId"`
	TenantName        string    `json:"TenantName"`
	TenantAdminUserId uint      `json:"TenantAdminUser"`
	TenantCreation    time.Time `json:"-"`
	TenantPwId        uint      `json:"-"`
}

// Creates a tenant and updates the user
func (t *Tenant) CreateTenantInDatabase(ctx context.Context, userId int, dh *database.DbHandler) bool {

	tx, err := dh.DbConnection.BeginTx(ctx, nil)
	if err != nil {
		logging.Log(logging.Error, "[Tenant | CreateTenantInDatabase] "+err.Error())
		return false
	}
	defer tx.Rollback()

	t.TenantCreation = time.Now()

	masterKeyID, ok := createMasterKey(*t, dh)
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

	var userExists bool
	err = tx.QueryRowContext(ctx, "SELECT EXISTS(SELECT * FROM Users WHERE UserId = ?)", userId).Scan(&userExists)
	if err != nil || !userExists {
		logging.Log(logging.Error, fmt.Sprintf("[Tenant | CreateTenantInDatabase] User with ID %d not found or query failed", userId))
		return false
	}

	updateQuery := "UPDATE Users SET TenantId = ? WHERE UserId = ?"
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
