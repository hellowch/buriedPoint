package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var Db *gorm.DB

func init() {
	var err error
	dsn := "root:123456@tcp(weichenhao.cn:3310)/chargeMs?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "tb_", // 表名前缀，`User` 的表名应该是 `tb_users`
			SingularTable: true,  // 使用单数表名
		},
	})
	if err != nil {
		panic(err.Error())
		return
	}

	//连接池配置
	sqlDB, err2 := Db.DB()
	if err2 != nil {
		panic(err2.Error())
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}