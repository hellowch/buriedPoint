package service

import (
	"buriedPoint/src/models/basic_fields"
	"buriedPoint/src/models/kafka"
	"buriedPoint/src/models/mysql"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func BPInsertDeploy(ctx *gin.Context) (result basic_fields.Result) {
	buriedPoint := basic_fields.BuriedPoint{}
	err := ctx.ShouldBind(&buriedPoint)
	if err != nil {
		result.Code = http.StatusBadRequest
		result.Message = "模型绑定失败"
		result.Data = err
	}
	buriedPoint.CreateTime = time.Now()
	buriedPoint.UpdateTime = time.Now()
	saveRes, err := mysql.BuriedPointInsertSql(ctx, &buriedPoint)
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

func BPSelectDeploy(ctx *gin.Context) (result basic_fields.Result) {
	buriedPoint := []basic_fields.BuriedPoint{}
	companyId := ctx.DefaultQuery("company_id", "nil")
	userId := ctx.DefaultQuery("user_id", "nil")
	id := ctx.DefaultQuery("id", "nil")
	name := ctx.DefaultQuery("name", "nil")
	BpField := ctx.DefaultQuery("bp_field", "nil")
	var saveRes bool
	var err error
	if companyId != "nil" {
		saveRes, err = mysql.BPSelectDeployCompany(&buriedPoint, companyId)
	} else if userId != "nil" {
		saveRes, err = mysql.BPSelectDeployUser(&buriedPoint, userId)
	} else {
		saveRes, err = mysql.BPSelectDeployOne(&buriedPoint, id, name, BpField)
	}
	if !saveRes {
		result.Code = -100
		result.Message = "错误"
		result.Data = err
		return result
	}
	result.Code = http.StatusOK
	result.Message = "查询成功"
	result.Data = buriedPoint
	return result
}

func BPDeleteDeploy(ctx *gin.Context) (result basic_fields.Result) {
	id := ctx.DefaultQuery("id", "nil")
	saveRes, err := mysql.BPDeleteDeployOne(id)
	if !saveRes {
		result.Code = -100
		result.Message = "错误"
		result.Data = err
		return result
	}
	result.Code = http.StatusOK
	result.Message = "查询成功"
	result.Data = saveRes
	return result
}

//将数据写入kafka
func BPInsertKafkaData(ctx *gin.Context) (result basic_fields.Result) {
	_ = ctx.Request.ParseForm()
	//获取输入的参数
	dataMap := map[string]string{}
	for k, v := range ctx.Request.PostForm {
		dataMap[k] = v[0]
	}
	//写入kafka
	kafka.KafkaProducer(dataMap)

	result.Code = http.StatusOK
	result.Message = "插入成功"
	result.Data = dataMap
	return result
}

//从kafka获取数据写入mongo
//func BPInsertMongoData(value []byte) {
//	dataMap := make(map[string]string)
//	//json转map
//	err := json.Unmarshal(value, &dataMap)
//	if err != nil {
//		fmt.Println("Umarshal failed:", err)
//		return
//	}
//	//写入mongo
//	err = mongo.InsertMongo(dataMap)
//	if err != nil {
//		fmt.Println("mongo failed:", err)
//		return
//	}
//	fmt.Println(dataMap)
//}

