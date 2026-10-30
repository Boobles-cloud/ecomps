package repositorys

import (
	"context"
	"errors"
	"uuid"

	smstructs "ecomps.boobles.cloud/backend/internal/secret_manager/sm_structs"
	vault "github.com/hashicorp/vault/api"
)

type StackitRepository struct {
	client                  *vault.Client
	defaultTenantPath       string
	defaultClientSecretPath string
}

func CreateNewStackitRepository(client *vault.Client) *StackitRepository {
	return &StackitRepository{
		client: client,
	}
}

func (s *StackitRepository) GetTenantPw(ctx context.Context, tenantId uuid.UUID) (smstructs.SecretEntry, error) {

	secret, err := s.client.Logical().Read(s.defaultTenantPath + tenantId.String())

	if err != nil {
		return smstructs.SecretEntry{}, err
	}

	if secret == nil || secret.Data["data"] == nil {
		return smstructs.SecretEntry{}, errors.New("Path is empty")
	}

	datamap, ok := secret.Data["data"].(map[string]any)

	if !ok {
		return smstructs.SecretEntry{}, errors.New("Failed to extract secret")
	}

	return smstructs.SecretEntry{
		Key:   tenantId.String(),
		Value: datamap["TenantPw"].(string),
	}, nil
}

func (s *StackitRepository) GetAllTenantPws(ctx context.Context) ([]smstructs.SecretEntry, error) {
	return []smstructs.SecretEntry{}, errors.ErrUnsupported
}

func (s *StackitRepository) GetClientSecret(ctx context.Context, providerKey string) (smstructs.SecretEntry, error) {
	return smstructs.SecretEntry{}, errors.ErrUnsupported
}

func (s *StackitRepository) CreateTenantPw(ctx context.Context, tenantId uuid.UUID) error {
	return errors.ErrUnsupported
}

func (s *StackitRepository) CheckProviderHealth(ctx context.Context) error {
	return errors.ErrUnsupported
}

func (s *StackitRepository) DeleteTenantPw(ctx context.Context, tenantId uuid.UUID) error {
	return errors.ErrUnsupported
}
