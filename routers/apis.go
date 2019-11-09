package routers

import (
	"github.com/catcherwong/rest-api-sample/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func initApi(e *gin.Engine) {

	e.GET("/", index)
	e.GET("/ping", pong)

	initUserController(e)
	initRedisController(e)
	initMetricsController(e)
	initMockController(e)
}

func pong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func index(c *gin.Context) {
	c.String(http.StatusOK, "welcome to rest-api-sample")
}

func initUserController(e *gin.Engine) {

	uc := controllers.NewUserController()

	rg := e.Group("/api/v1/users")
	{
		rg.GET("/", uc.GetUserList)
		rg.GET("/:id", uc.GetUserById)
		rg.POST("/", uc.AddUser)
	}
}

func initRedisController(e *gin.Engine) {

	c := controllers.NewRedisController()

	rg := e.Group("/api/v1/redis")
	{
		rg.GET("/string", c.GetString)
		rg.POST("/string", c.SetString)
		rg.DELETE("/", c.DeleteValue)
	}
}

func initMetricsController(e *gin.Engine) {

	c := controllers.NewMetricsController()

	rg := e.Group("/api/v1/metrics")
	{
		rg.GET("/", c.Record)
	}
}

func initMockController(e *gin.Engine) {

	c := controllers.NewMockController()

	rg := e.Group("/api/mock")
	{
		rg.GET("/", c.GetString)
		rg.GET("/get1", c.GetString1)
		rg.GET("/get2", c.GetString2)
	}
}
