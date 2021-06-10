package controller

import (
	"buriedPoint/src/models/basic_fields"
	"buriedPoint/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//用户登录
func UserLogin(ctx *gin.Context)  {
	result := basic_fields.Result{}
	result = service.UserLogin(ctx)
	ctx.JSON(http.StatusOK, result)
}

//用户注册
func UserRegister(ctx *gin.Context)  {
	result := basic_fields.Result{}
	result = service.UserRegister(ctx)
	ctx.JSON(http.StatusOK, result)
}

//用户注销
func UserLogout(ctx *gin.Context)  {
	
}


//用户信息修改
func UserUpdate(ctx *gin.Context)  {

}