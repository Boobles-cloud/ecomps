package startup

import (
	"context"
	"os"

	"ecomps.boobles.cloud/backend/internal/secret_manager/repositorys"
	"ecomps.boobles.cloud/backend/internal/secret_manager/services"
	vault "github.com/hashicorp/vault/api"
	auth "github.com/hashicorp/vault/api/auth/approle"
)

func ConfigureSecretManager(ctx context.Context) (*services.SecretManagerService, error) {

	// TODO: Add env vars

	conf := vault.DefaultConfig()
	conf.Address = os.Getenv("")

	secretId := auth.SecretID{FromString: os.Getenv("")}
	appRoleAuth, err := auth.NewAppRoleAuth(os.Getenv(""), &secretId)

	if err != nil {
		return &services.SecretManagerService{}, err
	}

	client, err := vault.NewClient(conf)

	if err != nil {
		return &services.SecretManagerService{}, err
	}

	// Authenticate our client
	if _, err := client.Auth().Login(ctx, appRoleAuth); err != nil {
		return &services.SecretManagerService{}, err
	}

	// Our current provider
	repository := repositorys.CreateNewStackitRepository(client)

	return services.CreateNewSecretManagerService(repository), nil
}
