package main

import "github.com/gin-gonic/gin"

func main(){
	server := gin.Default()
	server.GET("/getHandler",getHandler)
	server.POST("/postHandler",postHandler)
	server.PUT("/",putHandler)
	server.DELETE("/",deleteHandler)
	server.Run(":8008")

}

func getHandler(context *gin.Context){
	//to get something
}

func postHandler(context *gin.Context){
	//to post something
}

func putHandler(context *gin.Context){
	//to update something
}

func deleteHandler(context *gin.Context){
	//to delete something
}