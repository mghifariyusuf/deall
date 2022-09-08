package repositories

import (
	"context"
	"deall/cmd/entity"
)

type UserRepository interface {
	ListUser(ctx context.Context, page, limit int64) ([]*entity.User, error)
	GetUserByID(ctx context.Context, id string) (*entity.User, error)
	GetUserRoleID(ctx context.Context, roleID string) ([]*entity.User, error)
	GetUserByEmailOrPhone(ctx context.Context, data string) (*entity.User, error)
	InsertUser(ctx context.Context, user *entity.User) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User, id string) (*entity.User, error)
	DeleteUser(ctx context.Context, id string, userID string) error
}
