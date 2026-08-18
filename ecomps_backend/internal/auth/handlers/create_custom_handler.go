package handlers

import "ecomps.boobles.cloud/backend/database"

type AuthHandler struct {
	Dh *database.DbHandler
}

// Creates a new auth handler
func CreateAuthHandler(dh *database.DbHandler) *AuthHandler {
	return &AuthHandler{
		Dh: dh,
	}
}
