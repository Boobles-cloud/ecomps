package repositorys

import (
	"context"
	"database/sql"
	"errors"
	"math/rand"
	"time"

	"ecomps.boobles.cloud/backend/database"
	tenantstructs "ecomps.boobles.cloud/backend/internal/tenant/tenant_structs"
)

const (
	allCharacters  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	maxLengthOfKey = 32
)

type TenantRepository struct {
	*DatabaseRepository[tenantstructs.Tenant]
}

func NewTenantRepository(db *database.DbHandler, entityName string, toArgs ToArgsFunc[tenantstructs.Tenant]) *TenantRepository {
	return &TenantRepository{
		DatabaseRepository: NewDatabaseRepository(db, "Tenant", toArgs),
	}
}

func (t *TenantRepository) GetById(ctx context.Context, id uint) (tenantstructs.Tenant, error) {

	tenant, ok := database.QueryOne[tenantstructs.Tenant](ctx, t.db, "SelectTenantById", []any{id})

	if !ok {
		return tenant, sql.ErrNoRows
	}

	return tenant, nil
}

func (t *TenantRepository) GetPw(ctx context.Context, tenantId uint) (string, error) {

	tenant, ok := database.QueryOne[tenantstructs.TenantPwStruct](ctx, t.db, "SelectTenantPwByTenantId", []any{tenantId})

	if !ok {
		return "", sql.ErrNoRows
	}

	return tenant.TenantPwVal, nil
}

func (t *TenantRepository) Create(ctx context.Context, tenant tenantstructs.Tenant, userId uint) error {
	tx, err := t.db.DbConnection.Begin()

	if err != nil {
		return err
	}

	defer tx.Rollback()

	var userHasTenant bool

	err = tx.QueryRowContext(ctx,
		"SELECT UserHasTenant FROM Users WHERE UserId = ? FOR UPDATE", userId,
	).Scan(&userHasTenant)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return errors.New("User has no tenant")
	} else if err != nil {
		return err
	}

	if userHasTenant {
		return errors.New("User already has a tenant")
	}

	masterKeyId, ok := createMasterKey(ctx, tx, t.DatabaseRepository.db)

	if !ok {
		return errors.New("Failed to create tenant key")
	}

	tenant.TenantPwId = masterKeyId

	insertQuery := "INSERT INTO Tenant (TenantName, TenantCreation, TenantAdminUserId, TenantPwId) VALUES (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, insertQuery, tenant.TenantName, tenant.TenantCreation, tenant.TenantAdminUserId, tenant.TenantPwId)

	if err != nil {
		return err
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return err
	}

	tenant.TenantId = uint(lastId)

	updateQuery := "UPDATE Users SET TenantId = ?, UserHasTenant = TRUE WHERE UserId = ?"
	if _, err := tx.ExecContext(ctx, updateQuery, lastId, userId); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (t *TenantRepository) Update(ctx context.Context, tenant tenantstructs.Tenant) error {

	if result := database.UpdateDatabaseEntry[tenantstructs.Tenant](t.db, "UpdateTenant", "TenantId", tenant); !result {
		return errors.New("Failed to update")
	}

	return nil
}

// This func adds a tenant to the delete database
func (t *TenantRepository) Delete(ctx context.Context, userId, tenantId uint) error {

	tenantDelete := tenantstructs.TenantDeletionStruct{
		IssuedFrom:     userId,
		IssuedOn:       time.Now(),
		WhenToComplete: time.Now().AddDate(0, 2, 0),
		Deleted:        false,
		TenantId:       tenantId,
	}

	if result := t.db.ExecuteSQLStatement(ctx, "InsertTenantDeletion", []any{tenantDelete.IssuedFrom,
		tenantDelete.IssuedOn, tenantDelete.WhenToComplete, tenantDelete.Deleted, tenantDelete.TenantId}); !result.Ok {
		return errors.New("Failed to insert ")
	}

	return nil
}

// Creates the master key for the given tenant, inside the given transaction.
func createMasterKey(ctx context.Context, tx *sql.Tx, dh *database.DbHandler) (uint, bool) {
	tenantPw := createRandomString(maxLengthOfKey)
	result := dh.ExecuteSQLStatementTx(ctx, tx, "InsertTenantPw", []any{tenantPw})
	return result.LastId, result.Ok
}

// Create a random string of the given length.
func createRandomString(length int) string {
	randSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randSource)

	result := make([]byte, length)
	for i := range length {
		result[i] = allCharacters[random.Intn(len(allCharacters))]
	}
	return string(result)
}
