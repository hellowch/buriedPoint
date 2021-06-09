package service

import (
	"buriedPoint/src/models/basic_fields"
	mysql2 "buriedPoint/src/pkg/mysql"
	"fmt"
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
	tx := mysql2.Db.Exec("INSERT INTO company(company_name,create_time,update_time) VALUES (?,?,?);",company.CompanyName,company.CreateTime,company.UpdateTime)
	if tx.Error != nil {
		result.Code = http.StatusInternalServerError
		result.Message = "错误"
		result.Data = tx.Error
		return result
	}
	result.Code = http.StatusOK
	result.Message = "插入成功"
	result.Data = tx.RowsAffected
	return result
}

func CompanySelect(ctx *gin.Context) (result basic_fields.Result) {
	company := []basic_fields.Company{}
	companyLike := ctx.DefaultQuery("company_name", "nil")
	tx := mysql2.Db.Debug().Where(fmt.Sprintf("company_name LIKE %q",("%"+companyLike+"%"))).Find(&company)
	if tx.Error != nil {
		result.Code = http.StatusInternalServerError
		result.Message = "错误"
		result.Data = tx.Error
		return result
	}
	result.Code = http.StatusOK
	result.Message = "查询成功"
	result.Data = company
	return result
}
