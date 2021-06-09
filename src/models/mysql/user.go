package mysql

import (
	"buriedPoint/src/models/basic_fields"
	mysql2 "buriedPoint/src/pkg/mysql"
	"github.com/gin-gonic/gin"
)

func UserLoginSql(ctx *gin.Context, user *basic_fields.User) (bool,error) {
	tx := mysql2.Db.Raw("SELECT id,name,username,password,phone_number,role_id,company_id,avatar,email,available,create_time,update_time FROM user WHERE username = ? and password = ?", user.Username,user.Password).Scan(&user)
	if tx.Error != nil {
		return false, tx.Error
	}
	return true,nil
}

func UserRegisterSql(ctx *gin.Context, user *basic_fields.User) (bool,error) {
	tx := mysql2.Db.Exec("INSERT INTO user(name,username,password,phone_number,role_id,company_id,avatar,email,create_time,update_time) VALUES (?,?,?,?,?,?,?,?,?,?);",
		user.Name,user.Username,user.Password,user.PhoneNumber,user.RoleId,user.CompanyId,user.Avatar,user.Email,user.CreateTime,user.UpdateTime)
	if tx.Error != nil {
		return false, tx.Error
	}
	return true,nil
}
