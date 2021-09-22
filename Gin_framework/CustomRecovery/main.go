package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//what will server does what panic is happened
//just shutdown the server???
//or recovery the server state???
//let custom the recovery the behavior

func main(){
	server := gin.New()
	server.Use(gin.Logger())
	//here we're not using gin recovery
	//custom our self

	//custom gin recovery and use it with middleware
	server.Use(gin.CustomRecovery(func(c *gin.Context,r interface{}){
		fmt.Println(r)
		if err , ok:= r.(string);ok{
			//out put the panic message
			//and recovery server status
			c.String(http.StatusInternalServerError,fmt.Sprintf("error: %s",err))
		} //cast interface to string
		c.AbortWithStatus(http.StatusInternalServerError)
	}))


	//simulate the panic
	server.GET("/panic", func(c *gin.Context) {
		fmt.Println("testing custom recovery")
		panic("foo")
	})

	server.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK,"Hello!Welcome to panic world!")
	})

	server.Run(":8080")
}
