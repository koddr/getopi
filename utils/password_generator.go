package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// NormalizePassword ...
func NormalizePassword(pwd string) []byte {
	// Return the users input as a byte slice which will save us
	// from having to do this conversion later on
	return []byte(pwd)
}

// GeneratePassword ...
func GeneratePassword(pwd string) string {
	// Normalize password from string to []byte
	bytePwd := NormalizePassword(pwd)

	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(bytePwd, bcrypt.MinCost)
	if err != nil {
		return err.Error()
	}

	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

// ComparePasswords ...
func ComparePasswords(hashedPwd string, inputPwd string) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := NormalizePassword(hashedPwd)
	byteInput := NormalizePassword(inputPwd)

	// Return result
	if err := bcrypt.CompareHashAndPassword(byteHash, byteInput); err != nil {
		return false
	}

	return true
}
