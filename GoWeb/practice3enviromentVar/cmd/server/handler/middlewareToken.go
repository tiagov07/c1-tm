package handler

import (
	"github/tiagov07/c1-tm/GoWeb/practice3enviromentVar/pkg/web"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")
	if requiredToken == "" {
		log.Fatal("you must include the token in order to continue")
	}
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			c.AbortWithStatusJSON(401, web.NewResponse(401, nil, "not authorized", "token required"))
			return
		}
		if token != requiredToken {
			c.AbortWithStatusJSON(401, web.NewResponse(401, nil, "not authorized", "invalid token"))
			return
		}
		c.Next()

	}
}
