package passwordh

import (
	"golang.org/x/crypto/bcrypt"
)

// CreatePassword create hash password from string
func CreatePassword(pwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
}

// ComparePasswords check password
func ComparePasswords(hashedPwd string, plainPwd string) bool {

	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)

	bytePlain := []byte(plainPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)

	if err != nil {
		return false
	}

	return true
}
