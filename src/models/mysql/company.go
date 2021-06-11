package mysql

import (
	"buriedPoint/src/models/basic_fields"
	mysql2 "buriedPoint/src/pkg/mysql"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CompanyRegisterSql(ctx *gin.Context, company *basic_fields.Company) (bool,error) {
	tx := mysql2.Db.Exec("INSERT INTO company(company_name,create_time,update_time) VALUES (?,?,?);",company.CompanyName,company.CreateTime,company.UpdateTime)
	if tx.Error != nil {
		return false, tx.Error
	}
	return true,nil
}

func CompanySelectSql(ctx *gin.Context, companyLike string, company *[]basic_fields.Company) (bool,error) {
	tx := mysql2.Db.Where(fmt.Sprintf("company_name LIKE %q",("%"+companyLike+"%"))).Find(&company)
	if tx.Error != nil {
		return false, tx.Error
	}
	return true,nil
}
