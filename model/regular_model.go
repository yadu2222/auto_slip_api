package model

import (
	"log"
)

// "fmt"
// "math/rand"
// "time"
// groupeテーブル
// typeで型定義
type Regular struct {
	// カラムを指定しない場合は、変数名がそのままカラム名になる
	// pk primaryKey
	// autoincr 自動インクリメント
	// json json化する際のキー名
	RegularUuid  string `xorm:"varchar(36) pk" json:"regularUUID"` // 一意の値
	MagazineCode string `xorm:"varchar(36)" json:"magazineCode"`   // 雑誌ID
	CustomerUuid string `xorm:"varchar(36)" json:"customerUUID"`                      // 顧客コード
	Quantity     int    `json:"quantity"`                     // 冊数
}

func (Regular) TableName() string {
	return "regulars"
}

// FK制約の追加
func InitRegularFK() error {
	// magazine_code
	_, err := db.Exec("ALTER TABLE regulars ADD FOREIGN KEY (magazine_code) REFERENCES magazines(magazine_code) ON DELETE CASCADE ON UPDATE CASCADE")
	if err != nil {
		return err
	}
	// // customer_uuid
	_, err = db.Exec("ALTER TABLE regulars ADD FOREIGN KEY (customer_uuid) REFERENCES customers(customer_uuid) ON DELETE CASCADE ON UPDATE CASCADE")
	if err != nil {
		return err
	}

	return nil
}

func CreateRegularTestData() {
	regular1 := &Regular{
		RegularUuid:  "903e3147-1b8c-4e26-a5ee-f525a246e2df",
		MagazineCode: "29934",
		CustomerUuid: "d38678b7-b540-4893-96aa-a3f51cbb07f2",
		Quantity:     1,
	}
	db.Insert(regular1)
}



// 定期情報の登録
func RegisterRegular(regular Regular) error {
	exists, err := isRegularExists(regular)
	if err != nil {
		// エラーが発生した場合、ログを出力して処理を継続
		log.Printf("%sの重複チェック中にエラーが発生しました: %v", regular.CustomerUuid,err)
		return err
	}
	if exists {
		// 重複がある場合はログを出力して処理を継続
		log.Printf("%sはすでに登録されています", regular.CustomerUuid)
		return nil
	}

	// 定期を登録
	_, err = db.Insert(&regular)
	if err != nil {
		// エラーが発生した場合、ログを出力して処理を継続します
		log.Printf("%sの登録中にエラーが発生しました: %v", regular.CustomerUuid, err)
		// return err
	}
	log.Printf("%sを登録しました", regular.CustomerUuid)
	return nil
}

// 定期を一括登録
func RegisterRegulars(regulars []Regular) error {
	for _, regular := range regulars {
		exists, err := isRegularExists(regular)
		if err != nil {
			// エラーが発生した場合、ログを出力して処理を継続
			log.Printf("%sの重複チェック中にエラーが発生しました: %v", regular.CustomerUuid, err)
			return err
		}
		if exists {
			// 重複がある場合はログを出力して処理を継続
			log.Printf("%sはすでに登録されています", regular.CustomerUuid)
			continue
		}

		// 定期を登録
		_, err = db.Insert(&regular)
		if err != nil {
			// エラーが発生した場合、ログを出力して処理を継続します
			log.Printf("%sの登録中にエラーが発生しました: %v", regular.CustomerUuid, err)
			// return err
		}else{
			log.Printf("%sを登録しました", regular.CustomerUuid)
		}
	}
	return nil
}

// 指定された雑誌がすでに存在するかをチェックする関数
func isRegularExists(regular Regular) (bool, error) {
	// ここで具体的に雑誌の重複チェックを実装します
	var count int64
	session := db.Where("customer_uuid = ?", regular.CustomerUuid).And("magazine_code = ?", regular.MagazineCode)
	count, err := session.Count(&Regular{})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// 定期情報を一覧取得
func GetRegulars() ([]Regular, error) {
	var regulars []Regular

	err := db.Find(&regulars)
	if err != nil {
		return nil, err
	}
	return regulars, nil
}

// 雑誌を主キーに定期を一覧取得
func FindRegularByMagazine(magazineUuid string) ([]Regular, error) {
	var regulars []Regular
	err := db.Where("magazine_code = (?)", magazineUuid).Find(&regulars)
	if err != nil {
		return nil, err
	}
	return regulars, nil
}

// 顧客を主キーに定期を一覧取得
func FindRegularByCustomer(customerUuid string) ([]Regular, error) {
	var regulars []Regular
	err := db.Where("customer_uuid = (?)", customerUuid).Find(&regulars)
	if err != nil {
		return nil, err
	}
	return regulars, nil
}

