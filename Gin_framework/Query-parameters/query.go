package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	server := gin.Default()
	server.GET("/hello",func(c *gin.Context){
		//get the Uri parameter
		// /hello?firstName=x&lastName=y

		//if firstName is existed and set it,otherwise set to the default value
		firstName := c.DefaultQuery("firstName","Guest")

		//lastName := c.Request.URL.Query().Get("")
		//if lastName is not exist set it to empty
		lastName := c.Query("lastName")
		c.String(http.StatusOK,"Hello %s %s",firstName,lastName)
	})
	//quest form
	server.POST("/post_form", func(c *gin.Context) {
		//query data from a from
		msg := c.PostForm("message")
		name := c.DefaultPostForm("name","guest")
		c.JSON(http.StatusOK,gin.H{
			"status" : "posted",
			"message" : msg,
			"name" : name,
		})
	})
	//query and form
	server.POST("/posts",func(c *gin.Context){
		//with id and page
		ids := c.Query("ids")
		page := c.DefaultQuery("page","0")
		name := c.DefaultPostForm("name","guest")
		msg := c.PostForm("message")

		c.JSON(http.StatusOK,gin.H{
			"status":"posted",
			"id":ids,
			"page":page,
			"name":name,
			"message":msg,
		})
	})
	//query a map???
	server.POST("/postMap", func(c *gin.Context) {
		// /postMap?ids[dbID]=1&ids[pageID]=2
		//get the ids for the map dbID and PageID
		ids := c.QueryMap("ids") //used for same query strings??

		//get the name from the form for 2 same field name
		//eg: name[first] and name[last]
		names := c.PostFormMap("name")

		c.JSON(http.StatusOK,gin.H{
			"ids":ids,
			"names":names,
		})
	})
	server.Run(":8080")
}
