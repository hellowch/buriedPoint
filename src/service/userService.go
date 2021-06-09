package service

import (
	"buriedPoint/src/middleware"
	"buriedPoint/src/models/basic_fields"
	mysql2 "buriedPoint/src/pkg/mysql"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)


func UserLogin(ctx *gin.Context) (result basic_fields.Result) {
	out := map[string]interface{}{}
	user := basic_fields.User{}
	user.Username = ctx.PostForm("username")
	user.Password = ctx.PostForm("password")
	tx := mysql2.Db.Raw("SELECT id,name,username,password,phone_number,role_id,company_id,avatar,email,available,create_time,update_time FROM user WHERE username = ? and password = ?", user.Username,user.Password).Scan(&user)
	if tx.Error != nil {
		result.Code = -100
		result.Message = "错误"
		result.Data = tx.Error
		return result
	}
	token, err := middleware.GenToken(user.Username)
	if err != nil {
		result.Code = -101
		result.Message = "token生成错误"
		result.Data = err
		return result
	}
	out["user"] = user
	out["token"] = token
	result.Code = http.StatusOK
	result.Message = "查询成功"
	result.Data = out
	return result
}

func UserRegister(ctx *gin.Context) (result basic_fields.Result) {
	user := basic_fields.User{}
	err := ctx.ShouldBind(&user)
	if err != nil {
		result.Code = http.StatusBadRequest
		result.Message = "模型绑定失败"
		result.Data = err
	}
	if len(user.Avatar) == 0 {
		user.Avatar = "https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
	}
	if user.RoleId == 0 {
		user.RoleId = 1
	}
	user.CreateTime = time.Now()
	user.UpdateTime = time.Now()
	tx := mysql2.Db.Exec("INSERT INTO user(name,username,password,phone_number,role_id,company_id,avatar,email,create_time,update_time) VALUES (?,?,?,?,?,?,?,?,?,?);",
		user.Name,user.Username,user.Password,user.PhoneNumber,user.RoleId,user.CompanyId,user.Avatar,user.Email,user.CreateTime,user.UpdateTime)
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
