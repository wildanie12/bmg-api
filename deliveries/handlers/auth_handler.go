package handlers

import (
	_helpers "gin-bmg-restful/deliveries/helpers"
	_web "gin-bmg-restful/entities/web"
	_authService "gin-bmg-restful/services/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthHandler main struct object
type AuthHandler struct {
	authService *_authService.Service
}

// NewAuthHandler act as a constructor that return AuthHandler 
func NewAuthHandler(authService *_authService.Service) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Login a registered user
// AuthRequest entity must is used
func (handler AuthHandler) Login(c *gin.Context) {
	authRequest := _web.AuthRequest{}
	c.Bind(&authRequest)

	authResponse, err := handler.authService.Login(authRequest)
	if err != nil {
		_helpers.WebErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, _web.SuccessResponse{
		Status: "OK",
		Code: 200,
		Data: authResponse,
	})
}

// Register a new user
// Into the system
func (handler AuthHandler) Register(c *gin.Context) {
	userRequest := _web.UserRequest{}
	c.Bind(&userRequest)

	authResponse, err := handler.authService.Register(userRequest)
	if err != nil {
		_helpers.WebErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, _web.SuccessResponse{
		Status: "OK",
		Code: 200,
		Data: authResponse,
	})
}

// Me method can Get currently authenticated
// Userdata by their JWT token
func (handler AuthHandler) Me(c *gin.Context) {
	
}