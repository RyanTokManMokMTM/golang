package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//set static file
func main(){
	server := gin.Default()
	//localhost:8080/public to access StaticPublic folder

	//public whole folder/resources
	server.Static("/public","./StaticPublic")

	//just can access this files
	server.StaticFile("/picture","./StaticPublic/pic.jpg")

	//working with custom file system
	//gin default use gin.Dir
	//just list StaticFS
	server.StaticFS("/customFS",http.Dir("./StaticPublic"))
	server.Run(":8080")
}
