package handlers

import (
	"gin-bmg-restful/deliveries/helpers"
	_heroService "gin-bmg-restful/services/hero"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HeroHandler main struct object
type HeroHandler struct {
	heroService *_heroService.Service
}

// NewHeroHandler act as a constructor that return HeroHandler 
func NewHeroHandler(heroService *_heroService.Service) *HeroHandler {
	return &HeroHandler{
		heroService: heroService,
	}
}

// FindAll heroes by name search query
func (handler HeroHandler) FindAll(c *gin.Context) {
	name := c.Query("name")

	heroes, err := handler.heroService.FindAll(name)
	if err != nil {
		helpers.WebErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, heroes)
}