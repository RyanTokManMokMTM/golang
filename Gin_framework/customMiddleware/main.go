package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func customLogger() (handler gin.HandlerFunc){
	handler = func(c *gin.Context){
		t := time.Now()

		// set a variable to context
		c.Set("userKey","hello")

		c.Next()

		//??? After the request is done???
		//??? After the response????
		//**If the router doesn't return a response to client
		//the context will come back here
		latency := time.Since(t)
		log.Println(latency)

		status := c.Writer.Status()
		log.Println(status)
	}
	return
}

func main(){
	server := gin.New()
	server.Use(customLogger())
	server.GET("/middleware", func(c *gin.Context) {
		//try to get the key that middleware is set
		key := c.MustGet("userKey").(string)
		log.Println(key)
		//c.String(http.StatusOK,"Hello")
	})
	server.Run(":8080")
}
