package route

import (
	"github.com/gin-gonic/gin"
	"music_api_server/apiError"
	"music_api_server/controller/user_controller"
)

func UserRoute(r *gin.RouterGroup){
	userRoute := r.Group("/user")
	//userRoute.POST("/files",apiError.ErrorHandler(user_controller.UploadImage))
	////USER RESOURCE [GET]
	//userAuthedRoute := userRoute.Use(middleware.JWTAuth())
	//userAuthedRoute.GET("/profile")

	//USER RESOURCE [GET]


	//POST REQUEST
	authPath := userRoute.Group("/auth")
	authPath.POST("/signup", apiError.ErrorHandler(user_controller.RegisterHandler))
	authPath.POST("/login",apiError.ErrorHandler(user_controller.LoginHandler))

}

