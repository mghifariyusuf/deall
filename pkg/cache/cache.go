package cache

import (
	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
)

type redisPool struct {
	conn *redis.Pool
}

const (
	//FIVEMINUTE Expire Redis Key for Five Minutes
	FIVEMINUTE int64 = 300

	//TENMINUTE Expire Redis Key for Ten Minutes
	TENMINUTE int64 = 600

	//HALFHOUR Expire Redis Key for Half Hour
	HALFHOUR int64 = 1800

	//ONEHOUR Expire Redis Key for One Hour
	ONEHOUR int64 = 3600

	//TWOHOUR Expire Redis Key for Two Hour
	TWOHOUR int64 = 7200

	//SIXHOUR Expire Redis Key for Six Hour
	SIXHOUR int64 = 21600
)

//RedisCommand .
type RedisCommand interface {
	Get(key string) ([]byte, error)
	Set(key string, value string) error
	SetEx(key string, ttl int64, value string) error
	Del(key string) error
	Expire(key string, ttl int) error
}

//Init for initialization cache redis
func Init(conn *redis.Pool) RedisCommand {
	return &redisPool{conn}
}

func (r *redisPool) Set(key string, value string) error {
	conn := r.conn.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (r *redisPool) Get(key string) ([]byte, error) {
	conn := r.conn.Get()
	defer conn.Close()

	v, err := redis.String(conn.Do("GET", key))
	if err == redis.ErrNil {
		return nil, nil
	}

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return []byte(v), nil
}

func (r *redisPool) SetEx(key string, ttl int64, value string) error {
	conn := r.conn.Get()
	defer conn.Close()

	_, err := conn.Do("SETEX", key, ttl, value)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (r *redisPool) Del(key string) error {
	conn := r.conn.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", key)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (r *redisPool) Expire(key string, ttl int) error {
	conn := r.conn.Get()
	defer conn.Close()

	_, err := conn.Do("EXPIRE", key, ttl)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
