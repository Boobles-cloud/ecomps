package repositorys

import (
	"context"
	"errors"
	"strings"
	"uuid"

	repository "ecomps.boobles.cloud/backend/internal/secret_manager"
	smstructs "ecomps.boobles.cloud/backend/internal/secret_manager/sm_structs"
	"ecomps.boobles.cloud/backend/utils/logging"
	vault "github.com/hashicorp/vault/api"
)

type StackitRepository struct {
	client                  *vault.Client
	defaultTenantPath       string
	defaultClientSecretPath string
	defaultClientIdPath     string
	instanceId              string
}

func CreateNewStackitRepository(client *vault.Client, tenantPath, clientSecretPath, clientId, instanceId string) *StackitRepository {
	return &StackitRepository{
		client:                  client,
		defaultTenantPath:       tenantPath,
		defaultClientSecretPath: clientSecretPath,
		defaultClientIdPath:     clientId,
		instanceId:              instanceId,
	}
}

func (s *StackitRepository) GetTenantPw(ctx context.Context, tenantId uuid.UUID) (smstructs.SecretEntry, error) {

	secret, err := s.client.KVv2(s.instanceId).Get(ctx, tenantId.String())

	if err != nil {
		return smstructs.SecretEntry{}, err
	}

	dataMap, ok := secret.Data[repository.TENANT_PW_KEY].(string)

	if !ok {
		return smstructs.SecretEntry{}, errors.New("Value not found")
	}

	return smstructs.SecretEntry{
		Key:   tenantId.String(),
		Value: dataMap,
	}, nil
}

func (s *StackitRepository) GetAllTenantPws(ctx context.Context) ([]smstructs.SecretEntry, error) {

	allEntrys, err := s.client.Logical().ListWithContext(ctx, s.defaultTenantPath)

	if err != nil {
		return []smstructs.SecretEntry{}, err
	}

	if allEntrys == nil {
		return []smstructs.SecretEntry{}, errors.New("List is empty")
	}

	keys, ok := allEntrys.Data["keys"].([]any)

	if !ok {
		return []smstructs.SecretEntry{}, errors.New("Failed to get keys")
	}

	allPws := make([]smstructs.SecretEntry, 0, len(keys))

	for i := range keys {

		key, ok := keys[i].(string)

		if !ok {
			continue
		}

		if strings.HasSuffix(key, "/") {
			continue
		}

		secret, err := s.client.KVv2(s.instanceId).Get(ctx, s.defaultTenantPath+key)

		if err != nil {
			// We log the error and then just skipp the entry
			logging.Log(logging.Error, "[SecretMananger | GetAllTenantPws]"+err.Error())
			continue
		}

		pw, ok := secret.Data[repository.TENANT_PW_KEY].(string)

		if !ok {
			continue
		}

		allPws = append(allPws, smstructs.SecretEntry{Key: key, Value: pw})
	}

	return allPws, nil
}

func (s *StackitRepository) GetClientSecret(ctx context.Context, providerKey string) (smstructs.SecretEntry, error) {

	secret, err := s.client.KVv2(s.instanceId).Get(ctx, s.defaultClientSecretPath)

	if err != nil {
		return smstructs.SecretEntry{}, err
	}

	dataMap, ok := secret.Data[repository.CLIENT_SECRET_KEY].(string)

	if !ok {
		return smstructs.SecretEntry{}, errors.New("Key not found")
	}

	return smstructs.SecretEntry{
		Key:   providerKey,
		Value: dataMap,
	}, nil
}

func (s *StackitRepository) GetClientId(ctx context.Context, providerKey string) (smstructs.SecretEntry, error) {

	secret, err := s.client.KVv2(s.instanceId).Get(ctx, s.defaultClientIdPath)

	if err != nil {
		return smstructs.SecretEntry{}, err
	}

	dataMap, ok := secret.Data[repository.CLIENT_ID_KEY].(string)

	if !ok {
		return smstructs.SecretEntry{}, errors.New("Key not found")
	}

	return smstructs.SecretEntry{
		Key:   providerKey,
		Value: dataMap,
	}, nil
}

func (s *StackitRepository) CreateTenantPw(ctx context.Context, tenantId uuid.UUID, pw string) error {

	if _, err := s.client.KVv2(s.instanceId).Put(ctx, s.defaultTenantPath+tenantId.String(), map[string]any{repository.TENANT_PW_KEY: pw}); err != nil {
		return err
	}

	return nil
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

	if err := s.client.KVv2(s.instanceId).Delete(ctx, s.defaultTenantPath+tenantId.String()); err != nil {
		return err
	}

	return nil
}
