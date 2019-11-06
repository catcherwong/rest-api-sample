package routers

import (
	"github.com/catcherwong/rest-api-sample/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouters(e *gin.Engine) {

	e.GET("/", index)
	e.GET("/ping", pong)

	initUserController(e)
	initRedisController(e)
	initMetricsController(e)
}

func pong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func index(c *gin.Context) {
	c.String(http.StatusOK, "welcome to rest-api-sample")
}

func initUserController(e *gin.Engine) {

	uc := controllers.NewUserController()

	rg := e.Group("/v1/api/users")
	{
		rg.GET("/", uc.GetUserList)
		rg.GET("/:id", uc.GetUserById)
		rg.POST("/", uc.AddUser)
	}
}

func initRedisController(e *gin.Engine) {

	c := controllers.NewRedisController()

	rg := e.Group("/v1/api/redis")
	{
		rg.GET("/string", c.GetString)
		rg.POST("/string", c.SetString)
		rg.DELETE("/", c.DeleteValue)
	}
}

func initMetricsController(e *gin.Engine) {

	c := controllers.NewMetricsController()

	rg := e.Group("/v1/api/metrics")
	{
		rg.GET("/", c.Record)
	}
}