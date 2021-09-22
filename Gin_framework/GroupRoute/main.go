package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	server := gin.Default()
	//group with v1 route
	v1 := server.Group("/v1")
	{
		v1.POST("/login", func(c *gin.Context) {
			c.String(http.StatusOK,"welcome!v1")
		})

		v1.POST("/submit", func(c *gin.Context) {
			c.String(http.StatusOK,"summited!v1")
		})
	}

	v2 := server.Group("/v2")
	{
		v2.POST("/login", func(c *gin.Context) {
			c.String(http.StatusOK,"welcome!v2")
		})

		v2.POST("/submit", func(c *gin.Context) {
			c.String(http.StatusOK,"summited!v2")
		})
	}

	//group with v2 route


	server.Run(":8080")
}
