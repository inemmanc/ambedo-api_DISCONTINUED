package auth

import (
	"ambedo-api/src/config"
	"time"

	jwt "github.com/golang-jwt/jwt"
)

// CreateToken returns a signed token with user permissions
func CreateToken(userID uint64) (string, error) {
	perms := jwt.MapClaims{}
	perms["authorized"] = true
	perms["exp"] = time.Now().Add(time.Hour * 72).Unix()
	perms["userID"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, perms)

	return token.SignedString([]byte(config.DefaultSecretKey))
}
