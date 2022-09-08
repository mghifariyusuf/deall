package user

import (
	"context"
	"database/sql"
	"deall/cmd/entity"
	"deall/cmd/lib/authentication"
	"deall/cmd/lib/customError"
	"deall/cmd/repositories"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *userService {
	return &userService{userRepository: userRepository}
}

func (s *userService) ListUser(ctx context.Context, page, limit int64) ([]*entity.User, error) {

	users, err := s.userRepository.ListUser(ctx, page, limit)
	if err == sql.ErrNoRows || users == nil {
		return make([]*entity.User, 0), nil
	} else if err != nil {
		logrus.Error(err)
		return nil, customError.ErrInternalServerError
	}

	return users, nil
}

func (s *userService) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	user, err := s.userRepository.GetUserByID(ctx, id)
	if err == sql.ErrNoRows || user == nil {
		logrus.Error(err)
		return nil, customError.ErrDataNotFound
	} else if err != nil {
		logrus.Error(err)
		return nil, customError.ErrInternalServerError
	}

	return user, nil
}

func (s *userService) InsertUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	id := uuid.New()
	pass, err := authentication.SetPassword(user.Password)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	user.Password = pass
	user.ID = id.String()

	user, err = s.userRepository.InsertUser(ctx, user)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return user, nil
}

func (s *userService) UpdateUser(ctx context.Context, user *entity.User, id string) (*entity.User, error) {
	u, err := s.userRepository.GetUserByID(ctx, id)
	if err == sql.ErrNoRows || u == nil {
		logrus.Error(err)
		return nil, customError.ErrDataNotFound
	} else if err != nil {
		logrus.Error(err)
		return nil, err
	}
	user.Password = u.Password

	user, err = s.userRepository.UpdateUser(ctx, user, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return user, nil
}

func (s *userService) DeleteUser(ctx context.Context, id string, userID string) error {
	u, err := s.userRepository.GetUserByID(ctx, id)
	if err == sql.ErrNoRows || u == nil {
		logrus.Error(err)
		return customError.ErrDataNotFound
	} else if err != nil {
		logrus.Error(err)
		return err
	}
	err = s.userRepository.DeleteUser(ctx, id, userID)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
