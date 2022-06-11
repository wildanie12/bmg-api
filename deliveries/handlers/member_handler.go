package handlers

import (
	"fmt"
	_helpers "gin-bmg-restful/deliveries/helpers"
	_web "gin-bmg-restful/entities/web"
	_userService "gin-bmg-restful/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserHandler main struct
type UserHandler struct {
	userService *_userService.Service
}

// NewUserHandler act as a constructor that return UserHandler 
func NewUserHandler(userService *_userService.Service) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}


// FindAll returns list of user based on query param as a filter
func (handler UserHandler) FindAll(c *gin.Context) {

	filter := []map[string]string{}

	name := c.Query("name")
	if name != "" {
		filter = append(filter, map[string]string{
			"field": "name",
			"operator": "LIKE",
			"value": "%"+name+"%",
		})
	}

	userRes, err := handler.userService.FindAll(filter)
	if err != nil {
		_helpers.WebErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, _web.SuccessResponse{
		Status: "OK",
		Code: http.StatusOK,
		Data: userRes,
	})
}


// CheckReferral return the validity of a referral code
// it also check with currently authenticated user
func (handler UserHandler) CheckReferral(c *gin.Context) {
	request := _web.ReferralRequest{}
	c.Bind(&request)

	username := c.MustGet("auth_username").(string)
	owner, err := handler.userService.CheckReferral(request, username)
	if err != nil {
		_helpers.WebErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, _web.SuccessResponse{
		Status: "OK",
		Code: http.StatusOK,
		Data: owner,
	})
}

// UpdateProfile users
func (handler UserHandler) UpdateProfile(c *gin.Context) {
	request := _web.UserRequest{}
	c.Bind(&request)

	username := c.MustGet("auth_username").(string)
	fmt.Println(username)
	user, err := handler.userService.UpdateProfile(request, username)
	if err != nil {
		_helpers.WebErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, _web.SuccessResponse{
		Status: "OK",
		Code: http.StatusOK,
		Data: map[string]interface{} {
			"username": user.Username,
		},
	})
}