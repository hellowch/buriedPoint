package controller

import (
	"buriedPoint/src/models/basic_fields"
	"buriedPoint/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//写入新建埋点信息
func BPInsertDeploy(ctx *gin.Context)  {
	result := basic_fields.Result{}
	result = service.BPInsertDeploy(ctx)
	ctx.JSON(http.StatusOK, result)
}

//获取埋点列表(公司埋点，个人埋点，某个埋点)
func BPSelectDeploy(ctx *gin.Context)  {
	
}

//埋点删除
func BPDeleteDeploy(ctx *gin.Context)  {
	
}

//向埋点添加数据
func BPInsertData(ctx *gin.Context)  {
	
}

//读取埋点数据
func BPSelectData(ctx *gin.Context)  {

}