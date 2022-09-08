package redis

import (
	"context"
	"deall/cmd/entity"
	"deall/cmd/mocks"
	cache "deall/pkg/cache/mock"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/mock"
)

var mockUserRepo = new(mocks.UserRepository)
var mockCache = new(cache.MockRedisCommand)
var redis = NewUser(mockCache, mockUserRepo, "project")

func TestListUser(t *testing.T) {
	var limit int64 = 10
	var page int64 = 1
	var userPayload = []*entity.User{
		&entity.User{
			ID:          "ini id",
			Username:    "user1",
			Email:       "coba@mail.com",
			FirstName:   "first_name",
			LastName:    "last_name",
			Password:    "ok",
			PhoneNumber: "0813456789",
			CreatedAt:   time.Now(),
			CreatedBy:   "o",
			UpdatedAt:   nil,
			UpdatedBy:   nil,
		},
		&entity.User{
			ID:          "ini idi",
			Username:    "user2",
			Email:       "id2@email.com",
			FirstName:   "first2_name",
			LastName:    "last2_name",
			Password:    "ok",
			PhoneNumber: "08134567890",
			CreatedAt:   time.Now(),
			CreatedBy:   "o",
			UpdatedAt:   nil,
			UpdatedBy:   nil,
		},
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("ListUser", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("int64")).Return(userPayload, nil).Once()

		userResponse, err := redis.ListUser(context.TODO(), page, limit)
		require.NoError(t, err)
		require.NotNil(t, userResponse)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestGetUserByEmailOrPhone(t *testing.T) {
	var data = "id2@email.com"
	var userPayload = &entity.User{
		ID:          "ini idi",
		Username:    "user2",
		Email:       "id2@email.com",
		FirstName:   "first2_name",
		LastName:    "last2_name",
		Password:    "ok",
		PhoneNumber: "08134567890",
		CreatedAt:   time.Now(),
		CreatedBy:   "o",
		UpdatedAt:   nil,
		UpdatedBy:   nil,
	}

	mockCacheUser, _ := json.Marshal(userPayload)
	mockCacheUserError, _ := json.Marshal("OK")

	t.Run("success-cache-found", func(t *testing.T) {
		mockCache.On("Get", mock.Anything, mock.AnythingOfType("string")).Return(mockCacheUser, nil).Once()

		userResponse, err := redis.GetUserByEmailOrPhone(context.TODO(), data)
		require.NoError(t, err)
		require.NotNil(t, userResponse)
		require.Equal(t, userPayload.ID, userResponse.ID)
		mockCache.AssertExpectations(t)

	})

	t.Run("error-cache-found", func(t *testing.T) {
		mockCache.On("Get", mock.Anything, mock.AnythingOfType("string")).Return(mockCacheUserError, nil).Once()

		userResponse, err := redis.GetUserByEmailOrPhone(context.TODO(), data)
		require.Error(t, err)
		require.Nil(t, userResponse)
		mockCache.AssertExpectations(t)
	})

	t.Run("success-cache-not-found", func(t *testing.T) {
		mockCache.On("Get", mock.Anything, mock.AnythingOfType("string")).Return(nil, nil).Once()
		mockUserRepo.On("GetUserByEmailOrPhone", mock.Anything, mock.AnythingOfType("string")).Return(userPayload, nil).Once()
		mockCache.On("SetEx", mock.Anything, mock.Anything, mock.AnythingOfType("string")).Return(nil).Once()

		userResponse, err := redis.GetUserByEmailOrPhone(context.TODO(), data)
		require.NoError(t, err)
		require.NotNil(t, userResponse)
		require.Equal(t, userPayload.ID, userResponse.ID)
		mockCache.AssertExpectations(t)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("error-cache", func(t *testing.T) {
		mockCache.On("Get", mock.Anything, mock.AnythingOfType("string")).Return(nil, fmt.Errorf("error")).Once()

		userResponse, err := redis.GetUserByEmailOrPhone(context.TODO(), data)
		require.Error(t, err)
		require.Nil(t, userResponse)
		mockCache.AssertExpectations(t)
	})

	t.Run("error-repository", func(t *testing.T) {
		mockCache.On("Get", mock.Anything, mock.AnythingOfType("string")).Return(nil, nil).Once()
		mockUserRepo.On("GetUserByEmailOrPhone", mock.Anything, mock.AnythingOfType("string")).Return(nil, fmt.Errorf("error")).Once()

		userResponse, err := redis.GetUserByEmailOrPhone(context.TODO(), data)
		require.Error(t, err)
		require.Nil(t, userResponse)
		mockCache.AssertExpectations(t)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("error-cache-set-ex", func(t *testing.T) {
		mockCache.On("Get", mock.Anything, mock.AnythingOfType("string")).Return(nil, nil).Once()
		mockUserRepo.On("GetUserByEmailOrPhone", mock.Anything, mock.AnythingOfType("string")).Return(nil, nil).Once()
		mockCache.On("SetEx", mock.Anything, mock.Anything, mock.AnythingOfType("string")).Return(fmt.Errorf("error")).Once()

		userResponse, err := redis.GetUserByEmailOrPhone(context.TODO(), data)
		require.Error(t, err)
		require.Nil(t, userResponse)
		mockCache.AssertExpectations(t)
		mockUserRepo.AssertExpectations(t)
	})

}
