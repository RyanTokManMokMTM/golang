package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//use gin.New() -> to create a blank new gin server without any middleware
	server := gin.New()

	//add middleware ourselves
	server.Use(gin.Logger()) //at a logger(default) as middleware to server
	server.Use(gin.Recovery()) //at a Recovery(default) as middleware to server

	//we can add our custom middle ware to any round

	//this route will use 2 custom middleware
	server.GET("/middleware",customMiddleWareA,customMiddleWareB, func(c *gin.Context) {
		c.String(http.StatusOK,"welcome the middleware testing!")
	})
	//grouping the route

	mainRoute := server.Group("/")
	mainRoute.Use(authMiddleWare())
	{
		mainRoute.POST("/login", func(c *gin.Context) {
			fmt.Println("user login...")
			c.JSON(http.StatusOK,gin.H{
				"message":"LoggedIn",
			})
		})

		mainRoute.POST("/submit", func(c *gin.Context) {
			fmt.Println("user submitting...")
			c.JSON(http.StatusOK,gin.H{
				"message":"file submitted",
			})
		})

		//another route group inside current group
		testingRoute := mainRoute.Group("testing")
		testingRoute.GET("/analytics", func(c *gin.Context) {
			fmt.Println("analytics")
			c.String(http.StatusOK,"tested and analyzed")
		})
	}
	server.Run(":8080")
}

func customMiddleWareA(c *gin.Context){
	fmt.Println("i am passing the middleware A")
	// jump to next middle ware
	c.Next()
}

func customMiddleWareB(c *gin.Context){
	fmt.Println("i am passing the middleware B")
	// jump to next middle ware
	c.Next()
}

func auth(c *gin.Context){
	fmt.Println("Hello!You are authorized")
	c.Next()
}

func authMiddleWare() gin.HandlerFunc{
	return auth
}
