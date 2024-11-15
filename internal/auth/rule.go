package auth

import (
	"reflect"
	"unicode"

	"github.com/go-playground/validator/v10"
)

const (
	MIN_USERNAME_LENGTH int = 3
	MAX_USERNAME_LENGTH int = 100
	MIN_PASSWORD_LENGTH int = 8
	MAX_PASSWORD_LENGTH int = 72 // bcrypt length limit
)

/*
Validates password based on the following conditions:

  - length(byte) is between specified minLength and maxLength
  - includes a number
  - includes a special character
  - includes a letter
  - does not have unsupported characters
*/
// func isValidPassword(minLength int, maxLength int) func(v interface{}) error {
// 	return func(v interface{}) error {
// 		s, ok := v.(string)
// 		if !ok {
// 			return errors.New("failed to validate password")
// 		}

// 		length := 0
// 		hasLetter := false
// 		hasNumber := false
// 		hasSpecial := false
// 		for _, c := range s {
// 			switch {
// 			case unicode.IsNumber(c):
// 				hasNumber = true
// 			case unicode.IsPunct(c) || unicode.IsSymbol(c):
// 				hasSpecial = true
// 			case unicode.IsLetter(c) || c == ' ':
// 				hasLetter = true
// 			default:
// 				return errors.New("password contains unsupported character")
// 			}
// 			length++
// 		}

// 		switch {
// 		case length < minLength:
// 			return fmt.Errorf("password must be longer than %v characters", minLength)
// 		case length > maxLength:
// 			return errors.New("password is too long")
// 		case !hasLetter || !hasNumber || !hasSpecial:
// 			return errors.New("password must include at least one letter, one number and one special character")
// 		}

// 		return nil
// 	}
// }

/*
Validates password based on the following conditions:

  - length(byte) is between specified minLength and maxLength
  - includes a number
  - includes a special character
  - includes a letter
  - does not have unsupported characters
*/
func IsValidPassword(fl validator.FieldLevel) bool {
	if fl.Field().Kind() != reflect.String {
		return false
	}

	value := fl.Field().String()

	length := 0
	hasLetter := false
	hasNumber := false
	hasSpecial := false
	for _, c := range value {
		switch {
		case unicode.IsNumber(c):
			hasNumber = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecial = true
		case unicode.IsLetter(c) || c == ' ':
			hasLetter = true
		default:
			return false
		}
		length++
	}

	switch {
	case length < MIN_PASSWORD_LENGTH:
		return false
	case length > MAX_PASSWORD_LENGTH:
		return false
	case !hasLetter || !hasNumber || !hasSpecial:
		return false
	}

	return true
}
