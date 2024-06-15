package model

type TellType struct {
	TellTypeId int64 `xorm:"'tell_type_id' pk autoincr" json:"tellTypeId"`
	TellType string `xorm:"'tell_type' not null" json:"tellType"`
}

func (TellType) TableName() string {
	return "tell_types"
}