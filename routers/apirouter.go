package routeruser

import (
	"deploysimple/controllers"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init()  {
	router = gin.Default()
	router.GET("/getlistuser", controllers.HandleGetListRequest)
}

