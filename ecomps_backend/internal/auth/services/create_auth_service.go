package services

import (
	authstructs "ecomps.boobles.cloud/backend/internal/auth/auth_structs"
	"ecomps.boobles.cloud/backend/internal/repository"
	"ecomps.boobles.cloud/backend/internal/user/services"
)

const (
	AuthTokenCookieName = "EcompsAuthToken"
)

type AuthService struct {
	authRepository repository.AuthRepository
	userService    *services.UserService // We need the user service, because we want to get the user by id
}

func CreateNewAuthService(r repository.AuthRepository, u *services.UserService) *AuthService {
	return &AuthService{
		authRepository: r,
		userService:    u,
	}
}

func JwtToArgs(a authstructs.JWTDatabaseStruct) []any {
	return []any{
		a.UserAccessId,
		a.TokenVal,
		a.TokenExpire,
		a.UserId,
	}
}
