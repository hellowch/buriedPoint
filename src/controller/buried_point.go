package controller

import (
	"buriedPoint/src/models/basic_fields"
	"buriedPoint/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BuriedPointInsert(ctx *gin.Context)  {
	result := basic_fields.Result{}
	result = service.BuriedPointInsert(ctx)
	ctx.JSON(http.StatusOK, result)
}
