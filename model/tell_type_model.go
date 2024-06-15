package model

type TellType struct {
	TellTypeId int64 `xorm:"'tell_type_id' pk autoincr" json:"tellTypeId"`
	TellType string `xorm:"'tell_type' not null" json:"tellType"`
}

func (TellType) TableName() string {
	return "tell_types"
}

// TODO: テストデータ作成
func CreateTellTypeTestData(){
	tellType1 := &TellType{
		TellType: "不要",
	}
	tellType2 := &TellType{
		TellType: "要",
	}
	tellType3 := &TellType{
		TellType: "着信のみ",
	}
	db.Insert(tellType1)
	db.Insert(tellType2)
	db.Insert(tellType3)
	
}