package middleware

import (
	"urlshortner/src/authuser"
	"github.com/gin-gonic/gin"
)

var authPackage = new(authuser.controller)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authPackage.TokenValid(c)
		c.Next()
	}
}
