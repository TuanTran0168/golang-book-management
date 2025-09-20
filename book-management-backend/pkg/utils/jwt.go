package utils

import (
	configs "book-management/configs"
	"book-management/internal/models"
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

/*
This package handles JWT-related utilities:
- Claims: custom struct storing UserID, Role, and standard JWT claims.
- GenerateAccessToken: generates an access token with short TTL.
- GenerateRefreshToken: generates a refresh token with longer TTL.
- ParseAccessToken: validates and extracts claims from an access token.
*/

// Claims defines the custom JWT claims structure
type Claims struct {
	UserID uint        `json:"user_id"`
	Role   models.Role `json:"role"`
	jwt.RegisteredClaims
}

// GenerateAccessToken creates a signed JWT access token
// Params:
// - userID: unique identifier of the user
// - role: user's role (admin/user/...)
// - cfg: application config containing secret and TTL
// Returns: signed JWT string or error
func GenerateAccessToken(userID uint, role models.Role, cfg *configs.Config) (string, error) {
	expiresAt := time.Now().Add(time.Duration(cfg.AccessTokenTTLMinutes) * time.Minute)

	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWTSecret))
}

// GenerateRefreshToken creates a signed JWT refresh token
// Params:
// - userID: unique identifier of the user
// - cfg: application config containing secret and TTL
// Returns: signed JWT string, its expiration time, or error
func GenerateRefreshToken(userID uint, cfg *configs.Config) (string, time.Time, error) {
	expiresAt := time.Now().Add(time.Duration(cfg.RefreshTokenTTLHours) * time.Hour)

	claims := &jwt.RegisteredClaims{
		Subject:   strconv.FormatUint(uint64(userID), 10),
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(cfg.JWTSecret))
	return signed, expiresAt, err
}

// ParseAccessToken validates and decodes a JWT access token
// Params:
// - tokenStr: JWT access token string
// - cfg: application config containing secret
// Returns: pointer to Claims or error
func ParseAccessToken(tokenStr string, cfg *configs.Config) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
