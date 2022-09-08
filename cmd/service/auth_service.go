package service

import (
	"context"
	"deall/cmd/entity"
	"deall/pkg/auth"
)

type AuthService interface {
	Login(ctx context.Context, data string, password string) (*entity.Authorization, error)
	Logout(ctx context.Context, accessDetail *auth.AccessDetails) error
}
