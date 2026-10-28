package repositorys

import (
	"context"
	"database/sql"
	"errors"
	"time"
	"uuid"

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

func (t *TenantRepository) GetById(ctx context.Context, id uuid.UUID) (tenantstructs.Tenant, error) {

	tenant, ok := database.QueryOne[tenantstructs.Tenant](ctx, t.db, "SelectTenantById", []any{id})

	if !ok {
		return tenant, sql.ErrNoRows
	}

	return tenant, nil
}

func (t *TenantRepository) Create(ctx context.Context, tenant *tenantstructs.Tenant, userId uint) error {

	var userHasTenant, userDisabled bool

	err := t.db.DbConnection.QueryRowContext(ctx,
		"SELECT UserHasTenant, UserDisabled FROM Users WHERE UserId = ? FOR UPDATE", userId,
	).Scan(&userHasTenant, &userDisabled)

	if err != nil {
		return err
	}

	if userHasTenant || userDisabled {
		return errors.New("User is disabled or has a tenant")
	}

	if result := t.db.ExecuteSQLStatement(ctx, "InsertTenant", t.toArgs(*tenant)); !result.Ok {
		return errors.New("Failed to insert tenant")
	}

	// This is a bit messy, but who cares. Right?
	if _, err := t.db.DbConnection.Exec("UPDATE Users SET TenantId = ?, UserHasTenant = TRUE WHERE UserId = ?", []any{tenant.TenantId, userId}); err != nil {
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
func (t *TenantRepository) Delete(ctx context.Context, userId uint, tenantId uuid.UUID) error {

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

// Gets a deletion struct
func (t *TenantRepository) GetTenantDeletion(ctx context.Context, id uuid.UUID) (tenantstructs.TenantDeletionStruct, error) {

	deletion, ok := database.QueryOne[tenantstructs.TenantDeletionStruct](ctx, t.db, "SelectTenantDeletionFromTenantId", []any{id})

	if !ok {
		return deletion, sql.ErrNoRows
	}

	return deletion, nil
}
