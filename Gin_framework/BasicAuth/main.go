package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var secrets = gin.H{
	"ken":gin.H{
		"email":"adminA@admin.com",
		"group":"admin",
		"phone":"666",
	},
	"john":gin.H{
		"email":"adminB@admin.com",
		"group":"admin",
		"phone":"777",
	},
	"jackson":gin.H{
		"email":"adminC@admin.com",
		"group":"admin",
		"phone":"808",
	},
}

func main(){
	server := gin.Default()
	adminGroup := server.Group("./admin",gin.BasicAuth(gin.Accounts{
		//used as account and password
		//ac:ken and pw:adminA
		//then return ken as user
		"ken":"adminA",
		"john":"adminB",
		"tom":"no",
		"jackson":"adminC",
	}))

	adminGroup.GET("/secrets", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user];ok{
			c.JSON(http.StatusOK,gin.H{
				"user":user,
				"secretData":secret,
			})
		}else{
			c.JSON(http.StatusOK,gin.H{
				"user":user,
				"secretData":"no secret",
			})
		}
	})
	server.Run(":8080")
}
