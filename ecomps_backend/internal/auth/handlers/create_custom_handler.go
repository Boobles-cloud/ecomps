package handlers

import (
	"ecomps.boobles.cloud/backend/internal/auth/services"
	userService "ecomps.boobles.cloud/backend/internal/user/services"
)

type AuthHandler struct {
	authService *services.AuthService
	userService *userService.UserService
}

// Creates a new auth handler
func CreateAuthHandler(a *services.AuthService, u *userService.UserService) *AuthHandler {
	return &AuthHandler{
		authService: a,
		userService: u,
	}
}
