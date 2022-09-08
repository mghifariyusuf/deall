package database

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
)

// RedisConfig is Configuration for Redis
type RedisConfig struct {
	Host          string `envconfig:"HOST"`
	Port          int    `envconfig:"PORT"`
	Prefix        string `envconfig:"PREFIX" default:"project"`
	MaxIdle       int    `envconfig:"MAX_IDLE"`
	MaxIdleActive int    `envconfig:"MAX_ACTIVE"`
	auth          string `envconfig:"AUTH"`
	db            string `envconfig:"DATABASE"`
}

// InitRedis Initialize Redis Pool
func InitRedis(cfg RedisConfig) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     cfg.MaxIdle,
		IdleTimeout: 240 * time.Second,
		MaxActive:   cfg.MaxIdleActive,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
			if err != nil {
				logrus.Error(err)
				return nil, err
			}
			// _, err = conn.Do("AUTH", cfg.auth)
			// if err != nil {
			// 	conn.Close()

			// 	logrus.Error(err)
			// 	return nil, err
			// }
			// _, err = conn.Do("SELECT", cfg.db)
			// if err != nil {
			// 	conn.Close()

			// 	return nil, err
			// }
			return conn, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
