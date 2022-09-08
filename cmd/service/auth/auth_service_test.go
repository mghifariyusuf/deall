package service

import (
	"context"
	"database/sql"
	"deall/cmd/entity"
	"deall/cmd/lib/authentication"
	"deall/cmd/mocks"
	auth "deall/pkg/auth"
	authMock "deall/pkg/auth/mock"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var mockUserRepo = new(mocks.UserRepository)
var mockToken = new(authMock.MockToken)
var mockAuth = new(authMock.MockAuth)
var service = NewAuthService(mockUserRepo, mockAuth, mockToken)

func TestLogin(t *testing.T) {
	var data = "coba@mail.com"
	password, err := authentication.SetPassword("coba")
	require.NoError(t, err)
	var userPayload = &entity.User{
		ID:          "ini id",
		Username:    "user1",
		Email:       "coba@mail.com",
		FirstName:   "first_name",
		LastName:    "last_name",
		Password:    "coba",
		PhoneNumber: "0813456789",
		CreatedAt:   time.Now(),
		CreatedBy:   "o",
		UpdatedAt:   nil,
		UpdatedBy:   nil,
	}

	var tokenPayload = &auth.TokenDetails{
		AccessToken:         "okokokokok",
		RefreshToken:        "okokokokokok",
		RefreshTokenExpires: time.Now().Add(time.Hour * 24 * 7).Unix(),
		RefreshUUID:         "okokokokok",
		RoleID:              userPayload.RoleID,
		TokenExpires:        time.Now().Add(time.Minute * 24 * 7).Unix(),
		TokenUUID:           "okokokook",
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("GetUserByEmailOrPhone", mock.Anything, mock.AnythingOfType("string")).Return(userPayload, nil).Once()
		mockToken.On("CreateToken", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(tokenPayload, nil).Once()
		mockAuth.On("CreateAuth", mock.AnythingOfType("string"), mock.Anything).Return(nil).Once()

		loginResponse, err := service.Login(context.TODO(), data, password)
		require.NoError(t, err)
		require.NotNil(t, loginResponse)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("error-password-doesnt-match", func(t *testing.T) {
		mockUserRepo.On("GetUserByEmailOrPhone", mock.Anything, mock.AnythingOfType("string")).Return(userPayload, nil).Once()

		loginResponse, err := service.Login(context.TODO(), data, "salah")
		require.Error(t, err)
		require.Nil(t, loginResponse)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("error-data-not-found", func(t *testing.T) {
		mockUserRepo.On("GetUserByEmailOrPhone", mock.Anything, mock.AnythingOfType("string")).Return(nil, sql.ErrNoRows).Once()

		loginResponse, err := service.Login(context.TODO(), data, password)
		require.Error(t, err)
		require.Nil(t, loginResponse)
	})

	t.Run("error-repository", func(t *testing.T) {
		mockUserRepo.On("GetUserByEmailOrPhone", mock.Anything, mock.AnythingOfType("string")).Return(nil, fmt.Errorf("error")).Once()

		loginResponse, err := service.Login(context.TODO(), data, password)
		require.Error(t, err)
		require.Nil(t, loginResponse)
	})

	t.Run("error-token", func(t *testing.T) {
		mockUserRepo.On("GetUserByEmailOrPhone", mock.Anything, mock.AnythingOfType("string")).Return(userPayload, nil).Once()
		mockToken.On("CreateToken", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil, fmt.Errorf("token")).Once()

		loginResponse, err := service.Login(context.TODO(), data, password)
		require.Error(t, err)
		require.Nil(t, loginResponse)
	})

	t.Run("error-auth", func(t *testing.T) {
		mockUserRepo.On("GetUserByEmailOrPhone", mock.Anything, mock.AnythingOfType("string")).Return(userPayload, nil).Once()

		mockToken.On("CreateToken", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(tokenPayload, nil).Once()
		mockAuth.On("CreateAuth", mock.AnythingOfType("string"), mock.Anything).Return(fmt.Errorf("error")).Once()

		loginResponse, err := service.Login(context.TODO(), data, password)
		require.Error(t, err)
		require.Nil(t, loginResponse)
	})
}
