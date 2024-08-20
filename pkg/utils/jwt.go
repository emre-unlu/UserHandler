package utils

import (
	"github.com/emre-unlu/GinTest/internal/config"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(email string) (accessToken string, refreshToken string, err error) {

	expirationTimeAccess := time.Now().Add(config.AccessTokenExpiration)
	claimsAccess := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTimeAccess),
			Issuer:    "emre",
		},
	}
	tokenA := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsAccess)
	accessToken, err = tokenA.SignedString(config.JWTSecretKey)
	if err != nil {
		return "", "", err
	}

	expirationTimeRefresh := time.Now().Add(config.RefreshTokenExpiration)
	claimsRefresh := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTimeRefresh),
			Issuer:    "emre",
		},
	}
	tokenR := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	refreshToken, err = tokenR.SignedString(config.JWTSecretKey)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil

}

func ValidateJWT(accessToken string) (claims *Claims, err error) {
	claims = &Claims{}
	token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return config.JWTSecretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}
