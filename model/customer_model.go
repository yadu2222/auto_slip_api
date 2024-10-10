package model

import (
	"log"
)

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
	CustomerUuid string `xorm:"varchar(36) pk" json:"customerUUId"` // 一意の値
	CustomerName string `json:"customerName"`                       // 顧客名
	Ruby 		 string `json:"ruby"`                       // 顧客名のルビ
	MethodType   int    `json:"methodType"`                         // 処理のタイプ
	TellAddress  string `json:"tellAddress"`                        //電話番号	// unique検討
	TellType     int    `json:"tellType"`                           // 冊数
	Note         string `json:"note"`                               // 冊数
	CsvId        int    `json:"csvId"`                              // csv形式のときに使用していたid　nullを許容
}

func (Customer) TableName() string {
	return "customers"
}

// FK制約の追加
func InitCustomerFK() error {

	_, err := db.Exec("ALTER TABLE customers ALTER COLUMN tell_address SET DEFAULT NULL")
	if err != nil {
		return err
	}

	// // methodtype
	_, err = db.Exec("ALTER TABLE customers ADD FOREIGN KEY (method_type) REFERENCES method_types(method_id) ON DELETE CASCADE ON UPDATE CASCADE")
	if err != nil {
		return err
	}

	// // telltype
	_, err = db.Exec("ALTER TABLE customers ADD FOREIGN KEY (tell_type) REFERENCES tell_types(tell_type_id) ON DELETE CASCADE ON UPDATE CASCADE")
	if err != nil {
		return err
	}
	return nil
}

// テストデータ
func CreateCustomerTestData() {
	customer1 := &Customer{
		CustomerUuid: "d38678b7-b540-4893-96aa-a3f51cbb07f2",
		CustomerName: "ほげ岡",
		MethodType:   1,
		TellAddress:  "090-1234-5678",
		TellType:     1,
		Note:         "ほげほげ鳴いてます",
	}
	db.Insert(customer1)
}

// 顧客一覧を取得
func GetCustomers() ([]Customer, error) {
	var customers []Customer
	err := db.Find(&customers)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

// 名前で検索して取得
func FindCustomersByName(name string) ([]Customer, error) {
	var customers []Customer
	session := db.Table("customers")
	err := session.Where("customer_name like ?", "%"+name+"%").Find(&customers)
	if err != nil {
		return customers, err
	}
	return customers, nil
}

// お客様を登録する関数
func RegisterCustomer(customer Customer) error {
	// // 電話番号の重複を確認
	// exists, err := isCustomerExists(customer)

	_, err := db.Insert(customer)
	if err != nil {
		return err
	}
	return nil
}

// お客様を登録する関数
func RegisterCustomers(customers []Customer) error {
	for _, customer := range customers {

		if customer.TellAddress != "" {
			exists, err := ExistsustomerByCsvID(customer.CsvId)
			if err != nil {
				// エラーが発生した場合、ログを出力して処理を継続
				log.Printf("%s様の重複チェック中にエラーが発生しました: %v", customer.CustomerName, err)
				return err
			}
			if exists {
				// 重複がある場合はログを出力して処理を継続
				log.Printf("%s様はすでに登録されています", customer.CustomerName)
				continue
			}
		}

		// お客様を登録
		_, err := db.Insert(&customer)
		if err != nil {
			// エラーが発生した場合、ログを出力して処理を継続します
			log.Printf("%s様の登録中にエラーが発生しました: %v", customer.CustomerName, err)
			return err
		}
		log.Printf("%s様を登録しました", customer.CustomerName)
	}
	return nil
}

// 指定された顧客がすでに存在するかを電話番号でチェックする関数
func IsCustomerExists(customer Customer) (bool, error) {
	session := db.Where("tell_address = ?", customer.TellAddress)
	count, err := session.Count(&Customer{})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// uuidから顧客を取得
func FindCustomerByID(uuid string) (Customer, error) {
	var customer Customer
	session := db.Table("customers")
	_, err := session.Where("customer_uuid = ?", uuid).Get(&customer)
	if err != nil {
		return customer, err
	}
	return customer, nil
}

// csv_idから顧客を取得
func FindCustomerByCsvID(csvid string) (Customer, error) {
	var customer Customer
	session := db.Table("customers")
	_, err := session.Where("csv_id = ?", csvid).Get(&customer)
	if err != nil {
		return customer, err
	}
	return customer, nil
}

// uuidから顧客のidをチェック

// 指定された顧客がすでに存在するかをチェックする関数
func ExistsustomerByCsvID(csvid int) (bool, error) {
	var customer Customer
	session := db.Table("customers")
	// データが存在するかどうかをチェック
	exists, err := session.Where("csv_id = ?", csvid).Get(&customer)
	if err != nil {
		return false, err
	}
	if exists {
		// 顧客が存在する場合
		log.Println("顧客名:", customer.CustomerName)
		return true, nil
	}
	// 顧客が存在しない場合
	return false, nil
}

// お客様を削除する関数
func DeleteCustomer(customerUuid string) (*Customer, error) {
	customer, err := FindCustomerByID(customerUuid)
	if err != nil {
		return nil, err
	}
	_, err = db.Delete(&customer)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

// 顧客情報を更新
func UpdateCustomer(customer Customer) error {
	_, err := db.Where("customer_uuid = ?", customer.CustomerUuid).Update(&customer)
	if err != nil {
		return err
	}
	return nil
}

// 配達が必要な顧客を取得
func FindCustomersNeedDelivery() ([]Customer, error) {
	var customers []Customer
	session := db.Table("customers")
	err := session.Where("method_type = 1").Find(&customers)
	if err != nil {
		return nil, err
	}
	return customers, nil
}
