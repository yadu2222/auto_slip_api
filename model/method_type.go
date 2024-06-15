package model

type MethodType struct {
	MethodId int `xorm:"'tell_type_id' pk autoincr" json:"tellTypeId"`
	MethodType string `xorm:"'tell_type' not null" json:"tellType"`
}

func (MethodType) TableName() string {
	return "method_types"
}