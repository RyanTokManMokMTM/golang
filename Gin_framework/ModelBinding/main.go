package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

/*
Data Binding with the request
There are 2 methods -MustBind(Gin handle error) and ShouldBind(Developer handle error itself)
set the binding to data fields ,add a `json:"fieldname"` on data model
Type A: Must Bind
	`Bind`,`BindJSON`,`BingXML`,`Bindxxx` -> all use MustBindWith under the hood
	=> these binding method(MushBindWith),if there is a binding error,the request is aborted with
       `context.AbortWithError(400,err)`,the response status code is set to 400 anc content-type is set to
		`text/plain;character-set="utf-8"`
		DON'T TRY TO CHANGE RESPONSE HEADER(STATUS CODE) => If u need to, change the binding type to
		`ShouldBind`
Type B: Should Bind
		`ShouldBind`,`ShouldBindJSON`... => use should bind with under the hood
		if there is a binding error, the error is returned and it's the developer's responsibility
		to handle the request and error appropriately
		HANDLE ERROR OURSELVES
*/

/*
Gin tries to infer the binder depending on the Content-type
If we're sure what we are binding, we can use `ShouldBindWith`/`MustBindWith`

we can also specify fields are required => `binding:"required"`
if the required field has empty value, an error will be return
*/

//if the field no need to valid -> binding:"-" to skip the validation
type userLogin struct {
	Account string  `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main(){
	//try bind with json xml and from
	server := gin.Default()
	user := server.Group("/login")
	user.Use(func(c *gin.Context) {
		fmt.Println("new user is came in")
		c.Next()
	})
	{
		user.POST("/json",bindJSONHandle)
		user.POST("/xml",bindXMLHandle)
		user.POST("/form",bindFormHandle)
	}
	server.Run(":8080")
}

/*
Format
JSON{
	user : "admin",
	pw : "1324"
}
*/
func bindJSONHandle(c *gin.Context){
	var json userLogin
	fmt.Println(c.Request.Body)
	if err := c.ShouldBindJSON(&json); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":"missing data",
			"error": err.Error(),
		})
		return
	}

	//check json value
	fmt.Println(json)
	if json.Account != "admin" || json.Password != "1234"{
		c.JSON(http.StatusUnauthorized,gin.H{
			"message":"Unauthorized",
		})
		return
	}
	fmt.Println(json)

	c.JSON(http.StatusOK,"welcome!")
}

func bindXMLHandle(c *gin.Context){
	xml := userLogin{}
	fmt.Println("XML binding...")
	if err := c.ShouldBindXML(&xml); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":"missing data",
			"error": err.Error(),
		})
		return
	}

	if strings.TrimSpace(xml.Account) != "admin" || xml.Password != "1234"{
		c.JSON(http.StatusUnauthorized,gin.H{
			"message":"Unauthorized",
		})
		return
	}

	fmt.Println(xml)

	c.String(http.StatusOK,"Welcome")
}

func bindFormHandle(c *gin.Context){
	form := userLogin{}
	fmt.Println("FORM binding...")
	//Check the content-type
	if err := c.ShouldBind(&form); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":"missing data",
			"error":err.Error(),
		})
		return
	}
	fmt.Println(form)
	if form.Account != "admin" || form.Password != "1234"{
		c.JSON(http.StatusUnauthorized,gin.H{
			"message":"Unauthorized",
		})
		return
	}

	fmt.Println(form)

	c.String(http.StatusOK,"welcome!")
}