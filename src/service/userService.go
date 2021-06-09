package service

import (
	"buriedPoint/src/middleware"
	"buriedPoint/src/models/basic_fields"
	"buriedPoint/src/models/mysql"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)


func UserLogin(ctx *gin.Context) (result basic_fields.Result) {
	out := map[string]interface{}{}
	user := basic_fields.User{}
	user.Username = ctx.PostForm("username")
	user.Password = ctx.PostForm("password")
	saveRes, err := mysql.UserLoginSql(ctx, &user)
	if !saveRes {
		result.Code = -100
		result.Message = "错误"
		result.Data = err
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
	saveRes, err := mysql.UserRegisterSql(ctx, &user)
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
