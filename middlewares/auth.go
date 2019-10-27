package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		p := c.Request.URL.Path

		log.Println("request path ", p)

		if p == "/api/ping" || p == "/" {
			c.Next()
			return
		}

		token := c.Request.Header.Get("Authorization")

		log.Println("request token ", token)

		if token == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if token != "1234567890" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
