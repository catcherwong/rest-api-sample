package api

import (
	"github.com/catcherwong/bgadmin-rest-api/api/redis"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouters(e *gin.Engine) {

	e.GET("/", index)

	v1 := e.Group("/api")
	{
		v1.GET("/ping", pong)

		redis.InitRouters(v1)
	}
}

func pong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func index(c *gin.Context) {
	c.String(http.StatusOK, "welcome to bgadmin-rest-api")
}
