package userhttp

import (
	"context"
	"deall/cmd/entity"
	"deall/cmd/lib/customError"
	"deall/pkg/helper"
	"deall/pkg/response"
	"deall/pkg/router"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

type contextKey string

func (dashboard *userDashboard) ListUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()

	var page = 1
	var limit = 10
	var param = r.URL.Query()
	var err error

	if param.Get("page") != "" {
		page, err = strconv.Atoi(param.Get("page"))
		if err != nil {
			logrus.Error(err)
			response.Error(w, customError.ErrInvalidBodyRequest)
			return
		}
	}

	if param.Get("limit") != "" {
		limit, err = strconv.Atoi(param.Get("limit"))
		if err != nil {
			logrus.Error(err)
			response.Error(w, customError.ErrInvalidBodyRequest)
			return
		}
	}

	result, err := dashboard.service.ListUser(ctx, int64(page), int64(limit))
	if err != nil {
		logrus.Error(err)
		response.Error(w, err)
		return
	}

	response.Success(w, result)
}

func (dashboard *userDashboard) GetUserByID(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()

	id := router.Param(r.Context(), "id")

	result, err := dashboard.service.GetUserByID(ctx, id)
	if err != nil {
		logrus.Error(err)
		response.Error(w, err)
		return
	}

	response.Success(w, result)
}

func (dashboard *userDashboard) InserUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()

	var params insertRequest

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&params)

	params.CreatedBy = r.Header.Get("userId")

	err := helper.Validate(params)
	if err != nil {
		logrus.Error(err)
		response.Error(w, customError.ErrInvalidBodyRequest)
		return
	}

	user := &entity.User{
		Username:    params.Username,
		Email:       params.Email,
		Password:    params.Password,
		FirstName:   params.FirstName,
		LastName:    params.LastName,
		PhoneNumber: params.PhoneNumber,
		RoleID:      params.RoleID,
		CreatedBy:   params.CreatedBy,
	}

	result, err := dashboard.service.InsertUser(ctx, user)
	if err != nil {
		logrus.Error(err)
		response.Error(w, customError.ErrInternalServerError)
		return
	}

	response.Success(w, result)
}

func (dashboard *userDashboard) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()

	id := router.Param(r.Context(), "id")

	var params updateRequest

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&params)

	params.UpdatedBy = r.Header.Get("userId")
	err := helper.Validate(params)
	if err != nil {
		logrus.Error(err)
		response.Error(w, customError.ErrInvalidBodyRequest)
		return
	}

	user := &entity.User{
		Username:    params.Username,
		Email:       params.Email,
		FirstName:   params.FirstName,
		LastName:    params.LastName,
		PhoneNumber: params.PhoneNumber,
		RoleID:      params.RoleID,
		UpdatedBy:   &params.UpdatedBy,
	}

	result, err := dashboard.service.UpdateUser(ctx, user, id)
	if err != nil {
		logrus.Error(err)
		response.Error(w, err)
		return
	}

	response.Success(w, result)
}

func (dashboard *userDashboard) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()

	id := router.Param(r.Context(), "id")

	var params deleteRequest

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&params)

	params.DeletedBy = r.Header.Get("userId")

	err := helper.Validate(params)
	if err != nil {
		logrus.Error(err)
		response.Error(w, customError.ErrInvalidBodyRequest)
		return
	}

	err = dashboard.service.DeleteUser(ctx, id, params.DeletedBy)
	if err != nil {
		logrus.Error(err)
		response.Error(w, err)
		return
	}

	response.Success(w, "SUCCESS")
}
