package model



// groupeテーブル
// typeで型定義
type Employee struct {
	// カラムを指定しない場合は、変数名がそのままカラム名になる
	// pk primaryKey
	// autoincr 自動インクリメント
	// json json化する際のキー名
	EmployeeUuid string  `xorm:"varchar(36) pk" json:"enployeeUUId"`
	EmployeeName string  `xorm:"varchar(20) not null" json:"employeeName"`
	EmployeeTypeId int  `xorm:"not null" json:"employeeType"`
}

// テストデータ
func CreateEmployeeTestData() {
	employee1 := &Employee{
		EmployeeUuid: "c99cb6c4-42b9-4d6b-9884-ae6664f9df00",
		EmployeeName: "やづ",
		EmployeeTypeId: 2,
	}
	db.Insert(employee1)
}

// テーブル名
func (Employee) TableName() string {
	return "employees"
}

// FK制約の追加
func InitEmployeeFK() error {
	// UserTypeId
	_, err := db.Exec("ALTER TABLE employees ADD FOREIGN KEY (employee_type_id) REFERENCES employee_types(employee_type_id) ON DELETE CASCADE ON UPDATE CASCADE")
	if err != nil {
		return err
	}

	return nil
}
