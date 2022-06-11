package routes

import (
	_handlers "gin-bmg-restful/deliveries/handlers"
	_middleware "gin-bmg-restful/deliveries/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoute contains all endpoint registered
// based on user endpoint category needs
func RegisterUserRoute(r *gin.Engine, authHandler *_handlers.AuthHandler, userHandler *_handlers.UserHandler) {
	r.POST("/auth/login", authHandler.Login)
	r.POST("/auth/register", authHandler.Register)
	r.GET("/auth/me", _middleware.VerifyJWT(), authHandler.Me)
	r.GET("/users", _middleware.VerifyJWT(), userHandler.FindAll)
	r.PUT("/users", _middleware.VerifyJWT(), userHandler.UpdateProfile)
	r.POST("/users/referral", _middleware.VerifyJWT() ,userHandler.CheckReferral)
}


// RegisterHeroRoute registers all hero endpoint
func RegisterHeroRoute(r *gin.Engine, heroHandler *_handlers.HeroHandler) {
	r.GET("/hero", heroHandler.FindAll)
}