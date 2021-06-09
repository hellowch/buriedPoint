package mysql

import (
	"buriedPoint/src/models/basic_fields"
	mysql2 "buriedPoint/src/pkg/mysql"
	"github.com/gin-gonic/gin"
)

func BuriedPointInsertSql(ctx *gin.Context, buriedPoint *basic_fields.BuriedPoint) (bool,error) {
	tx := mysql2.Db.Exec("INSERT INTO buried_point(Name,user_id,company_id,business_line,layer_name,numeric_field,create_time,update_time) VALUES (?,?,?,?,?,?,?,?);",
		buriedPoint.Name,buriedPoint.UserId,buriedPoint.CompanyId,buriedPoint.BusinessLine,buriedPoint.LayerName,buriedPoint.NumericField,buriedPoint.CreateTime,buriedPoint.UpdateTime)
	if tx.Error != nil {
		return false, tx.Error
	}
	return true,nil
}
