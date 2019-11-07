package main

import (
	"fmt"
	"github.com/catcherwong/rest-api-sample/config"
	"github.com/catcherwong/rest-api-sample/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"

	"github.com/catcherwong/rest-api-sample/db"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	log.Println("welcome to rest-api-sample")

	defer db.RedisClient.Close()

	gin.SetMode(config.AppCfg.GinMode)

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	initSwagger(r)

	r.GET("/metrics", metrics)

	//r.Use(middlewares.AuthMiddleware())

	// enable cors
	initCors(r)

	//api.InitRouters(r)

	routers.InitRouters(r)

	r.Run(":9999")
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
