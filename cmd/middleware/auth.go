package middleware

import (
	"context"
	"deall/cmd/entity"
	"deall/cmd/lib/customError"
	"deall/pkg/auth"
	"deall/pkg/response"
	"deall/pkg/router"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

type contextKey string
type AuthMiddleware interface {
	RequiresAccessToken(next http.Handler) http.Handler
}

type authMidlleware struct {
	authModule auth.Auth
}

func New(authModule auth.Auth) *authMidlleware {
	return &authMidlleware{authModule: authModule}
}

func (c contextKey) String() string {
	return "myPackage context key " + string(c)
}

func (a *authMidlleware) RequiresAccessToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			logrus.Error(err)
			response.Error(w, err)
			return
		}

		claims, err := auth.ExtractTokenMetadata(r)
		if err != nil {
			logrus.Error(err)
			response.Error(w, err)
			return
		}
		fmt.Println(claims.TokenUUID)
		userID, err := a.authModule.FetchAuth(claims.TokenUUID)
		if err != nil {
			logrus.Error(err)
			response.Error(w, customError.ErrInternalServerError)
			return
		}

		id := router.Param(r.Context(), "id")

		if id != "" && claims.RoleID != "admin" && id != claims.UserID {
			response.Error(w, customError.ErrNotAuthorize)
			return
		}

		r.Header.Add("userId", userID)
		r.Header.Add("token", claims.TokenUUID)

		if r.Method == "PUT" {
			var user entity.User
			decoder := json.NewDecoder(r.Body)
			decoder.Decode(&user)

			if claims.RoleID != "admin" {
				user.RoleID = "user"
			}

			b, err := json.Marshal(user)
			if err != nil {
				log.Println(err)
				response.Error(w, customError.ErrInvalidBodyRequest)
				return
			}
			r.Body = io.NopCloser(strings.NewReader(string(b)))
		}
		contextClaims := contextKey("claims")
		ctx := context.WithValue(r.Context(), contextClaims, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RequiresAuthorization(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contextClaims := contextKey("claims")
		value := r.Context().Value(contextClaims).(*auth.AccessDetails)
		roleID := value.RoleID

		if roleID != "admin" {
			logrus.Error(customError.ErrNotAuthorize.Detail)
			response.Error(w, customError.ErrNotAuthorize)
			return
		}

		next.ServeHTTP(w, r)

	})
}
