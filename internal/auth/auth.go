package auth

import (
	"context"

	"github.com/massivebugs/home-feature-server/api/dto"
)

type Auth struct{}

func NewAuth() *Auth {
	return &Auth{}
}

func (a *Auth) CreateUser(ctx context.Context, req *dto.CreateUserRequestDTO) error {
	return nil
}
