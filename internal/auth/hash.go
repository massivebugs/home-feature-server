package auth

import "golang.org/x/crypto/bcrypt"

func GeneratePasswordHash(password string) (string, error) {
	// Cost is the number of times the key is stretched.
	// GenerateFromPassword returns a zero value []byte,
	// but we're returning a "" just for documentation purposes
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashed), err
}

func CheckPasswordHash(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
