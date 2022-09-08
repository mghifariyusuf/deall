package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

// DBConfig Database Config
type DBConfig struct {
	Dialect         string        `envconfig:"DIALECT"`
	Host            string        `envconfig:"HOST"`
	Port            int           `envconfig:"PORT"`
	Name            string        `envconfig:"NAME"`
	Username        string        `envconfig:"USER_NAME"`
	Password        string        `envconfig:"PASSWORD"`
	MaxConnOpen     int           `envconfig:"MAX_CONN_OPEN"`
	MaxConnLifetime time.Duration `envconfig:"MAX_CONN_LIFETIME"`
	MaxConnIdle     int           `envconfig:"MAX_CONN_IDLE"`
}

// InitDB for Initialize Database SQL
func InitDB(config DBConfig) (db *sql.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", config.Username, config.Password, config.Host, config.Port, config.Name)
	db, err = sql.Open(config.Dialect, dsn)
	if err != nil {
		logrus.Error(err)
		return
	}

	db.SetMaxOpenConns(config.MaxConnOpen)
	db.SetMaxIdleConns(config.MaxConnIdle)
	db.SetConnMaxLifetime(config.MaxConnLifetime)

	return
}
