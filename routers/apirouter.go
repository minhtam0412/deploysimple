package routers

import (
	"deploysimple/controllers"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()
	router.GET("/getall", controllers.HandleGetAllUserRequest)
	router.POST("/add", controllers.HandleAddUserRequest)
	router.PUT("/update", controllers.HandleUpdateUserRequest)
	router.DELETE("/delete", controllers.HandleDeleteUserRequest)
	router.GET("/get/:id", controllers.HandleGetUserByKey)
	router.GET("/ping", controllers.Ping)
}

func Run(port string) {
	router.Run(port)
}
