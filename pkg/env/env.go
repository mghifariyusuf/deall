package env

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

func Load(prefix string, out interface{}, filename ...string) error {
	err := godotenv.Load(filename...)
	if err != nil {
		logrus.Error(err)
		return err
	}

	err = envconfig.Process(prefix, out)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
