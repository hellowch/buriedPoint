package controller

import (
	"buriedPoint/src/models/basic_fields"
	"buriedPoint/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//公司注册
func CompanyRegister(ctx *gin.Context)  {
	result := basic_fields.Result{}
	result = service.CompanyRegister(ctx)
	ctx.JSON(http.StatusOK, result)
}

//公司查询
func CompanySelect(ctx *gin.Context)  {
	result := basic_fields.Result{}
	result = service.CompanySelect(ctx)
	ctx.JSON(http.StatusOK, result)
}

//公司注销
func CompanyLogout(ctx *gin.Context)  {

}

//公司信息修改
func CompanyUpdate(ctx *gin.Context)  {

}

