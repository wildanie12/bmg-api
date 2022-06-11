package main

import (
	_config "gin-bmg-restful/config"
	_handlers "gin-bmg-restful/deliveries/handlers"
	"gin-bmg-restful/deliveries/routes"
	_heroRepository "gin-bmg-restful/repositories/hero"
	_userRepository "gin-bmg-restful/repositories/user"
	_authService "gin-bmg-restful/services/auth"
	_heroService "gin-bmg-restful/services/hero"
	_userService "gin-bmg-restful/services/user"
	_utils "gin-bmg-restful/utils"

	"github.com/gin-gonic/gin"
)

func main() {

	config := _config.New()
	db := _utils.NewPostgresConnection(config)
	_utils.Migrate(db)


	r := gin.Default()

	// Repositories
	userRepository := _userRepository.NewRepository(db)
	heroRepository := _heroRepository.NewRepository()

	// Services
	authService := _authService.NewService(userRepository)
	userService := _userService.NewService(userRepository)
	heroService := _heroService.NewService(heroRepository)

	// Handler
	authHandler := _handlers.NewAuthHandler(authService)
	userHandler := _handlers.NewUserHandler(userService)
	heroHandler := _handlers.NewHeroHandler(heroService)

	// Route
	routes.RegisterUserRoute(r, authHandler, userHandler)
	routes.RegisterHeroRoute(r, heroHandler)
	
	r.Run(":" + config.App.Port)
}