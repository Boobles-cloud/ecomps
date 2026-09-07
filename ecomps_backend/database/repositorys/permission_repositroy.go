package repositorys

import (
	"context"
	"database/sql"
	"errors"

	"ecomps.boobles.cloud/backend/database"
	permssionstructs "ecomps.boobles.cloud/backend/internal/permission/permission_structs"
)

type PermissionRepository struct {
	*DatabaseRepository[permssionstructs.Permission]
}

// Creates a new user permission repository
func NewPermissionRepository(db *database.DbHandler, entityName string, toArgs ToArgsFunc[permssionstructs.Permission]) *PermissionRepository {
	return &PermissionRepository{
		DatabaseRepository: NewDatabaseRepository(db, entityName, toArgs),
	}
}

func (p *PermissionRepository) AsingUserPermission(ctx context.Context, userId, permissionId uint) error {

	if result := p.db.ExecuteSQLStatement(ctx, "InsertUserPermission", []any{permissionId, userId}); !result.Ok {
		return errors.New("Failed to add permission to user")
	}

	return nil
}

func (p *PermissionRepository) GetAllPermissionsByLanguageId(ctx context.Context, langId uint) ([]permssionstructs.Permission, error) {

	results, ok := database.QueryMany[permssionstructs.Permission](ctx, p.db, "SelectPermissionsByLanguageId", []any{langId})

	if !ok {
		return []permssionstructs.Permission{}, sql.ErrNoRows
	}

	return results, nil
}

func (p *PermissionRepository) GetAllPermissionsForUserId(ctx context.Context, userId uint) ([]permssionstructs.Permission, error) {

	results, ok := database.QueryMany[permssionstructs.Permission](ctx, p.db, "SelectPermissionsByUserId", []any{userId})

	if !ok {
		return []permssionstructs.Permission{}, sql.ErrNoRows
	}

	return results, nil
}

func (p *PermissionRepository) RemoveUserPermission(ctx context.Context, userId, permissionId uint) error {

	if result := p.db.ExecuteSQLStatement(ctx, "DeletePermissionFromUserByUserIdAndPermissionId", []any{userId, permissionId}); !result.Ok {
		return errors.New("Failed to delete permission")
	}
	return nil
}
