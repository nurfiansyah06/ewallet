package middleware

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(ttl time.Duration, payload interface{}, secretJWT string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	now := time.Now().UTC()
	claim := token.Claims.(jwt.MapClaims)

	claim["sub"] = payload
	claim["exp"] = now.Add(ttl).Unix()
	claim["iat"] = now.Unix()
	claim["nbf"] = now.Unix()

	tokenString, err := token.SignedString([]byte(secretJWT))
	if err != nil {
		return "", fmt.Errorf("failed token: %w", err)
	}

	return tokenString, nil
}

func ValidateToken(token string, signedJWT string) (interface{}, error) {
	tokenString, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}

		return []byte(signedJWT), nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token %w", err)
	}

	claims, ok := tokenString.Claims.(jwt.MapClaims)
	if !ok || !tokenString.Valid {
		return nil, fmt.Errorf("invalid token claim")
	}

	return claims["sub"], nil
}