package model

// "fmt"
// "math/rand"
// "time"
// TODO: 外部キー制約
// groupeテーブル
// typeで型定義
type Regular struct {
	// カラムを指定しない場合は、変数名がそのままカラム名になる
	// pk primaryKey
	// autoincr 自動インクリメント
	// json json化する際のキー名
	RegularUuid  string `xorm:"varchar(36) pk" json:"regularUUID"` // 一意の値
	MagazineUuid string `xorm:"varchar(36)" json:"magazineUUID"`                      // 雑誌ID
	CustomerUuid string `json:"customerUUID"`                      // 顧客コード
	Quantity     int    `json:"magazine_name"`                     // 冊数
}

func (Regular) TableName() string {
	return "regulars"
}

// FK制約の追加
func InitRegularFK() error {
	// magazine_uuid
	_, err := db.Exec("ALTER TABLE regulars ADD FOREIGN KEY (magazine_uuid) REFERENCES magazines(magazine_uuid) ON DELETE CASCADE ON UPDATE CASCADE")
	if err != nil {
		return err
	}
	// customer_uuid
	_, err = db.Exec("ALTER TABLE regulars ADD FOREIGN KEY (customer_uuid) REFERENCES customers(customer_uuid) ON DELETE CASCADE ON UPDATE CASCADE")
	if err != nil {
		return err
	}

	return nil
}
