package auth

import (
	jwt "github.com/golang-jwt/jwt"
)

func CreateTokne(userID uint64) (string, error) {
	perms := jwt.MapClaims{}
	perms["authorized"] = true

	// TEMP RETURN ---- REMOVE
	return "", nil
}
