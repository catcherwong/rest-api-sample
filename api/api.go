package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouters(e *gin.Engine) {

	e.GET("/", index)
	e.GET("/ping", pong)

	redisApi := NewRedisApi()
	redisApi.InitRouter(e)

	userApi := NewUserApi()
	userApi.InitRouter(e)

	metricsApi := NewMetricsApi()
	metricsApi.InitRouter(e)

}

func pong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func index(c *gin.Context) {
	c.String(http.StatusOK, "welcome to rest-api-sample")
}
