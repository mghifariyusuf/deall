package config

import (
	"deall/pkg/database"
	"deall/pkg/env"
	"deall/pkg/server"
	"log"
)

const prefix = "PROJECT"

var ACCESSSECRET = "SECRET"
var REFRESHSECRET = "SECRET"

// Config For DB
type Config struct {
	ServerConfig  server.Options       `envconfig:"SERVER"`
	DBConfig      database.DBConfig    `envconfig:"DATABASE"`
	RedisConfig   database.RedisConfig `envconfig:"REDIS"`
	accessSecret  string               `envconfig:"ACCESS_SECRET"`
	refreshSecret string               `envconfig:"REFRESH_SECRET"`
}

// LoadConfig for load all configuration from .env file
func LoadConfig() *Config {
	var config Config
	err := env.Load(prefix, &config)
	if err != nil {
		log.Fatalf("Failed to load configuration. err: " + err.Error())
	}
	ACCESSSECRET = config.accessSecret
	REFRESHSECRET = config.refreshSecret
	log.Println("Config successfully Loaded")
	return &config
}
