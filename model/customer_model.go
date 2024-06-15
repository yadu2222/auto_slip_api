package model

// "fmt"
// "math/rand"
// "time"
// TODO: 外部キー制約
// groupeテーブル
// typeで型定義
type Customer struct {
	// カラムを指定しない場合は、変数名がそのままカラム名になる
	// pk primaryKey
	// autoincr 自動インクリメント
	// json json化する際のキー名
	CustomerUuid string  `xorm:"varchar(36) pk" json:"customerUUId"`// 一意の値
	CustomerName string  `json:"customerName"`	// 雑誌コード
	MethodType int  `json:"methodType"`			// 処理のタイプ
	TellAddress string `json:"tellAddress"`						// 冊数
	TellType int `json:"tellType"`						// 冊数
	Note string `json:"note"`						// 冊数
}

func (Customer) TableName() string {
	return "customers"
}

func RegisterCustomer(customer Customer) error {
	_, err := db.Insert(customer)
	if err != nil {
		return err
	}
	return nil
}

// FK制約の追加
func InitCustomerFK() error {
	// methodtype
	_, err := db.Exec("ALTER TABLE customers ADD FOREIGN KEY (method_type) REFERENCES method_types(method_id) ON DELETE CASCADE ON UPDATE CASCADE")
	if err != nil {
		return err
	}
	// telltype
	_, err = db.Exec("ALTER TABLE customers ADD FOREIGN KEY (tell_type) REFERENCES tell_types(tell_type_id) ON DELETE CASCADE ON UPDATE CASCADE")
	if err != nil {
		return err
	}
	return nil
}

// テストデータ
func CreateCustomerTestData() {
	customer1 := &Customer{
		CustomerUuid:  "d38678b7-b540-4893-96aa-a3f51cbb07f2",
		CustomerName: "ほげ岡",
		MethodType: 1,
		TellAddress: "090-1234-5678",
		TellType: 1,
		Note: "ほげほげ鳴いてます",
	}
	db.Insert(customer1)
}