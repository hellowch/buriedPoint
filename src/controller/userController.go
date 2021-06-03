package controller

import (
	"buriedPoint/src/models/basic_fields"
	"buriedPoint/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var result basic_fields.Result

func UserLogin(ctx *gin.Context)  {
	result = service.UserLogin(ctx)
	ctx.JSON(http.StatusOK, result)
}

func UserRegister(ctx *gin.Context)  {
	result = service.UserRegister(ctx)
	ctx.JSON(http.StatusOK, result)
}