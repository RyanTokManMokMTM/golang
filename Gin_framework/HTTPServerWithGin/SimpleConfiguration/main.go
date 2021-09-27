package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/
func main(){
	//with go package net/http
	ginRoute := gin.New()
	ginRoute.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK,"welcome to http configure")
	})
	http.ListenAndServe(":8080",ginRoute) //serveMux
}
