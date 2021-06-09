package controller

import (
	"buriedPoint/src/models/basic_fields"
	"buriedPoint/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)


func UserLogin(ctx *gin.Context)  {
	result := basic_fields.Result{}
	result = service.UserLogin(ctx)
	ctx.JSON(http.StatusOK, result)
}

func UserRegister(ctx *gin.Context)  {
	result := basic_fields.Result{}
	result = service.UserRegister(ctx)
	ctx.JSON(http.StatusOK, result)
}
