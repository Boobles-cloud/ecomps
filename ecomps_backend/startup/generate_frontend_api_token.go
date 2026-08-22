package startup

import (
	"crypto/rand"
	"math/big"
	"os"

	"ecomps.boobles.cloud/backend/utils/logging"
)

// Used on first init to generate a api token
// This func sets the api token as enviroment var
func GenerateFrontendApiToken() bool {

	allCharacters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	apiKey := make([]rune, 64)

	max := big.NewInt(int64(len(allCharacters)))

	for i := range apiKey {
		tmp, err := rand.Int(rand.Reader, max)

		if err != nil {
			logging.Log(logging.Error, "[Startup | GenerateFrontendApiToken] "+err.Error())
		}

		apiKey[i] = allCharacters[tmp.Int64()]
	}

	// Set it to this enviroment vars, so we dont need to read a file everytime we want to access it
	os.Setenv("API_Key", string(apiKey))

	// Store the api key to our .env file so the frontend can access it
	file, err := os.OpenFile(os.Getenv("env_path"), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		logging.Log(logging.Error, "[Startup | GenerateFrontendApiToken]"+err.Error())
		return false
	}

	defer file.Close()

	apiKeyEntry := "\nAPI_KEY=" + string(apiKey)

	if _, err := file.WriteString(apiKeyEntry); err != nil {
		logging.Log(logging.Error, "[Startup | GenerateFrontendApiToken]"+err.Error())
		return false
	}

	return true
}
