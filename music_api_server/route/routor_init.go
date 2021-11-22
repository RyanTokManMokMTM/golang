package route

import (
	"github.com/gin-gonic/gin"
)

// RouterInit - /api/v1/path....
func RouterInit(r *gin.Engine){
	//Init for all available route
	apiV1 := r.Group("/api/v1")
	UserRoute(apiV1) //User route
	//r.Use(middleware.JWTAuth()).GET("/test/jwt" ,func(context *gin.Context) {
	//	//name := context.Request.Header.Get("name")
	//	userName := context.MustGet("userName")
	//	//fmt.Println(email)
	//	context.JSON(http.StatusOK, gin.H{"message":userName})
	//})

}
