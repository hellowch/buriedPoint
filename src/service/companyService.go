package service

import (
	"buriedPoint/src/models/basic_fields"
	"buriedPoint/src/models/mysql"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CompanyRegister(ctx *gin.Context) (result basic_fields.Result) {
	company := basic_fields.Company{}
	err := ctx.ShouldBind(&company)
	if err != nil {
		result.Code = http.StatusBadRequest
		result.Message = "模型绑定失败"
		result.Data = err
	}
	company.CreateTime = time.Now()
	company.UpdateTime = time.Now()
	saveRes, err := mysql.CompanyRegisterSql(ctx, &company)
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

func CompanySelect(ctx *gin.Context) (result basic_fields.Result) {
	company := []basic_fields.Company{}
	companyLike := ctx.DefaultQuery("company_name", "nil")
	saveRes, err := mysql.CompanySelectSql(ctx, companyLike, &company)
	if !saveRes {
		result.Code = -100
		result.Message = "错误"
		result.Data = err
		return result
	}
	result.Code = http.StatusOK
	result.Message = "查询成功"
	result.Data = company
	return result
}
