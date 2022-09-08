package cmd

import (
	"deall/cmd/config"
	dashboardAuth "deall/cmd/delivery/http/v1/dashboard/auth"
	dashboardUser "deall/cmd/delivery/http/v1/dashboard/user"
	"deall/cmd/middleware"
	"deall/cmd/repositories"
	"deall/cmd/repositories/user/mysql"
	"deall/cmd/repositories/user/redis"
	service "deall/cmd/service/auth"
	serviceUser "deall/cmd/service/user"
	"deall/pkg/auth"
	"deall/pkg/cache"
	"deall/pkg/database"
	"deall/pkg/server"
	"log"
	"os"
	"os/signal"
	"syscall"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/sirupsen/logrus"
)

func Run() {
	formatter := runtime.Formatter{ChildFormatter: &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "01/02/06 15:04:05",
	}}
	formatter.Line = true
	logrus.SetFormatter(&formatter)
	logrus.SetOutput(os.Stdout)
	logrus.StandardLogger()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	config := config.LoadConfig()
	logrus.Info("Config Loaded")

	// Database
	db, err := database.InitDB(config.DBConfig)
	if err != nil {
		logrus.Error(err)
	}
	logrus.Info("Success Initialize Database")

	// Redis
	redisConn := database.InitRedis(config.RedisConfig)
	defer redisConn.Close()

	redisCommand := cache.Init(redisConn)
	logrus.Info("Success Initialize Redis")

	// Auth
	authModule := auth.New(redisCommand)
	tokenModule := auth.NewToken()

	// middleware
	middlewareModule := middleware.New(authModule)

	// Casbin
	// casbin := casbin.New(db)
	// casbin.Register()

	// Repository
	var authRepository repositories.UserRepository = mysql.NewUser(db)
	authRepository = redis.NewUser(redisCommand, authRepository, config.RedisConfig.Prefix)

	// Service
	var authService = service.NewAuthService(authRepository, authModule, tokenModule)
	var userService = serviceUser.NewUserService(authRepository)

	// Server
	server := server.New(&config.ServerConfig)

	// router
	router := server.Router()

	// add CORS
	router.Use(middleware.RequiresCORS)

	// Handler
	v1 := router.Group("/v1")
	dashboardAuth.NewAuthDashboard(authService, middlewareModule).Register(v1)
	dashboardUser.NewUserDashboard(userService, middlewareModule).Register(v1)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	serveChan := server.Run()

	select {
	case err := <-serveChan:
		if err != nil {
			panic(err)
		}
	case <-sigChan:
	}

	server.Stop()
}
