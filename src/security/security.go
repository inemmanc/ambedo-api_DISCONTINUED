package security

import "golang.org/x/crypto/bcrypt"

// takes a password string and hashes it
func Hash(password string,) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// compares the default password to the hashed password and returns if its matchs
func VerifyPassword(hashedPassword, stringPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(stringPassword))
}