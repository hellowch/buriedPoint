package basic_fields

import "time"

type Company struct {
	Id          int       `json:"id"`
	CompanyName string    `json:"company_name"`//公司名
	CreateTime  time.Time `json:"create_time"`
	UpdateTime  time.Time `json:"update_time"`
}