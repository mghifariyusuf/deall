package redis

import (
	"context"
	"deall/cmd/entity"
	"deall/cmd/repositories"
	"deall/pkg/cache"
	"encoding/json"
	"fmt"
	"log"

	"github.com/sirupsen/logrus"
)

type Redis struct {
	conn   cache.RedisCommand
	repo   repositories.UserRepository
	prefix string
	module string
}

func NewUser(conn cache.RedisCommand, repo repositories.UserRepository, prefix string) *Redis {
	return &Redis{conn, repo, prefix, "user"}
}

func (r *Redis) ListUser(ctx context.Context, page, limit int64) ([]*entity.User, error) {
	return r.repo.ListUser(ctx, page, limit)
}

func (r *Redis) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	return r.repo.GetUserByID(ctx, id)
}

func (r *Redis) GetUserRoleID(ctx context.Context, roleID string) ([]*entity.User, error) {
	return r.repo.GetUserRoleID(ctx, roleID)
}

func (r *Redis) GetUserByEmailOrPhone(ctx context.Context, data string) (*entity.User, error) {
	var key = fmt.Sprintf("%s:%s:%s", r.prefix, r.module, data)
	var user = new(entity.User)

	value, err := r.conn.Get(key)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	if value == nil {
		user, err = r.repo.GetUserByEmailOrPhone(ctx, data)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		data, err := json.Marshal(user)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		err = r.conn.SetEx(key, cache.ONEHOUR, string(data))
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return user, err
	}

	err = json.Unmarshal(value, user)
	if err != nil {
		logrus.Println(err)
		return nil, err
	}
	return user, err
}

func (r *Redis) InsertUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	return r.repo.InsertUser(ctx, user)
}

func (r *Redis) UpdateUser(ctx context.Context, user *entity.User, id string) (*entity.User, error) {
	return r.repo.UpdateUser(ctx, user, id)
}

func (r *Redis) DeleteUser(ctx context.Context, id string, userID string) error {
	return r.repo.DeleteUser(ctx, id, userID)
}
