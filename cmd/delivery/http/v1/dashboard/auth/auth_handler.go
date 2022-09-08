package http

import (
	"context"
	"deall/cmd/lib/authentication"
	"deall/cmd/lib/customError"
	"deall/pkg/auth"
	"deall/pkg/helper"
	"deall/pkg/response"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

func (dashboard *authDashboard) Login(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()
	var data string
	var params loginRequest

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&params)

	err := helper.Validate(params)
	if err != nil {
		logrus.Error(err)
		response.Error(w, customError.ErrInvalidBodyRequest)
		return
	}

	if params.Email != "" {
		if authentication.IsEmail(params.Email) == true {
			data = params.Email
		} else {
			logrus.Error("error body request")
			response.Error(w, customError.ErrInvalidBodyRequest)
			return
		}
	} else if params.PhoneNumber != "" {
		data = params.PhoneNumber
	} else {
		logrus.Error("error body request")
		response.Error(w, customError.ErrInvalidBodyRequest)
		return
	}

	res, err := dashboard.service.Login(ctx, data, params.Password)
	if err != nil {
		logrus.Error(err)
		response.Error(w, err)
		return
	}

	response.Success(w, res)
}

func (dashboard *authDashboard) Logout(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()

	accessDetail := auth.AccessDetails{
		TokenUUID: r.Header.Get("token"),
		UserID:    r.Header.Get("userId"),
	}
	fmt.Println(accessDetail.TokenUUID)
	err := dashboard.service.Logout(ctx, &accessDetail)
	if err != nil {
		logrus.Error(err)
		response.Error(w, customError.ErrInternalServerError)
		return
	}

	response.Success(w, "SUCCESS")
}
