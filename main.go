package main

import (
	"github.com/catcherwong/rest-api-sample/common"
	"github.com/catcherwong/rest-api-sample/routers"
	"log"

	_ "github.com/catcherwong/rest-api-sample/docs"
)

// @title rest-api-sample
// @version 1.0
// @description This is a sample about rest api .
// @termsOfService https://github.com/catcherwong

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @contact.name Catcher Wong
// @contact.url https://github.com/catcherwong
// @contact.email catcher_hwq@outlook.com

// @license.name MIT
// @license.url https://github.com/catcherwong/rest-api-sample/blob/master/LICENSE

// @schemes http
// @BasePath /
func main() {
	log.Println("welcome to rest-api-sample")

	common.Close()

	r := routers.InitRouter()

	r.Run(":9999")
}
