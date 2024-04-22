package auth

import (
	"ambedo-api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

// ValidateToken checks whether the received request token is valid
func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		// TEMP ---- REMOVE
		fmt.Printf(" >> invalid token request")
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// TEMP ---- REMOVE
		fmt.Printf(" >> %v", token.Claims)
		return nil
	}

	return errors.New("invalid token")
}

// ExtractUserID returns the userID saved in the token
func ExtractUserID(r *http.Request) (uint64, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		// TEMP ---- REMOVE
		fmt.Printf(" >> invalid token request")
		return 0, err
	}

	if perms, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, err := strconv.ParseUint(fmt.Sprintf("%.0f", perms["userID"]), 10, 64)
		if err != nil {
			return 0, err
		}

		return userID, nil
	}
	return 0, errors.New("invalid token")
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) != 2 {
		return ""
	}

	return strings.Split(token, " ")[1]
}

func returnVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpect signing method %v", token.Header["alg"])
	}

	return config.DefaultSecretKey, nil
}
