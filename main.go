package main

import (
	_config "gin-bmg-restful/config"
	_handlers "gin-bmg-restful/deliveries/handlers"
	"gin-bmg-restful/deliveries/routes"
	_userRepository "gin-bmg-restful/repositories/user"
	_authService "gin-bmg-restful/services/auth"
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

	// Services
	authService := _authService.NewService(userRepository)
	_userService.NewService(userRepository)

	// Handler
	authHandler := _handlers.NewAuthHandler(authService)

	// Route
	routes.RegisterUserRoute(r, *authHandler)
	
	r.Run(":" + config.App.Port)
}