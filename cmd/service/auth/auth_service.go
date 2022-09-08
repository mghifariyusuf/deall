package service

import (
	"context"
	"database/sql"
	"deall/cmd/entity"
	"deall/cmd/lib/authentication"
	"deall/cmd/lib/customError"
	"deall/cmd/repositories"
	"deall/pkg/auth"
	"log"

	"github.com/sirupsen/logrus"
)

type authService struct {
	userRepository repositories.UserRepository
	authorization  auth.Auth
	token          auth.Token
}

func NewAuthService(userRepository repositories.UserRepository, authorization auth.Auth, token auth.Token) *authService {
	return &authService{
		userRepository: userRepository,
		authorization:  authorization,
		token:          token,
	}
}

func (s *authService) Login(ctx context.Context, data string, password string) (*entity.Authorization, error) {
	user, err := s.userRepository.GetUserByEmailOrPhone(ctx, data)
	if err == sql.ErrNoRows {
		logrus.Error(err)
		return nil, customError.ErrInvalidLogin
	}
	if err != nil {
		log.Println(err)
		return nil, customError.ErrInternalServerError
	}

	ok := authentication.ComparePassword(user.Password, password)
	if !ok {
		logrus.Error("password doesn't match")
		return nil, customError.ErrInvalidLogin
	}

	tokenDetail, err := s.token.CreateToken(user.ID, user.RoleID)
	if err != nil {
		log.Println(err)
		return nil, customError.ErrUnProcessableEntity
	}

	err = s.authorization.CreateAuth(user.ID, tokenDetail)
	if err != nil {
		log.Println(err)
		return nil, customError.ErrUnProcessableEntity
	}

	result := &entity.Authorization{
		User:    *user,
		Token:   tokenDetail.AccessToken,
		Refresh: tokenDetail.RefreshToken,
	}
	return result, err
}

func (s *authService) Logout(ctx context.Context, accessDetail *auth.AccessDetails) error {
	err := s.authorization.DeleteToken(accessDetail)
	if err != nil {
		logrus.Error(err)
		return customError.ErrInternalServerError
	}
	return nil
}
