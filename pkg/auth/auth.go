package auth

import (
	"deall/cmd/lib/customError"
	"deall/pkg/cache"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type TokenDetails struct {
	AccessToken         string
	RefreshToken        string
	TokenUUID           string
	RefreshUUID         string
	RoleID              string
	TokenExpires        int64
	RefreshTokenExpires int64
}

type AccessDetails struct {
	TokenUUID string
	UserID    string
	RoleID    string
}

type Auth interface {
	CreateAuth(userID string, tokenDetails *TokenDetails) error
	FetchAuth(tokenUUID string) (string, error)
	DeleteToken(accessDetail *AccessDetails) error
	DeleteRefresh(refreshUUID string) error
}

type service struct {
	conn cache.RedisCommand
}

func New(conn cache.RedisCommand) *service {
	return &service{conn}
}

func (s *service) CreateAuth(userID string, tokenDetails *TokenDetails) error {
	token := time.Unix(tokenDetails.TokenExpires, 0)
	refreshToken := time.Unix(tokenDetails.RefreshTokenExpires, 0)
	now := time.Now()

	err := s.conn.SetEx(tokenDetails.TokenUUID, token.Unix()-now.Unix(), userID)
	if err != nil {
		logrus.Error(err)
		return err
	}

	err = s.conn.SetEx(tokenDetails.RefreshToken, refreshToken.Unix()-now.Unix(), userID)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (s *service) FetchAuth(tokenUUID string) (string, error) {
	userID, err := s.conn.Get(tokenUUID)
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	if string(userID) == "" {
		logrus.Error(customError.ErrDataNotFound)
		return "", customError.ErrToken
	}
	return string(userID), nil
}

func (s *service) DeleteToken(accessDetail *AccessDetails) error {
	refreshUUID := fmt.Sprintf("%s++%s", accessDetail.TokenUUID, accessDetail.UserID)
	err := s.conn.Del(accessDetail.TokenUUID)
	if err != nil {
		logrus.Error(err)
		return err
	}

	err = s.conn.Del(refreshUUID)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (s *service) DeleteRefresh(refreshUUID string) error {
	err := s.conn.Del(refreshUUID)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
