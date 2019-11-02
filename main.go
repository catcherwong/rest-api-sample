package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	//swaggerFiles "github.com/swaggo/files"

	//swaggerFiles "github.com/swaggo/files"
	//"github.com/swaggo/gin-swagger"
	"log"

	"github.com/catcherwong/bgadmin-rest-api/api"
	"github.com/catcherwong/bgadmin-rest-api/db"
	"github.com/catcherwong/bgadmin-rest-api/middlewares"
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
	log.Println("welcome to bgadmin-rest-api")

	defer db.RedisClient.Close()

	//r := gin.New()
	//
	//url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r := gin.Default()

	r.Use(middlewares.AuthMiddleware())

	// enable cors
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	r.Use(cors.New(config))

	api.InitRouters(r)

	r.Run(":9999")
}
