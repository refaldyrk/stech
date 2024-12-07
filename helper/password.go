package helper

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes the given password using bcrypt.
//
// password: the password to be hashed.
// string: the hashed password.
// error: an error if the hashing process fails.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

// CheckPasswordHash checks if a given password matches a given hash.
//
// Parameters:
// - password: the password string to check.
// - hash: the hash string to compare against the password.
//
// Returns:
// - bool: true if the password matches the hash, false otherwise.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
