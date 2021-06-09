package controller

import (
	"buriedPoint/src/models/basic_fields"
	"buriedPoint/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)


func CompanyRegister(ctx *gin.Context)  {
	result := basic_fields.Result{}
	result = service.CompanyRegister(ctx)
	ctx.JSON(http.StatusOK, result)
}

func CompanySelect(ctx *gin.Context)  {
	result := basic_fields.Result{}
	result = service.CompanySelect(ctx)
	ctx.JSON(http.StatusOK, result)
}
