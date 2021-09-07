package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func MyLogger() gin.HandlerFunc{
	return gin.LoggerWithFormatter(func(p gin.LogFormatterParams) string{
		return fmt.Sprintf("IP:%s,Mehod:%s,Paht:%s",p.ClientIP,p.Method,p.Path)
	})
}