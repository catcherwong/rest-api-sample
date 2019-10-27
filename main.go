package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"

	"github.com/catcherwong/bgadmin-rest-api/api"
	"github.com/catcherwong/bgadmin-rest-api/middlewares"
)

func main() {
	log.Println("welcome to bgadmin-rest-api")

	r := gin.Default()

	r.Use(middlewares.AuthMiddleware())

	// enable cors
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	r.Use(cors.New(config))

	api.InitRouters(r)

	r.Run(":9999")
}
