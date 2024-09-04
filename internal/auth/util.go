package auth

import (
	"strconv"
	"time"
)

func GenerateRandomString(length int) string {
	// TODO - implement
	return strconv.Itoa(int(time.Now().Unix()))
}
