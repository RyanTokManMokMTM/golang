package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main(){
	//gin.DisableConsoleColor()
	gin.ForceConsoleColor()

	//file,err := os.Create("server.log")
	//if err != nil{
	//	log.Fatalln(err)
	//}
	//
	//gin.DefaultWriter = io.MultiWriter(file)

	server := gin.New()
	server.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string{
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s\"%s \" %s\"\n",
						param.ClientIP,
						param.TimeStamp.Format(time.RFC1123),
						param.Method,
						param.Path,
						param.Request.Proto,
						param.StatusCode,
						param.Latency,
						param.Request.UserAgent(),
						param.ErrorMessage,
			)
	}))
	server.Use(gin.Recovery()) //default one return 500 code only
	//server.Use(gin.CustomRecovery(func(c *gin.Context, recovery interface{}) {
	//	if err,ok := recovery.(string);ok{
	//		c.String(http.StatusInternalServerError,fmt.Sprintf("error : %s ",err))
	//	}
	//	//code = 500 StatusInternalServerError
	//	c.Status(http.StatusInternalServerError)
	//}))
	server.GET("/panicTest", func(c *gin.Context) {
		panic("test panic")
	})

	server.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK,"Pong")
	})
	server.Run(":8080")
}
