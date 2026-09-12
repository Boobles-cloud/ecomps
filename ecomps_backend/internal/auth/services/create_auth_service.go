package services

import (
	authstructs "ecomps.boobles.cloud/backend/internal/auth/auth_structs"
	"ecomps.boobles.cloud/backend/internal/repository"
)

const (
	AuthTokenCookieName = "EcompsAuthToken"
)

type AuthService struct {
	authRepository repository.AuthRepository
}

func CreateNewAuthService(r repository.AuthRepository) *AuthService {
	return &AuthService{
		authRepository: r,
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
