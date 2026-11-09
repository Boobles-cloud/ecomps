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
	instanceId              string
}

func CreateNewStackitRepository(client *vault.Client, tenantPath, clientSecretPath, instanceId string) *StackitRepository {
	return &StackitRepository{
		client:                  client,
		defaultTenantPath:       tenantPath,
		defaultClientSecretPath: clientSecretPath,
		instanceId:              instanceId,
	}
}

func (s *StackitRepository) GetTenantPw(ctx context.Context, tenantId uuid.UUID) (smstructs.SecretEntry, error) {

	secret, err := s.client.KVv2(s.instanceId).Get(ctx, tenantId.String())

	if err != nil {
		return smstructs.SecretEntry{}, err
	}

	dataMap, ok := secret.Data["tenantPw"].(string)

	if !ok {
		return smstructs.SecretEntry{}, errors.New("Value not found")
	}

	return smstructs.SecretEntry{
		Key:   tenantId.String(),
		Value: dataMap,
	}, nil
}

func (s *StackitRepository) GetAllTenantPws(ctx context.Context) ([]smstructs.SecretEntry, error) {

	// TODO: implement meeee!!

	return []smstructs.SecretEntry{}, errors.ErrUnsupported
}

func (s *StackitRepository) GetClientSecret(ctx context.Context, providerKey string) (smstructs.SecretEntry, error) {

	secret, err := s.client.KVv2(s.instanceId).Get(ctx, "Test123")

	if err != nil {
		return smstructs.SecretEntry{}, err
	}

	dataMap, ok := secret.Data["client_secret"].(string)

	if !ok {
		return smstructs.SecretEntry{}, errors.New("Key not found")
	}

	return smstructs.SecretEntry{
		Key:   providerKey,
		Value: dataMap,
	}, nil
}

func (s *StackitRepository) GetClientId(ctx context.Context, providerKey string) (smstructs.SecretEntry, error) {
	secret, err := s.client.KVv2(s.instanceId).Get(ctx, "Test123")

	if err != nil {
		return smstructs.SecretEntry{}, err
	}

	dataMap, ok := secret.Data["client_id"].(string)

	if !ok {
		return smstructs.SecretEntry{}, errors.New("Key not found")
	}

	return smstructs.SecretEntry{
		Key:   providerKey,
		Value: dataMap,
	}, nil
}

func (s *StackitRepository) CreateTenantPw(ctx context.Context, tenantId uuid.UUID, pw string) error {

	// TODO: implement mee!!

	return errors.ErrUnsupported
}

func (s *StackitRepository) CheckProviderHealth(ctx context.Context) error {

	health, err := s.client.Sys().Health()

	if err != nil {
		return err
	}

	if !health.Initialized {
		return errors.New("Client isnt initialized!")
	}

	return nil
}

func (s *StackitRepository) DeleteTenantPw(ctx context.Context, tenantId uuid.UUID) error {

	// TODO: implement mee!!

	return errors.ErrUnsupported
}
