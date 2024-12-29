package helper

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

// Generate Jwt
func GenerateToken(ttl time.Duration, payload interface{}, privateKey string) (string, error) {
	decodedKey, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", fmt.Errorf("could not decode key: %w", err)
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedKey)
	if err != nil {
		return "", fmt.Errorf("create: parse key: %w", err)
	}
	now := time.Now().UTC()
	tokenClaims := make(jwt.MapClaims)
	tokenClaims["sub"] = payload
	tokenClaims["exp"] = now.Add(ttl).Unix()
	tokenClaims["iat"] = now.Unix()
	tokenClaims["nbf"] = now.Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, tokenClaims).SignedString(key)
	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}
	return token, nil
}