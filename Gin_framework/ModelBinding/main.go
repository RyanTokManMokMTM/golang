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

type User struct {
	Name string `form:"name"`
	Address string `form:"addr"`
	School string `form:"school"`
}

type Character struct {
	ID string `uri:"id" binding:"required"`
	Name string `uri:"name" binding:"required"`
}

//access the header data
type HeaderData struct {
	Domain string `header:"domain"`
	CookiesData bool `header:"cookie"`

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

	server.Any("/testing",bindUriQueryHandle)
	server.GET("bindUri/:id/:name",bindWithUriHandle)
	server.GET("/customHeader",bindWithHeaderHandle)

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

//not working with body etc... just working with uri query string
//?xxx=xxx&yyy=yyy...
func bindUriQueryHandle(c *gin.Context){
	//Only bind the query string that passing in uri
	var user User
	if err := c.ShouldBindQuery(&user);err != nil{
		c.JSON(http.StatusBadRequest,err)
		return
	}
	fmt.Println(user)
	c.String(http.StatusOK,"success")
}

//Bind model with define params in uri
//behind the scene -> range the c.Params to get our param
func bindWithUriHandle(c *gin.Context){
	fmt.Println(c.Param("id"))
	var character Character
	if err := c.ShouldBindUri(&character);err != nil{
		c.String(http.StatusBadRequest,err.Error())
		return
	}

	fmt.Println(character)
	c.JSON(http.StatusOK,gin.H{
		"message":"succeed",
		"uuid":character.ID,
		"name":character.Name,
	})
}

//bind with some header data that client has set to header
/* example
{
    "Domain": "game",
    "hasCookie": true,
    "message": "accessed"
}
 */
func bindWithHeaderHandle(c *gin.Context){
	var headerData HeaderData
	if err := c.ShouldBindHeader(&headerData);err != nil{
		c.String(http.StatusBadRequest,err.Error())
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"message":"accessed",
		"Domain":headerData.Domain,
		"hasCookie":headerData.CookiesData,
	})
}