package model

import (
	"time"
)

// 数取りテーブル
type CountingRegular struct {
	CountingUuid string  `xorm:"pk varchar(36)" json:"countingUUId"`
	RegularUuid string `xorm:"varchar(36) not null" json:"RegularUUId"`
	TakeData time.Time `xorm:"DATETIME not null" json:"taskData"`
}

// テーブル名
func (CountingRegular) TableName() string {
	return "counting_regulars"
}

// テストデータ
func CreateCountingRegularTestData() {
	regular1 := &CountingRegular{
		CountingUuid: "ff934ac7-ab6c-4dc9-8449-e2bcb4b69d28",
		RegularUuid: "ff934ac7-ab6c-4dc9-8449-e2bcb4b69d29",
		TakeData: time.Now(),
	}
	db.Insert(regular1)
}

// FK制約の追加
func InitCountingRegularFK() error {
	// regularid
	_, err := db.Exec("ALTER TABLE counting_regulars ADD FOREIGN KEY (regular_uuid) REFERENCES regulars(regular_uuid) ON DELETE CASCADE ON UPDATE CASCADE")
	if err != nil {
		return err
	}
	return nil
}