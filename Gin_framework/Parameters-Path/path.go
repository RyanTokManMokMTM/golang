package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	server := gin.Default()
	//matching path example

	//matching path : /users/(anyArgument)
	server.GET("/users/:name",func(c *gin.Context){
		//getting the name param
		name := c.Param("name")  // get the key of the url
		c.String(http.StatusOK,"your name is :%s",name)
	})

	//example 2
	//matching path : /users/(anyArgument)/ and /users/(anyArgument)/(anyAction)
	server.GET("/users/:name/*action",func(c *gin.Context){
		name := c.Param("name")
		action := c.Param("action")
		c.String(http.StatusOK,"%s is now doing an action of : %s",name,action)
	})

	//each mathing request context will hold the fullPath as route definition
	server.POST("/users/:name/*action",func(c *gin.Context){
		path := c.FullPath() == "/users/:name/*action" //is true due to matching data
		c.String(http.StatusOK,"%t",path)
	})


	//new route other than /users/:name
	server.GET("users/group",func(c *gin.Context){
		c.String(http.StatusOK,"here is a group of users ~~")
	})

	server.Run(":8080")
}
