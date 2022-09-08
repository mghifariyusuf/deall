package cache

import (
	"errors"
	"testing"

	"github.com/gomodule/redigo/redis"
	"github.com/rafaeljusto/redigomock"
	"github.com/stretchr/testify/assert"
)

var connection = redigomock.NewConn()
var pool = NewConn(connection)
var conn = Init(pool)

func NewConn(conn redis.Conn) *redis.Pool {

	pool := &redis.Pool{
		Dial:    func() (redis.Conn, error) { return conn, nil },
		MaxIdle: 10,
	}
	return pool
}

func TestSet(t *testing.T) {
	key := `test`
	value := "test"

	t.Run("success", func(t *testing.T) {
		connection.Command("SET", key, value)

		err := conn.Set(key, value)
		assert.Nil(t, err)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		connection.Command("SET", key, value).ExpectError(errors.New("Error"))

		err := conn.Set(key, value)
		assert.NotNil(t, err)
		assert.Error(t, err)
	})
}

func TestGet(t *testing.T) {
	key := `test`

	t.Run("success", func(t *testing.T) {
		connection.Command("GET", key).Expect("test")

		value, err := conn.Get(key)
		assert.Nil(t, err)
		assert.NotNil(t, value)
		assert.Equal(t, value, []byte("test"))
	})

	t.Run("not-found", func(t *testing.T) {
		connection.Command("GET", key).Expect(nil)

		value, err := conn.Get(key)
		assert.Nil(t, err)
		assert.Nil(t, value)
	})

	t.Run("error", func(t *testing.T) {
		connection.Command("GET", key).ExpectError(errors.New("error"))

		value, err := conn.Get(key)
		assert.NotNil(t, err)
		assert.Nil(t, value)
	})
}

func TestSetEx(t *testing.T) {
	key := `test`
	value := "test"
	var ttl int64 = FIVEMINUTE

	t.Run("success", func(t *testing.T) {
		connection.Command("SETEX", key, ttl, value)

		err := conn.SetEx(key, ttl, value)
		assert.Nil(t, err)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		connection.Command("SETEX", key, ttl, value).ExpectError(errors.New("error"))

		err := conn.SetEx(key, ttl, value)
		assert.NotNil(t, err)
		assert.Error(t, err)
	})
}

func TestDelete(t *testing.T) {
	key := `test`

	t.Run("success", func(t *testing.T) {
		connection.Command("DEL", key)

		err := conn.Del(key)
		assert.Nil(t, err)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		connection.Command("DEL", key).ExpectError(errors.New("error"))

		err := conn.Del(key)
		assert.NotNil(t, err)
		assert.Error(t, err)
	})
}

func TestExpire(t *testing.T) {
	key := `test`
	expire := 40
	connection.Command("SET", key, "test")

	t.Run("success", func(t *testing.T) {
		connection.Command("EXPIRE", key, expire)

		err := conn.Expire(key, expire)
		assert.Nil(t, err)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		connection.Command("EXPIRE", key, expire).ExpectError(errors.New("error"))

		err := conn.Expire(key, expire)
		assert.NotNil(t, err)
		assert.Error(t, err)
	})
}
