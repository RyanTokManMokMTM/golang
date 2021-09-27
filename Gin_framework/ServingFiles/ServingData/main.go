package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	server := gin.Default()
	server.GET("/local/file", func(c *gin.Context){
		//writing the specific file to body
		c.File("../public/pic.jpg")
	})

	var fs http.FileSystem = http.Dir("../public")
	server.GET("/local/FS", func(c *gin.Context) {
		//write the file from http.FileSystem to body in efficiency way
		c.FileFromFS("/b.jpg",fs)
	})
	server.Run(":8080")
}
