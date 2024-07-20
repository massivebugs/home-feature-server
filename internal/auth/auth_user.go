package auth

import "time"

type AuthUser struct {
	ID        uint32
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
