package service

import (
	"buriedPoint/src/models/basic_fields"
	"buriedPoint/src/models/mysql"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func BuriedPointInsert(ctx *gin.Context) (result basic_fields.Result) {
	buriedPoint := basic_fields.BuriedPoint{}
	err := ctx.ShouldBind(&buriedPoint)
	if err != nil {
		result.Code = http.StatusBadRequest
		result.Message = "模型绑定失败"
		result.Data = err
	}
	buriedPoint.CreateTime = time.Now()
	buriedPoint.UpdateTime = time.Now()
	saveRes, err := mysql.BuriedPointInsertSql(ctx, &buriedPoint)
	if !saveRes {
		result.Code = -100
		result.Message = "错误"
		result.Data = err
		return result
	}
	result.Code = http.StatusOK
	result.Message = "插入成功"
	result.Data = saveRes
	return result

}
