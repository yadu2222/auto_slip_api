package model

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var db *xorm.Engine // インスタンス

// SQL接続とテーブル作成
func DBConnect() error {
	// 環境変数から取得
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	dbDB := os.Getenv("MYSQL_DATABASE")

	// Mysqlに接続
	var err error
	db, err = xorm.NewEngine( // dbとエラーを取得
		"mysql", // dbの種類"root:root@tcp(db:3306)/cgroup?charset=utf8"
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", dbUser, dbPass, dbHost, dbPort, dbDB), // 接続情報
	)
	if err != nil { // エラー処理
		// エラーが出たら終了する
		log.Fatal("MySQL への接続に失敗しました;~~~;", err)
		return err
	} else {
		log.Println("MySQL への接続に成功しました!!!!!!!:")
	}
	return nil
}

// 接続を取得
func DBInstance() *xorm.Engine {
	return db // 接続を返す
}

// マイグレーション関連
func MigrationTable() error {
	// テーブルがないなら自動で作成。 // xormがテーブル作成時に列名をスネークケースにしてくれる。  // 列情報の追加変更は反映するが列の削除は反映しない。
	exists, _ := db.IsTableExist(&Employee{}) // この判定で、外部キー設定済みのテーブルの再Sync2時に外部キーのインデックスを消せないエラーを防いでいる。
	if !exists {
		err := db.Sync2(
			new(Employee),
			new(EmployeeType),
			new(TellType),
			new(MethodType),
			new(Magazine),
			new(Customer),
			new(Regular),
			new(CountingRegular),
			new(DeliveryLog),
			new(InvoiceLog),
			new(OparateLog),
		)
		if err != nil {
			fmt.Println("Failed to sync database.", err)
			return err
		}

	}

	// FK
	err := initFK()
	if err != nil {
		fmt.Println("Failed to set foreign key.", err)
		return err
	}

	// サンプルデータ作成
	RegisterSample()

	return nil
}

// 外部キーを設定
func initFK() error {
	// 従業員
	err := InitEmployeeFK()
	if err != nil {
		return err
	}
	// 雑誌
	err = InitMagazineFK()
	if err != nil {
		return err
	}
	// 顧客
	err = InitCustomerFK()
	if err != nil {
		return err
	}
	// 定期
	err = InitRegularFK()
	if err != nil {
		return err
	}
	// 定期カウント
	err = InitCountingRegularFK()
	if err != nil {
		return err
	}
	// 納品書
	err = InitDeliveryLogFK()
	if err != nil {
		return err
	}
	// 請求書
	err = InitInvoiceLogFK()
	if err != nil {
		return err
	}
	// 操作履歴
	err = InitOparateLogFK()
	if err != nil {
		return err
	}
	return err
}

// サンプルデータ作成
// 外部キーの参照先テーブルを先に登録する必要がある。
func RegisterSample() {
	// 基本データ
	CreateEmployeeTypeData()	// 従業員種別
	CreateTellTypeTestData()	// 連絡方法
	CreateMethodTypeTestData()	// 支払方法

	// テストデータ
	CreateEmployeeTestData()	// 従業員
	CreateMagazineTestData()	// 雑誌
	CreateCustomerTestData()	// 顧客
	CreateRegularTestData()		// 定期購読
	CreateCountingRegularTestData()	// 集計定期購読
	CreateDeliveryLogTestData()		// 納品履歴
	CreateInvoiceLogTestData()		// 請求履歴
	CreateOparateLogTestData()		// 操作履歴

	
}
