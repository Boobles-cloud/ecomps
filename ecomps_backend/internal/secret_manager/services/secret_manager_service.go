package services

import (
	"context"
	"errors"
	"math/rand"
	"time"
	"uuid"

	smstructs "ecomps.boobles.cloud/backend/internal/secret_manager/sm_structs"
)

const (
	allCharacters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	keyLength     = 32
)

// Gets all tenant Pws and stores them inside the cache
func (s *SecretManagerService) SetupCache(ctx context.Context) error {

	entrys, err := s.secretManagerRepo.GetAllTenantPws(ctx)

	if err != nil {
		return err
	}

	// Setup our cache with all entrys
	for i := range entrys {

		if entrys[i].Key == uuid.Nil().String() || entrys[i].Key == uuid.Max().String() {
			continue
		}

		s.cacheManager.SetOrUpdateItem(entrys[i].Key, entrys[i], uuid.Nil())
	}

	return nil
}

func (s *SecretManagerService) GetTenantPw(ctx context.Context, tenantId uuid.UUID) (smstructs.SecretEntry, error) {

	if tenantId.Compare(uuid.Nil()) == 0 || tenantId.Compare(uuid.Max()) == 0 {
		return smstructs.SecretEntry{}, errors.New("TenantId is invalid")
	}

	cacheEntry, ok := s.cacheManager.GetItem(tenantId.String())

	if ok {
		return cacheEntry, nil
	}

	entry, err := s.secretManagerRepo.GetTenantPw(ctx, tenantId)

	if err != nil {
		return entry, err
	}

	s.cacheManager.SetOrUpdateItem(entry.Key, entry, uuid.Nil())

	return entry, nil
}

func (s *SecretManagerService) GetClientSecret(ctx context.Context, providerKey string) (smstructs.SecretEntry, error) {

	if providerKey == "" {
		return smstructs.SecretEntry{}, errors.New("ProviderId is emtpty")
	}

	return s.secretManagerRepo.GetClientSecret(ctx, providerKey)
}

// Creates a tenant pw
func (s *SecretManagerService) CreateTenantPw(ctx context.Context, tenantId uuid.UUID) error {

	if tenantId.Compare(uuid.Nil()) == 0 || tenantId.Compare(uuid.Max()) == 0 {
		return errors.New("TenantId is invalid")
	}

	randSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randSource)

	result := make([]byte, keyLength)

	for i := range keyLength {
		result[i] = allCharacters[random.Intn(len(allCharacters))]
	}

	return s.secretManagerRepo.CreateTenantPw(ctx, tenantId, string(result))
}

// Just to check if the provider is reachable
func (s *SecretManagerService) CheckProviderHealth(ctx context.Context) error {
	return s.CheckProviderHealth(ctx)
}

func (s *SecretManagerService) DelteTenantPw(ctx context.Context, tenantId uuid.UUID) error {

	if tenantId.Compare(uuid.Nil()) == 0 || tenantId.Compare(uuid.Max()) == 0 {
		return errors.New("TenantId is invalid")
	}

	s.cacheManager.RemoveItem(tenantId.String())

	return s.secretManagerRepo.DeleteTenantPw(ctx, tenantId)
}
