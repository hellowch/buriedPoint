package mysql

import (
	"buriedPoint/src/models/basic_fields"
	mysql2 "buriedPoint/src/pkg/mysql"
	"fmt"
	"github.com/gin-gonic/gin"
)

func BuriedPointInsertSql(ctx *gin.Context, buriedPoint *basic_fields.BuriedPoint) (bool,error) {
	tx := mysql2.Db.Exec("INSERT INTO buried_point(Name,user_id,bp_field,company_id,business_line,layer_name,numeric_field,create_time,update_time) VALUES (?,?,?,?,?,?,?,?);",
		buriedPoint.Name,buriedPoint.UserId,buriedPoint.BpField,buriedPoint.CompanyId,buriedPoint.BusinessLine,buriedPoint.LayerName,buriedPoint.NumericField,buriedPoint.CreateTime,buriedPoint.UpdateTime)
	if tx.Error != nil {
		return false, tx.Error
	}
	return true,nil
}

func BPSelectDeployCompany(buriedPoint *[]basic_fields.BuriedPoint,companyId string) (bool,error)  {
	tx := mysql2.Db.Raw("select * from buried_point where company_id = ?", companyId).Find(&buriedPoint)
	if tx.Error != nil {
		return false, tx.Error
	}
	return true,nil
}

func BPSelectDeployUser(buriedPoint *[]basic_fields.BuriedPoint, userId string) (bool,error)  {
	tx := mysql2.Db.Raw("select * from buried_point where user_id = ?", userId).Find(&buriedPoint)
	if tx.Error != nil {
		return false, tx.Error
	}
	return true,nil
}

func BPSelectDeployOne(buriedPoint *[]basic_fields.BuriedPoint, id string, name string, BpField string) (bool,error)  {
	if id != "nil" {
		tx := mysql2.Db.Raw("select * from buried_point where id = ?", id).Find(&buriedPoint)
		if tx.Error != nil {
			return false, tx.Error
		}
	} else if name != "nil" {
		tx := mysql2.Db.Where(fmt.Sprintf("name LIKE %q",("%"+name+"%"))).Find(&buriedPoint)
		if tx.Error != nil {
			return false, tx.Error
		}
	} else {
		//tx := mysql2.Db.Raw("select * from buried_point where bp_field = ?", BpField).Find(&buriedPoint)
		tx := mysql2.Db.Where(fmt.Sprintf("bp_field LIKE %q",("%"+BpField+"%"))).Find(&buriedPoint)
		if tx.Error != nil {
			return false, tx.Error
		}
	}
	return true,nil
}

func BPDeleteDeployOne(id string) (bool,error) {
	tx := mysql2.Db.Exec("DELETE FROM buried_point WHERE id = ?", id)
	if tx.Error != nil {
		return false, tx.Error
	}
	return true,nil
}
