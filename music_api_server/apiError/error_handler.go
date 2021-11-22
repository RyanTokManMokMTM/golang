package apiError

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorAPIHandle func(ctx *gin.Context) (interface{},error)

func ErrorHandler(handler ErrorAPIHandle) gin.HandlerFunc{
	return func(ctx *gin.Context){
		data, err := handler(ctx)
		if err != nil {
			apiError := err.(APIError)
			ctx.JSON(apiError.Code,apiError)
			return
		}

		ctx.JSON(http.StatusOK,gin.H{
			"code":http.StatusOK,
			"result":data,
		})
	}
}

