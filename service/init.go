package service

import (
	"auto_slip_api/model"
	"fmt"
	
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"

	"github.com/joho/godotenv"
)

var DbEngine *xorm.Engine

func Init() error{

	// .envから定数をプロセスの環境変数にロード
	err := godotenv.Load(".env") // エラーを格納
	if err != nil {             // エラーがあったら
		fmt.Println("Error loading .env file", err)
		return err
	}


	// DB初期化
	err = model.DBConnect() // 接続
	if err != nil {
		
		return err
	}
	err = model.MigrationTable() // テーブル作成
	if err != nil {
		return err
	}

	// 接続を取得
	db := model.DBInstance()
	db.ShowSQL(true)       // SQL文の表示
	db.SetMaxOpenConns(10) // 接続数を制限
	return nil
}

