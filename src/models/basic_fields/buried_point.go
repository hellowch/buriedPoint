package basic_fields

import "time"

type BuriedPoint struct {
	Id           int        `json:"id"`
	Name         string     `json:"name"`
	UserId       int        `json:"user_id"`
	BusinessLine string     `json:"business_line"`//业务线
	LayerName    string     `json:"layer_name"`//层名
	NumericField string     `json:"numeric_field"`//字段
	CreateTime   time.Time  `json:"create_time"`
	UpdateTime   time.Time  `json:"update_time"`
}