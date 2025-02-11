package model

type MethodType struct {
	MethodId   int    `xorm:"'method_id' pk autoincr" json:"methodId"`
	MethodType string `xorm:"'method_type' not null" json:"methodType"`
}

func (MethodType) TableName() string {
	return "method_types"
}

// TODO: テストデータ作成
func CreateMethodTypeTestData() {
	methodType1 := &MethodType{
		MethodType: "配達",
	}
	methodType2 := &MethodType{
		MethodType: "店取",
	}
	methodType3 := &MethodType{
		MethodType: "店取伝票",
	}
	methodType4 := &MethodType{
		MethodType: "図書館",
	}
	methodType5 := &MethodType{
		MethodType: "暁光高校",
	}
	methodType6 := &MethodType{
		MethodType: "丸長",
	}
	methodType7 := &MethodType{
		MethodType: "図書館スポンサー",
	}
	db.Insert(methodType1)
	db.Insert(methodType2)
	db.Insert(methodType3)
	db.Insert(methodType4)
	db.Insert(methodType5)
	db.Insert(methodType6)
	db.Insert(methodType7)
}
