package auth

import "golang.org/x/crypto/bcrypt"

func PasswordEncrypt(password string) (string, error) {
	// Cost is the number of times the key is stretched.
	// GenerateFromPassword returns a zero value []byte,
	// but we're returning a "" just for documentation purposes
	hp, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hp), err
}

func CheckHashPassword(hp, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hp), []byte(password))
}
