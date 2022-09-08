package http

import (
	"deall/cmd/middleware"
	"deall/cmd/service"
	"deall/pkg/router"
	"time"
)

var timeout = 5 * time.Second

type authDashboard struct {
	service    service.AuthService
	middleware middleware.AuthMiddleware
}

func NewAuthDashboard(service service.AuthService, middleware middleware.AuthMiddleware) *authDashboard {
	return &authDashboard{service, middleware}
}

func (dashboard *authDashboard) Register(router *router.Router) {
	authRouter := router.Group("/auth")
	authRouter.POST("/login", dashboard.Login)
	authRouter.Use(dashboard.middleware.RequiresAccessToken)
	authRouter.POST("/logout", dashboard.Logout)
}
