package auth

import (
	"deall/cmd/config"
	"deall/cmd/lib/customError"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"github.com/twinj/uuid"
)

type tokenService struct{}

func NewToken() *tokenService {
	return &tokenService{}
}

type Token interface {
	CreateToken(userID string, roleID string) (*TokenDetails, error)
	ExtractTokenMetadata(r *http.Request) (*AccessDetails, error)
}

func (t *tokenService) CreateToken(userID string, roleID string) (*TokenDetails, error) {
	token := &TokenDetails{}
	token.RoleID = roleID
	token.TokenExpires = time.Now().Add(time.Minute * 30).Unix()
	token.TokenUUID = uuid.NewV4().String()

	var err error

	accessTokenClaims := jwt.MapClaims{}
	accessTokenClaims["access_uuid"] = token.TokenUUID
	accessTokenClaims["role_id"] = token.RoleID
	accessTokenClaims["user_id"] = userID
	accessTokenClaims["exp"] = token.TokenExpires

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	token.AccessToken, err = accessToken.SignedString([]byte(config.ACCESSSECRET))
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	token.RefreshTokenExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	token.RefreshUUID = fmt.Sprintf("%s++%s", token.TokenUUID, userID)

	refreshTokenClaims := jwt.MapClaims{}
	refreshTokenClaims["refresh_uuid"] = token.RefreshUUID
	refreshTokenClaims["role_id"] = token.RoleID
	refreshTokenClaims["user_id"] = userID
	refreshTokenClaims["exp"] = token.RefreshTokenExpires
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	token.RefreshToken, err = refreshToken.SignedString([]byte(config.REFRESHSECRET))
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return token, nil
}

func (t *tokenService) ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := verifyToken(r)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	acc, err := extract(token)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return acc, nil
}

func verifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			logrus.Error(token.Header["alg"])
			return nil, customError.ErrUnexpectedSigning
		}
		return []byte(config.ACCESSSECRET), nil
	})
	if err != nil {
		logrus.Error(err)
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, customError.ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, customError.ErrToken
			} else {
				return nil, customError.ErrToken
			}
		}
		return nil, customError.ErrToken
	}
	return token, nil
}

func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func extract(token *jwt.Token) (*AccessDetails, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		userID, userOK := claims["user_id"].(string)
		roleID, roleOK := claims["role_id"].(string)
		if !ok || !userOK || !roleOK {
			logrus.Error(customError.ErrNotAuthorize.Detail)
			return nil, customError.ErrNotAuthorize
		}
		return &AccessDetails{
			TokenUUID: accessUUID,
			UserID:    userID,
			RoleID:    roleID,
		}, nil
	}
	return nil, customError.ErrToken
}

func TokenValid(r *http.Request) error {
	token, err := verifyToken(r)
	if err != nil {
		logrus.Error(err)
		return err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		logrus.Error(customError.ErrToken.Detail)
		return customError.ErrToken
	}
	return nil
}

func ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := verifyToken(r)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	acc, err := extract(token)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return acc, nil
}
