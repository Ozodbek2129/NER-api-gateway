package auth

import (
	"errors"
	"gateway/config"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func extractToken(tokenStr string) string {
	return strings.TrimPrefix(tokenStr, "Bearer ")
}

func parseToken(tokenStr string) (jwt.MapClaims, error) {
	cfg := config.Load()

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
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

func ValidateAccessToken(tokenStr string) (jwt.MapClaims, error) {
	return parseToken(extractToken(tokenStr))
}

func GetUserInfoFromAccessToken(tokenStr string) (string, string, string, error) {
	claims, err := parseToken(extractToken(tokenStr))
	if err != nil {
		return "", "", "", err
	}

	userID, ok1 := claims["user_id"].(string)
	email, ok2 := claims["email"].(string)
	role, ok3 := claims["role"].(string)

	if !ok1 || !ok2 || !ok3 {
		return "", "", "", errors.New("invalid claim types")
	}

	return userID, email, role, nil
}
