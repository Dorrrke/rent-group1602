package auth

import (
	"errors"
	"time"

	"github.com/Dorrrke/rent-group1602/internal/domain/users"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("super-secret")

type TokenClaims struct {
	UserID string     `json:"user_id"`
	Role   users.Role `json:"role"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(uid string, role users.Role) (string, error) {
	claims := TokenClaims{
		UserID: uid,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

func GenerateRefreshToken(uid string, role users.Role) (string, error) {
	claims := TokenClaims{
		UserID: uid,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

func ParseToken(tokenString string) (string, users.Role, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&TokenClaims{},
		func(t *jwt.Token) (any, error) {
			return jwtSecret, nil
		},
	)
	if err != nil {
		return "", -1, err
	}

	claims, ok := token.Claims.(*TokenClaims)

	if !ok || !token.Valid {
		return "", -1, errors.New("invalid jwt")
	}

	return claims.UserID, claims.Role, nil
}
