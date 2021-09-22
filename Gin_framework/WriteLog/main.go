package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func main(){
	//some gin setting
	//just disable the ConsoleColor
	//gin default writter is  os.Stdout?
	gin.DisableConsoleColor()

	//create a log file

	file,err := os.Create("server.log")
	if err != nil{
		log.Fatalln(err)
	}

	//set the gin writer is our log file
	//var DefaultWriter io.Writer = os.Stdout
	//change gin default writer to our file
	gin.DefaultWriter = io.MultiWriter(file) //Multiple writer can write it

	server := gin.Default()

	server.GET("/log", func(c *gin.Context) {
		c.String(http.StatusOK,"logged")
	})

	server.Run(":8080")
}
