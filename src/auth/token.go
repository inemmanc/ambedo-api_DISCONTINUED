package auth

import (
	"time"

	jwt "github.com/golang-jwt/jwt"
)

func CreateToken(userID uint64) (string, error) {
	perms := jwt.MapClaims{}
	perms["authorized"] = true
	perms["exp"] = time.Now().Add(time.Hour * 72).Unix()
	perms["userID"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, perms)

	// TEMP SECRET KEY ---- REMOVE
	return token.SignedString([]byte("Secret"))
}
