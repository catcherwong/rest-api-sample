package routers

import (
	"fmt"
	"github.com/catcherwong/rest-api-sample/config"
	"github.com/catcherwong/rest-api-sample/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "github.com/catcherwong/rest-api-sample/docs"
)

func InitRouter() *gin.Engine {

	gin.SetMode(config.AppCfg.GinMode)

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	initSwagger(r)

	r.GET("/metrics", metrics)

	r.Use(middlewares.AuthMiddleware())

	// enable cors
	initCors(r)

	initApi(r)

	return r
}

func metrics(c *gin.Context) {
	promhttp.Handler().ServeHTTP(c.Writer, c.Request)
}

func initSwagger(e *gin.Engine) {
	url := fmt.Sprintf("http://localhost:%d/swagger/doc.json", config.AppCfg.Port)
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL(url)))
}

func initCors(e *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	e.Use(cors.New(config))
}
