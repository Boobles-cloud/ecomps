package services

import (
	"context"
	"errors"
	"net/http"
	"os"
	"time"

	authstructs "ecomps.boobles.cloud/backend/internal/auth/auth_structs"
	"github.com/golang-jwt/jwt/v4"
)

func (a *AuthService) CreateAccessTokenInDatabase(ctx context.Context, item authstructs.JWTDatabaseStruct) error {
	return a.CreateAccessTokenInDatabase(ctx, item)
}

func (a *AuthService) DeleteAccessTokenByValue(ctx context.Context, cookieVal string) error {
	return a.DeleteAccessTokenByValue(ctx, cookieVal)
}

func (a *AuthService) CreateAuthCookie(ctx context.Context, userId, tenantId uint) (http.Cookie, error) {

	if userId == 0 {
		return http.Cookie{}, errors.New("UserId cant be 0")
	}

	if tenantId == 0 {
		return http.Cookie{}, errors.New("TenantId cant be 0")
	}

	claims := authstructs.JWTClaimsStruct{
		UserId:   userId,
		TenantId: tenantId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Ecomps_backend_server",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, 3)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString([]byte(os.Getenv("jwt_secret")))

	if err != nil {
		return http.Cookie{}, err
	}

	tokenDb := authstructs.JWTDatabaseStruct{
		UserAccessId: 0,
		TokenVal:     tokenSigned,
		TokenExpire:  time.Now().AddDate(0, 0, 3),
		UserId:       userId,
	}

	if err := a.CreateAccessTokenInDatabase(ctx, tokenDb); err != nil {
		return http.Cookie{}, err
	}

	return http.Cookie{
		Name:     AuthTokenCookieName,
		Value:    tokenSigned,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().AddDate(0, 0, 3),
	}, nil
}
