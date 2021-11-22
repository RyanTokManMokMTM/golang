package middleware

import (
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
	"music_api_server/Tool"
	"net/http"
	"strings"
)

func JWTAuth() gin.HandlerFunc{
	return func(ctx* gin.Context){
		authHeader := ctx.Request.Header.Get("authorization")
		if authHeader == ""{
			ctx.JSON(http.StatusOK,gin.H{
				"code":http.StatusBadRequest,
				"msg":"token is not found",
			})
			ctx.Abort()
			return
		}

		jwt := Tool.NewJsonWebToken()

		//parse token into jwt object
		tokenStr := strings.Split(authHeader, " ")
		if tokenStr[0] != "Bearer" && tokenStr[1] == ""{
			ctx.JSON(http.StatusOK,gin.H{
				"code":http.StatusBadRequest,
				"msg":"token error",
			})
			ctx.Abort()
			return
		}
		claimsInfo, err := jwt.ParseToken(tokenStr[1])
		if err != nil {
			ctx.JSON(http.StatusOK,gin.H{
				"Code":http.StatusUnauthorized,
				"msg":err.Error(),
			})
			ctx.Abort()
			return
		}
		Tool.Logger.WithFields(logs.Fields{
			"Client":ctx.ClientIP(),
		}).Info("authenticating....")

		ctx.Set("userName",claimsInfo.Name)
		ctx.Set("userEmail",claimsInfo.Name)
		ctx.Next()
	}
}
