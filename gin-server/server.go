package main

import (
	img "git-server/example/controller/imageController"
	myLog "git-server/example/middleware"
	"github.com/gin-gonic/gin"
)

func main(){
	server := gin.Default()

	//setting the static
	server.Static("/resource","./resource")
	server.StaticFile("/sharefile","./shareFile/test.png")
	server.Use(myLog.MyMiddleWare())
	//create a controller
	imgController := img.NewImageController()

	//define a group of controller
	imageRoute := server.Group("/image")
	imageRoute.Use(myLog.MyLogger())
	imageRoute.GET("/",imgController.GetAll)
	imageRoute.POST("/",imgController.Create)
	imageRoute.PUT("/:id",imgController.Update)
	imageRoute.DELETE("/:id",imgController.Delete)

	server.Run(":8080")
}
