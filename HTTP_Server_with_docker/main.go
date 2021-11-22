package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

var logger *log.Logger = log.New(os.Stdout,"HTTP Server",log.LstdFlags)

func main(){
	 server := gin.New()
	 server.GET("/",homeHandler)

	logger.Println("SERVER IS LISTING ON 0.0.0.0:8080")
	err := server.Run(":8080")
	if err != nil {
		return
	}
	//log.Println("Server is on")
	//logger.Println("SERVER IS LISTING ON 0.0.0.0:8080")

}

func homeHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK,gin.H{
		"message":"hello",
	})
}