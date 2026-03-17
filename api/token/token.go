package auth

import (
	"errors"
	"gateway/config"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var cfg = config.Load()

func extractToken(tokenStr string) string {
	return strings.TrimPrefix(tokenStr, "Bearer ")
}

func ValidateAccessToken(tokenStr string) (jwt.MapClaims, error) {
	tokenStr = extractToken(tokenStr)

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(cfg.SIGNING_KEY), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	return claims, nil
}

func GetUserInfoFromAccessToken(tokenStr string) (string, string, string, error) {
	tokenStr = extractToken(tokenStr)

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(cfg.SIGNING_KEY), nil
	})
	if err != nil || !token.Valid {
		return "", "", "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", "", errors.New("invalid claims")
	}

	userID, ok1 := claims["user_id"].(string)
	email, ok2 := claims["email"].(string)
	role, ok3 := claims["role"].(string)

	if !ok1 || !ok2 || !ok3 {
		return "", "", "", errors.New("invalid claim types")
	}

	return userID, email, role, nil
}