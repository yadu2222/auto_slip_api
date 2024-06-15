package model

type EmployeeType struct {
	EmployeeTypeId int `xorm:"'employee_type_id' pk autoincr" json:"employeeTypeId"`
	EmployeeType string `xorm:"'employee_type' not null" json:"employeeType"`
}

func (EmployeeType) TableName() string {
	return "employee_types"
}

// データ
func CreateEmployeeTypeData() {
	employeeType1 := &EmployeeType{
		EmployeeType: "店長",
	}
	db.Insert(employeeType1)
	employeeType2 := &EmployeeType{
		EmployeeType: "アルバイト",
	}
	db.Insert(employeeType2)
}