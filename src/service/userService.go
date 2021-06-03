package service

import (
	"buriedPoint/src/models/basic_fields"
	mysql2 "buriedPoint/src/pkg/mysql"
	"github.com/gin-gonic/gin"
	"net/http"
)

var result basic_fields.Result

func UserLogin(ctx *gin.Context) basic_fields.Result {
	//捕获异常
	defer func() {
		err := recover()
		if err != nil {
			result.Code = http.StatusBadRequest
			result.Message = "错误1"
			result.Data = err
			return
		}
	}()
	user := basic_fields.User{}
	ctx.ShouldBind(&user)
	tx := mysql2.Db.Raw("SELECT id,name,phoneNumber,carid,cellid,avatar FROM user WHERE username = ? and password = ?", user.Username,user.Password).Scan(&user)
	if tx.Error != nil {
		result.Code = http.StatusBadRequest
		result.Message = "错误"
		result.Data = tx.Error
		return result
	}
	result.Code = http.StatusOK
	result.Message = "查询成功"
	result.Data = user
	return result
}

func UserRegister(ctx *gin.Context) basic_fields.Result {
	//捕获异常
	defer func() {
		err := recover()
		if err != nil {
			result.Code = http.StatusBadRequest
			result.Message = "错误1"
			result.Data = err
			return
		}
	}()
	user := basic_fields.User{}
	ctx.ShouldBind(&user)
	if len(user.Avatar) == 0 {
		user.Avatar = "https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
	}
	tx := mysql2.Db.Exec("INSERT INTO user(name,username,password,phoneNumber,avatar) VALUES (?,?,?,?,?);",user.Name,user.Username,user.Password,user.PhoneNumber,user.Avatar)
	if tx.Error != nil {
		result.Code = http.StatusBadRequest
		result.Message = "错误"
		result.Data = tx.Error
		return result
	}
	result.Code = http.StatusOK
	result.Message = "插入成功"
	result.Data = tx.RowsAffected
	return result
}