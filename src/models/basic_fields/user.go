package basic_fields

import "time"

type User struct {
	Id          int       `form:"id"           json:"id"            gorm:"column:id;primary_key"`
	Name        string    `form:"name"         json:"name"          gorm:"column:name"`
	Username    string    `form:"username"     json:"username"      gorm:"column:username"`
	Password    string    `form:"password"     json:"password"      gorm:"column:passowrd"`
	PhoneNumber string    `form:"phone_number" json:"phone_number"  gorm:"column:phone_number"`
	RoleId      int       `form:"role_id"      json:"role_id"       gorm:"column:role_id"` //账号权限10:管理员 1可操作 2只读
	CompanyId   int       `form:"company_id"   json:"company_id"    gorm:"column:company_id"`
	Avatar      string    `form:"avatar"       json:"avatar"        gorm:"column:avatar"`
	Email       string    `form:"email"        json:"email"         gorm:"column:email"`
	Available   int       `form:"available"    json:"available"     gorm:"column:available"` //帐号状态1可用2不可用
	CreateTime  time.Time `form:"create_time"  json:"create_time"   gorm:"column:create_time"`
	UpdateTime  time.Time `form:"update_time"  json:"update_time"   gorm:"column:update_time"`
}
