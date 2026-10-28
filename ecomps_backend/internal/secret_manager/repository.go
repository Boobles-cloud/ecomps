package secretmanager

import (
	"context"
	"uuid"

	smstructs "ecomps.boobles.cloud/backend/internal/secret_manager/sm_structs"
)

type SecretManagerRepository interface {
	GetTenantPw(ctx context.Context, tenantId uuid.UUID) (smstructs.SecretEntry, error)
	GetAllTenantPws(ctx context.Context) ([]smstructs.SecretEntry, error) // Only used for internal secrets cache
	GetClientSecret(ctx context.Context, providerKey string) (smstructs.SecretEntry, error)
	CreateTenantPw(ctx context.Context, tenantId uuid.UUID) error
	CheckProviderHealth(ctx context.Context) error
	DeleteTenantPw(ctx context.Context, tenantId uuid.UUID) error
}
