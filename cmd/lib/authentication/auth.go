package authentication

import (
	"regexp"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func SetPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	return string(hash), nil
}

func ComparePassword(hashPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(plainPassword))
	if err != nil {
		return false
	}
	return true
}

func IsEmail(s string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(s)
}
