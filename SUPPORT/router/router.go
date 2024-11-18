package router

import (
	"SUPPORT/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())         //允许跨域请求
	loadGroup := r.Group("/user") //设定路由组
	{
		loadGroup.POST("/add", controllers.AddUser)
		loadGroup.DELETE("/delete/:Userid", controllers.DeleteUser)
		loadGroup.POST("/update", controllers.UpdateUser)
		loadGroup.POST("/load", controllers.LoadUser)
	}
	return r
}
