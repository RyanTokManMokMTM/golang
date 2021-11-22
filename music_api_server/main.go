package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"music_api_server/config"
	"music_api_server/route"

	swaggerFile "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "music_api_server/docs"
)


// @title music api server
// @version 1.0
// @description  IOS Music Web Service

// @contact.name jackson.tmm
// @contact.url https://github.com/RyanTokManMokMTM
// @contact.email RyanTokManMokMTM@hotmaiol.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @schemes http



func main() {

	server := gin.New()
	//TODO - API DOC - DEBUG MODE ONLY
	if gin.Mode() == gin.DebugMode{
		url := ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json",config.Server.Port))
		server.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFile.Handler,url))
	}

	//TODO -Serving Static files
	resource := server.Group("/resource")
	resource.Static("/", "./public")

	route.RouterInit(server) //init all available route
	log.Fatalln(server.Run(fmt.Sprintf("%s:%d", config.Server.Address, config.Server.Port)))
}

