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

	//company
	company := Engine.Group("/company")
	company.POST("/register", controller.CompanyRegister)
	company.GET("/select", controller.CompanySelect)

	//user 路由组
	user := Engine.Group("/user")
	user.POST("/login",controller.UserLogin)
	user.POST("/register",controller.UserRegister)

	//buried_point
	buried_point := Engine.Group("/buriedPoint")
	buried_point.POST("/insertDeploy", controller.BPInsertDeploy)
	buried_point.GET("/selectDeploy", controller.BPSelectDeploy)
	buried_point.GET("/deleteDeploy", controller.BPDeleteDeploy)

	Engine.Run(":3001")
}
