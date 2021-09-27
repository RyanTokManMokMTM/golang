package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main(){
	router := gin.New()
	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK,"http configuration for more detail")
	})
	//server configuration
	server := http.Server{
		Addr: ":8080",
		Handler: router,
		ReadHeaderTimeout: 2 * time.Second,
		ReadTimeout: 2*time.Second,
		WriteTimeout: 2 * time.Second,
		IdleTimeout: 5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
