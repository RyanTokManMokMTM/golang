package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func MyMiddleWare() gin.HandlerFunc{
	return func(ctx *gin.Context){
		fmt.Println("Hello my middleware")
		ctx.Next()
	}
}
