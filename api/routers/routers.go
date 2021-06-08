package routers

import (
	"buriedPoint/src/controller"
	"buriedPoint/src/middleware"
	"github.com/gin-gonic/gin"
)

var Engine *gin.Engine

func InitRouter()  {
	Engine := gin.Default()
	Engine.Use(middleware.RequestInfos(),middleware.Cors())

	//user 路由组
	user := Engine.Group("/user")
	user.POST("/login",controller.UserLogin)
	user.POST("/register",controller.UserRegister)

	Engine.Run(":3001")
}