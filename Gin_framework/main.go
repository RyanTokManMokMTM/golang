package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

//simple gin http server
func main(){
	server := gin.Default() //return a default engine with logger as middleware

	//http method
	server.GET("/", func(context *gin.Context) {
			//send back a josn
			context.JSON(200,gin.H{
				"message" : "hello",
			})
	})

	log.Fatalln(server.Run("localhost:8080"))
}