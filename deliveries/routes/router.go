package routes

import (
	_handlers "gin-bmg-restful/deliveries/handlers"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoute contains all endpoint registered
// based on user endpoint category needs
func RegisterUserRoute(r *gin.Engine, authHandler _handlers.AuthHandler) {
	r.POST("/auth/login", authHandler.Login)
	r.POST("/auth/register", authHandler.Register)
	r.GET("/auth/me", authHandler.Me)
}