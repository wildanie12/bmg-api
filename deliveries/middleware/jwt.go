package middleware

import (
	"fmt"
	"gin-bmg-restful/entities/web"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// VerifyJWT makes sure that every request through here
// will always checked for it's authentication status
func VerifyJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		prefix := "Bearer"
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, web.ErrorResponse {
				Status: "ERROR",
				Code: http.StatusUnauthorized,
				Error: "Unauthorized user, you must be logged in first",
			})
			c.Abort()
		}

		tokenString := authHeader[len(prefix) + 1:]
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte("thisissecret33"), nil
		})
		if err != nil {
			panic(err)
		}

		if token.Valid {
			claims, _ := token.Claims.(jwt.MapClaims)
			c.Set("auth_username", claims["username"])
			c.Set("auth_email", claims["email"])
			c.Next()
			return
		} 

		c.JSON(http.StatusUnauthorized, web.ErrorResponse {
			Status: "ERROR",
			Code: http.StatusUnauthorized,
			Error: "Unauthorized user, you must be logged in first",
		})
		c.Abort()
	}
}