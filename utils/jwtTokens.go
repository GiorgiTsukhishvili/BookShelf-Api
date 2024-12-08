package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserID uint
	Email  string
	jwt.RegisteredClaims
}

type JWTInfo struct {
	Token                  string
	RefreshToken           string
	TokenExpiration        time.Time
	RefreshTokenExpiration time.Time
}

func GenerateJWTTokens(userID uint, email string) (*JWTInfo, error) {

	expiration, err := strconv.Atoi(os.Getenv("JWT_TOKEN_EXPIRATION_DATE"))
	if err != nil {
		return nil, err
	}

	refreshExpiration, err := strconv.Atoi(os.Getenv("JWT_REFRESH_TOKEN_EXPIRATION_DATE"))
	if err != nil {
		return nil, err
	}

	tokenExpiration := time.Now().Add(time.Duration(expiration) * time.Minute)
	refreshTokenExpiration := time.Now().Add(time.Duration(refreshExpiration) * time.Minute)

	claims := CustomClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(tokenExpiration),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString([]byte(os.Getenv("ACCESS_TOKEN_SECRET")))
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtRefreshToken, err := refreshToken.SignedString([]byte(os.Getenv("REFRESH_TOKEN_SECRET")))
	if err != nil {
		return nil, err
	}

	jwtInfo := JWTInfo{
		Token:                  jwtToken,
		RefreshToken:           jwtRefreshToken,
		TokenExpiration:        tokenExpiration,
		RefreshTokenExpiration: refreshTokenExpiration,
	}

	return &jwtInfo, nil
}
